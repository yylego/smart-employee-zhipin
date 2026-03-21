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

type Uc需求匹配 struct {
	data     *data.Data
	repo     *gormrepo.Repo[models.T需求匹配项, *models.T需求匹配项Columns]
	repo岗位 *gormrepo.Repo[models.T岗位, *models.T岗位Columns]
	zapLog   *zaplog.Zap
}

func NewUc需求匹配(data *data.Data, zapKratos *zapkratos.ZapKratos) *Uc需求匹配 {
	return &Uc需求匹配{
		data:     data,
		repo:     gormrepo.NewRepo(gormclass.Use(&models.T需求匹配项{})),
		repo岗位: gormrepo.NewRepo(gormclass.Use(&models.T岗位{})),
		zapLog:   zapKratos.SubZap(),
	}
}

type Req需求匹配项 struct {
	R岗位要求 string
	M匹配状态 models.E匹配状态
	R简历对应 string
	R补充说明 string
	S排序序号 int32
}

func (uc *Uc需求匹配) Xqt批量设置(ctx context.Context, j岗位编号 string, items []*Req需求匹配项) (int32, *ebzkratos.Ebz) {
	must.True(len(j岗位编号) == 28)
	db := uc.data.DB()

	v岗位, err := uc.repo岗位.With(ctx, db).First(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.J岗位编号.Eq(j岗位编号))
	})
	if err != nil {
		return 0, ebzkratos.New(pb.ErrorPositionNotFound("job_id=%s", j岗位编号))
	}
	p岗位主键 := v岗位.ID

	if err := uc.repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *models.T需求匹配项Columns) *gorm.DB {
		return db.Where(cls.P岗位主键.Eq(p岗位主键))
	}); err != nil {
		return 0, ebzkratos.New(pb.ErrorDbError("delete: %v", err))
	}
	for _, item := range items {
		v匹配 := &models.T需求匹配项{
			P岗位主键: p岗位主键,
			J岗位编号: j岗位编号,
			R岗位要求: item.R岗位要求,
			M匹配状态: item.M匹配状态,
			R简历对应: item.R简历对应,
			R补充说明: item.R补充说明,
			S排序序号: item.S排序序号,
		}
		if err := uc.repo.With(ctx, db).Create(v匹配); err != nil {
			return 0, ebzkratos.New(pb.ErrorDbError("create: %v", err))
		}
	}
	return int32(len(items)), nil
}

func (uc *Uc需求匹配) Get匹配列表(ctx context.Context, j岗位编号 string) ([]*models.T需求匹配项, *ebzkratos.Ebz) {
	must.True(len(j岗位编号) == 28)
	db := uc.data.DB()

	v岗位, err := uc.repo岗位.With(ctx, db).First(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.J岗位编号.Eq(j岗位编号))
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorPositionNotFound("job_id=%s", j岗位编号))
	}

	v匹配们, err := uc.repo.With(ctx, db).Find(func(db *gorm.DB, cls *models.T需求匹配项Columns) *gorm.DB {
		return db.Where(cls.P岗位主键.Eq(v岗位.ID)).Order(cls.S排序序号.Ob("ASC").Ox())
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("list: %v", err))
	}
	return v匹配们, nil
}
