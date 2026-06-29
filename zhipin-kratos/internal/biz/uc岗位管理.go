package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v3/errors"
	"github.com/yylego/gormcnm"
	"github.com/yylego/gormrepo"
	"github.com/yylego/gormrepo/gormclass"
	"github.com/yylego/kratos-ebz/ebzkratos"
	"github.com/yylego/kratos-gorm/gormkratos"
	"github.com/yylego/kratos-zap/zapkratos"
	"github.com/yylego/must"
	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/data"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"
	"github.com/yylego/zaplog"
	"gorm.io/gorm"
)

type Uc岗位管理 struct {
	data   *data.Data
	repo   *gormrepo.Repo[models.T岗位, *models.T岗位Columns]
	zapLog *zaplog.Zap
}

func NewUc岗位管理(data *data.Data, zapKratos *zapkratos.ZapKratos) *Uc岗位管理 {
	return &Uc岗位管理{
		data:   data,
		repo:   gormrepo.NewRepo(gormclass.Use(&models.T岗位{})),
		zapLog: zapKratos.SubZap(),
	}
}

type Req创建岗位 struct {
	J岗位编号 string
	T岗位名称 string
	C公司名称 string
	S薪资范围 string
	S薪资下限 int32
	S薪资上限 int32
	C城市名称 string
	L岗位链接 string
	R招聘者  string
	E招聘者号 string
	I猎头标记 bool
	S岗位状态 models.E岗位状态
	D岗位职责 string
	R岗位要求 string
	N备注信息 string
}

func (uc *Uc岗位管理) Xqt创建岗位(ctx context.Context, req *Req创建岗位) (*models.T岗位, *ebzkratos.Ebz) {
	if len(req.J岗位编号) != 28 {
		return nil, ebzkratos.New(pb.ErrorInvalidJobId("jobId must be 28 chars, got %d", len(req.J岗位编号)))
	}
	must.Nice(req.T岗位名称)
	must.Nice(req.C公司名称)
	must.Nice(req.C城市名称)
	must.Nice(req.S薪资范围)
	must.Nice(string(req.S岗位状态))

	v岗位 := &models.T岗位{
		J岗位编号: req.J岗位编号, T岗位名称: req.T岗位名称, C公司名称: req.C公司名称,
		S薪资范围: req.S薪资范围, S薪资下限: req.S薪资下限, S薪资上限: req.S薪资上限,
		C城市名称: req.C城市名称, L岗位链接: req.L岗位链接, R招聘者: req.R招聘者,
		E招聘者号: req.E招聘者号, I猎头标记: req.I猎头标记, S岗位状态: req.S岗位状态,
		D岗位职责: req.D岗位职责, R岗位要求: req.R岗位要求, N备注信息: req.N备注信息,
	}

	db := uc.data.DB()
	if erk, err := gormkratos.Transaction(ctx, db, func(db *gorm.DB) *errors.Error {
		if err := uc.repo.With(ctx, db).Create(v岗位); err != nil {
			return pb.ErrorDbError("create: %v", err)
		}
		return nil
	}); err != nil {
		if erk != nil {
			return nil, ebzkratos.New(erk)
		}
		return nil, ebzkratos.New(pb.ErrorTxError("tx: %v", err))
	}
	return v岗位, nil
}

func (uc *Uc岗位管理) Xqt全量更新(ctx context.Context, id uint, req *Req创建岗位) *ebzkratos.Ebz {
	must.True(id > 0)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(id))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		return cls.Kw(cls.T岗位名称.Kv(req.T岗位名称)).
			Kw(cls.C公司名称.Kv(req.C公司名称)).
			Kw(cls.S薪资范围.Kv(req.S薪资范围)).
			Kw(cls.S薪资下限.Kv(req.S薪资下限)).
			Kw(cls.S薪资上限.Kv(req.S薪资上限)).
			Kw(cls.C城市名称.Kv(req.C城市名称)).
			Kw(cls.L岗位链接.Kv(req.L岗位链接)).
			Kw(cls.R招聘者.Kv(req.R招聘者)).
			Kw(cls.E招聘者号.Kv(req.E招聘者号)).
			Kw(cls.I猎头标记.Kv(req.I猎头标记)).
			Kw(cls.S岗位状态.Kv(req.S岗位状态)).
			Kw(cls.D岗位职责.Kv(req.D岗位职责)).
			Kw(cls.R岗位要求.Kv(req.R岗位要求)).
			Kw(cls.N备注信息.Kv(req.N备注信息))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("update: %v", err))
	}
	return nil
}

type Req更新薪资 struct {
	ID    uint
	S薪资范围 string
	S薪资下限 int32
	S薪资上限 int32
}

