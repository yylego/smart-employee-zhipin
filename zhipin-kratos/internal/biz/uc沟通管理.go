package biz

import (
	"context"

	"github.com/yylego/gormcnm"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/kratos-ebz/ebzkratos"
	"github.com/yylego/kratos-zap/zapkratos"
	"github.com/yylego/must"
	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/data"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"
	"github.com/yylego/zaplog"
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
		zapLog:   zapKratos.SubZap(),
	}
}

type Req聊天消息 struct {
	D消息方向 int32
	C消息内容 string
	T消息时间 int64
	B简历消息 bool
	R简历版本 string
}

type Res同步聊天 struct {
	L最后沟通 int64
	L最后方向 int32
	L简历版本 string
	N消息数量 int32
}

func (uc *Uc沟通管理) Xqt同步聊天(ctx context.Context, j岗位编号 string, items []*Req聊天消息) (*Res同步聊天, *ebzkratos.Ebz) {
	must.True(len(j岗位编号) == 28)
	db := uc.data.DB()

	v岗位, err := uc.repo岗位.With(ctx, db).First(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.J岗位编号.Eq(j岗位编号))
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorPositionNotFound("job_id=%s", j岗位编号))
	}
	p岗位主键 := v岗位.ID

	// hard delete old messages for this position (not soft delete, since we do full replacement)
	cls := (&models.T沟通记录{}).Columns()
	if err := db.Unscoped().Where(cls.P岗位主键.Eq(p岗位主键)).Delete(&models.T沟通记录{}).Error; err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("delete: %v", err))
	}

	// insert all messages
	for _, item := range items {
		v记录 := &models.T沟通记录{
			P岗位主键: p岗位主键,
			J岗位编号: j岗位编号,
			D消息方向: item.D消息方向,
			C消息内容: item.C消息内容,
			T消息时间: item.T消息时间,
			B简历消息: item.B简历消息,
			R简历版本: item.R简历版本,
		}
		if err := uc.repo.With(ctx, db).Create(v记录); err != nil {
			return nil, ebzkratos.New(pb.ErrorDbError("create: %v", err))
		}
	}

	// extract summary from messages
	res := &Res同步聊天{N消息数量: int32(len(items))}
	if len(items) > 0 {
		last := items[len(items)-1]
		res.L最后沟通 = last.T消息时间
		res.L最后方向 = last.D消息方向
	}
	// find latest resume version
	for i := len(items) - 1; i >= 0; i-- {
		if items[i].B简历消息 && items[i].R简历版本 != "" {
			res.L简历版本 = items[i].R简历版本
			break
		}
	}

	// sync summary to position table
	if err := uc.repo岗位.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(p岗位主键))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		kv := cls.Kw(cls.L最后沟通.Kv(res.L最后沟通)).Kw(cls.L最后方向.Kv(res.L最后方向))
		if res.L简历版本 != "" {
			kv = kv.Kw(cls.L简历版本.Kv(res.L简历版本))
		}
		return kv
	}); err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("sync position: %v", err))
	}

	return res, nil
}

func (uc *Uc沟通管理) Get聊天记录(ctx context.Context, j岗位编号 string) ([]*models.T沟通记录, *ebzkratos.Ebz) {
	must.True(len(j岗位编号) == 28)
	db := uc.data.DB()

	v岗位, err := uc.repo岗位.With(ctx, db).First(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.J岗位编号.Eq(j岗位编号))
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorPositionNotFound("job_id=%s", j岗位编号))
	}

	v记录们, err := uc.repo.With(ctx, db).Find(func(db *gorm.DB, cls *models.T沟通记录Columns) *gorm.DB {
		return db.Where(cls.P岗位主键.Eq(v岗位.ID)).Order(cls.T消息时间.Ob("ASC").Ox())
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("get: %v", err))
	}
	return v记录们, nil
}
