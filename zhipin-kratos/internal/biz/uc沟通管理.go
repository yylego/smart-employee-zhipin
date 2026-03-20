package biz

import (
	"context"
	"time"

	"github.com/yylego/kratos-zap/zapkratos"
	"github.com/yylego/zaplog"
	"github.com/yylego/gormcnm"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/kratos-ebz/ebzkratos"
	"github.com/yylego/must"
	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/data"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"
	"gorm.io/gorm"
)

type Uc沟通管理 struct {
	data     *data.Data
	repo     *gormrepo.Repo[models.T沟通记录, *models.T沟通记录Columns]
	repo岗位 *gormrepo.Repo[models.T岗位, *models.T岗位Columns]
	zapLog   *zaplog.Zap
}

func NewUc沟通管理(data *data.Data, zapKratos *zapkratos.ZapKratos) *Uc沟通管理 {
	return &Uc沟通管理{
		data:     data,
		repo:     gormrepo.NewRepo(gormclass.Use(&models.T沟通记录{})),
		repo岗位: gormrepo.NewRepo(gormclass.Use(&models.T岗位{})),
		zapLog: zapKratos.SubZap(),
	}
}

type Req记录事件 struct {
	P岗位主键 uint
	E事件类型 models.E事件类型
	C消息内容 string
	D消息方向 int32
}

func (uc *Uc沟通管理) Xqt记录事件(ctx context.Context, req *Req记录事件) (*models.T沟通记录, *ebzkratos.Ebz) {
	must.True(req.P岗位主键 > 0)
	must.Nice(string(req.E事件类型))

	v记录 := &models.T沟通记录{
		P岗位主键: req.P岗位主键,
		E事件类型: req.E事件类型,
		E事件时间: time.Now(),
		C消息内容: req.C消息内容,
		D消息方向: req.D消息方向,
	}

	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).Create(v记录); err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("create: %v", err))
	}

	// sync last comm info to position table for fast querying
	if err := uc.repo岗位.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(req.P岗位主键))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		kv := cls.Kw(cls.L最后沟通.Kv(v记录.E事件时间.Unix())).Kw(cls.L最后方向.Kv(req.D消息方向))
		if req.E事件类型 == models.C事件类型_发简历 {
			kv = kv.Kw(cls.L简历版本.Kv(req.C消息内容))
		}
		return kv
	}); err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("sync position: %v", err))
	}

	return v记录, nil
}

func (uc *Uc沟通管理) Get沟通列表(ctx context.Context, p岗位主键 uint) ([]*models.T沟通记录, *ebzkratos.Ebz) {
	must.True(p岗位主键 > 0)
	db := uc.data.DB()
	v记录们, err := uc.repo.With(ctx, db).Find(func(db *gorm.DB, cls *models.T沟通记录Columns) *gorm.DB {
		return db.Where(cls.P岗位主键.Eq(p岗位主键)).Order(cls.E事件时间.Ob("ASC").Ox())
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("list: %v", err))
	}
	return v记录们, nil
}