func (uc *Uc岗位管理) Xqt更新薪资(ctx context.Context, req *Req更新薪资) *ebzkratos.Ebz {
	must.True(req.ID > 0)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(req.ID))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		return cls.Kw(cls.S薪资范围.Kv(req.S薪资范围)).Kw(cls.S薪资下限.Kv(req.S薪资下限)).Kw(cls.S薪资上限.Kv(req.S薪资上限))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("update: %v", err))
	}
	return nil
}

type Req更新状态 struct {
	ID    uint
	S岗位状态 models.E岗位状态
}

func (uc *Uc岗位管理) Xqt更新状态(ctx context.Context, req *Req更新状态) *ebzkratos.Ebz {
	must.True(req.ID > 0)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(req.ID))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		return cls.Kw(cls.S岗位状态.Kv(req.S岗位状态))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("update: %v", err))
	}
	return nil
}

type Req标记跳过 struct {
	ID    uint
	S跳过原因 string
}

func (uc *Uc岗位管理) Xqt标记跳过(ctx context.Context, req *Req标记跳过) *ebzkratos.Ebz {
	must.True(req.ID > 0)
	must.Nice(req.S跳过原因)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(req.ID))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		return cls.Kw(cls.S岗位状态.Kv(models.C岗位状态_已跳过)).Kw(cls.S跳过原因.Kv(req.S跳过原因))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("update: %v", err))
	}
	return nil
}

type Req更新职责 struct {
	ID    uint
	D岗位职责 string
}

func (uc *Uc岗位管理) Xqt更新职责(ctx context.Context, req *Req更新职责) *ebzkratos.Ebz {
	must.True(req.ID > 0)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(req.ID))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		return cls.Kw(cls.D岗位职责.Kv(req.D岗位职责))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("update: %v", err))
	}
	return nil
}

type Req更新要求 struct {
	ID    uint
	R岗位要求 string
}

func (uc *Uc岗位管理) Xqt更新要求(ctx context.Context, req *Req更新要求) *ebzkratos.Ebz {
	must.True(req.ID > 0)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(req.ID))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		return cls.Kw(cls.R岗位要求.Kv(req.R岗位要求))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("update: %v", err))
	}
	return nil
}

type Req更新匹配度 struct {
	ID   uint
	M匹配度 int32
}

func (uc *Uc岗位管理) Xqt更新匹配度(ctx context.Context, req *Req更新匹配度) *ebzkratos.Ebz {
	must.True(req.ID > 0)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(req.ID))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		return cls.Kw(cls.M匹配度.Kv(req.M匹配度))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("update: %v", err))
	}
	return nil
}

type Req更新备注 struct {
	ID    uint
	N备注信息 string
}

func (uc *Uc岗位管理) Xqt更新备注(ctx context.Context, req *Req更新备注) *ebzkratos.Ebz {
	must.True(req.ID > 0)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(req.ID))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		return cls.Kw(cls.N备注信息.Kv(req.N备注信息))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("update: %v", err))
	}
	return nil
}

type Req更新招聘者 struct {
	ID   uint
	R招聘者 string
}

func (uc *Uc岗位管理) Xqt更新招聘者(ctx context.Context, req *Req更新招聘者) *ebzkratos.Ebz {
	must.True(req.ID > 0)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(req.ID))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		return cls.Kw(cls.R招聘者.Kv(req.R招聘者))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("update: %v", err))
	}
	return nil
}

func (uc *Uc岗位管理) Get获取岗位(ctx context.Context, id uint) (*models.T岗位, *ebzkratos.Ebz) {
	must.True(id > 0)
	db := uc.data.DB()
	v岗位, erb := uc.repo.With(ctx, db).FirstE(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(id))
	})
	if erb != nil {
		if erb.NotExist {
			return nil, ebzkratos.New(pb.ErrorPositionNotFound("id=%d", id))
		}
		return nil, ebzkratos.New(pb.ErrorDbError("db: %v", erb.Cause))
	}
	return v岗位, nil
}

func (uc *Uc岗位管理) Get按编号查(ctx context.Context, j岗位编号 string) (*models.T岗位, *ebzkratos.Ebz) {
	must.Nice(j岗位编号)
	db := uc.data.DB()
	v岗位, erb := uc.repo.With(ctx, db).FirstE(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.J岗位编号.Eq(j岗位编号))
	})
	if erb != nil {
		if erb.NotExist {
			return nil, nil
		}
		return nil, ebzkratos.New(pb.ErrorDbError("db: %v", erb.Cause))
	}
	return v岗位, nil
}

type Req岗位列表 struct {
	S岗位状态    models.E岗位状态
	Page     int32
	PageSize int32
}

type Res岗位列表 struct {
	Items []*models.T岗位
	Total int64
}

