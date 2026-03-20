package biz

import (
	"context"

	"github.com/yylego/kratos-zap/zapkratos"
	"github.com/yylego/zaplog"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/kratos-ebz/ebzkratos"
	"github.com/yylego/must"
	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/data"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"
	"gorm.io/gorm"
)

type Uc黑名单管理 struct {
	data *data.Data
	repo *gormrepo.Repo[models.T黑名单, *models.T黑名单Columns]
	zapLog *zaplog.Zap
}

func NewUc黑名单管理(data *data.Data, zapKratos *zapkratos.ZapKratos) *Uc黑名单管理 {
	return &Uc黑名单管理{
		data: data,
		repo: gormrepo.NewRepo(gormclass.Use(&models.T黑名单{})),
		zapLog: zapKratos.SubZap(),
	}
}

type Req添加黑名单 struct {
	C公司名称 string
	R拉黑原因 string
}

func (uc *Uc黑名单管理) Xqt添加黑名单(ctx context.Context, req *Req添加黑名单) (*models.T黑名单, *ebzkratos.Ebz) {
	must.Nice(req.C公司名称)
	db := uc.data.DB()
	v黑名单 := &models.T黑名单{
		C公司名称: req.C公司名称,
		R拉黑原因: req.R拉黑原因,
	}
	if err := uc.repo.With(ctx, db).Create(v黑名单); err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("create: %v", err))
	}
	return v黑名单, nil
}

func (uc *Uc黑名单管理) Get检查黑名单(ctx context.Context, c公司名称 string) bool {
	must.Nice(c公司名称)
	db := uc.data.DB()
	v黑名单, erb := uc.repo.With(ctx, db).FirstE(func(db *gorm.DB, cls *models.T黑名单Columns) *gorm.DB {
		return db.Where(cls.C公司名称.Eq(c公司名称))
	})
	if erb != nil {
		return false
	}
	return v黑名单 != nil
}

func (uc *Uc黑名单管理) Get黑名单列表(ctx context.Context) ([]*models.T黑名单, *ebzkratos.Ebz) {
	db := uc.data.DB()
	v黑名单们, err := uc.repo.With(ctx, db).Find(func(db *gorm.DB, cls *models.T黑名单Columns) *gorm.DB {
		return db.Order(cls.ID.Ob("DESC").Ox())
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("list: %v", err))
	}
	return v黑名单们, nil
}

func (uc *Uc黑名单管理) Xqt移除黑名单(ctx context.Context, id uint) *ebzkratos.Ebz {
	must.True(id > 0)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *models.T黑名单Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(id))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("delete: %v", err))
	}
	return nil
}