func (uc *Uc岗位管理) Get岗位列表(ctx context.Context, req *Req岗位列表) (*Res岗位列表, *ebzkratos.Ebz) {
	db := uc.data.DB()
	v岗位们, total, err := uc.repo.With(ctx, db).FindC(
		func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
			if req.S岗位状态 != "" {
				db = db.Where(cls.S岗位状态.Eq(req.S岗位状态))
			}
			return db
		},
		func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
			if req.Page > 0 && req.PageSize > 0 {
				db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
			}
			return db.Order(cls.ID.Ob("DESC").Ox())
		},
	)
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("list: %v", err))
	}
	return &Res岗位列表{Items: v岗位们, Total: total}, nil
}

type Res检查编号 struct {
	Exists bool
	ID     uint
	S岗位状态  models.E岗位状态
}

func (uc *Uc岗位管理) Get检查编号(ctx context.Context, j岗位编号 string) (*Res检查编号, *ebzkratos.Ebz) {
	v岗位, ebz := uc.Get按编号查(ctx, j岗位编号)
	if ebz != nil || v岗位 == nil {
		return nil, nil
	}
	return &Res检查编号{Exists: true, ID: v岗位.ID, S岗位状态: v岗位.S岗位状态}, nil
}

func (uc *Uc岗位管理) Get批量检查编号(ctx context.Context, jobIds []string) (map[string]*models.T岗位, *ebzkratos.Ebz) {
	if len(jobIds) == 0 {
		return nil, nil
	}
	db := uc.data.DB()
	v岗位们, err := uc.repo.With(ctx, db).Find(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.J岗位编号.In(jobIds))
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("batch check: %v", err))
	}
	result := make(map[string]*models.T岗位, len(v岗位们))
	for _, v := range v岗位们 {
		result[v.J岗位编号] = v
	}
	return result, nil
}

type Req更新招聘者号 struct {
	ID    uint
	E招聘者号 string
}

func (uc *Uc岗位管理) Xqt更新招聘者号(ctx context.Context, req *Req更新招聘者号) *ebzkratos.Ebz {
	must.True(req.ID > 0)
	db := uc.data.DB()
	if err := uc.repo.With(ctx, db).UpdatesM(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.ID.Eq(req.ID))
	}, func(cls *models.T岗位Columns) gormcnm.ColumnValueMap {
		return cls.Kw(cls.E招聘者号.Kv(req.E招聘者号))
	}); err != nil {
		return ebzkratos.New(pb.ErrorDbError("update: %v", err))
	}
	return nil
}

// Get待跟进 finds positions not contacted in the last staleHours hours
func (uc *Uc岗位管理) Get待跟进(ctx context.Context, staleHours int32) (*Res岗位列表, *ebzkratos.Ebz) {
	db := uc.data.DB()
	cutoff := time.Now().Add(-time.Duration(staleHours) * time.Hour).Unix()
	v岗位们, err := uc.repo.With(ctx, db).Find(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.S岗位状态.In(models.C岗位活跃状态)).
			Where(cls.L最后沟通.Gt(int64(0))).
			Where(cls.L最后沟通.Lte(cutoff)).
			Order(cls.L最后沟通.Ob("ASC").Ox())
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("list: %v", err))
	}
	return &Res岗位列表{Items: v岗位们, Total: int64(len(v岗位们))}, nil
}

// Get待回复 finds positions where last message is from recruiter (we need to reply)
func (uc *Uc岗位管理) Get待回复(ctx context.Context) (*Res岗位列表, *ebzkratos.Ebz) {
	db := uc.data.DB()
	v岗位们, err := uc.repo.With(ctx, db).Find(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.S岗位状态.In(models.C岗位活跃状态)).
			Where(cls.L最后方向.Eq(int32(1))). // 1 = incoming from recruiter
			Order(cls.L最后沟通.Ob("ASC").Ox())
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("list: %v", err))
	}
	return &Res岗位列表{Items: v岗位们, Total: int64(len(v岗位们))}, nil
}

// Get待补发 finds positions that need resume resend (last resume != given version)
func (uc *Uc岗位管理) Get待补发(ctx context.Context, latestVersion string) (*Res岗位列表, *ebzkratos.Ebz) {
	must.Nice(latestVersion)
	db := uc.data.DB()
	v岗位们, err := uc.repo.With(ctx, db).Find(func(db *gorm.DB, cls *models.T岗位Columns) *gorm.DB {
		return db.Where(cls.S岗位状态.In(models.C岗位活跃状态)).
			Where(cls.L简历版本.Qx("=?", "").OR(cls.L简历版本.Qx("!=?", latestVersion)).Qx2()).
			Order(cls.ID.Ob("DESC").Ox())
	})
	if err != nil {
		return nil, ebzkratos.New(pb.ErrorDbError("list: %v", err))
	}
	return &Res岗位列表{Items: v岗位们, Total: int64(len(v岗位们))}, nil
}
