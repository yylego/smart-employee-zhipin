package data

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/yylego/go-migrate/checkmigration"
	"github.com/yylego/kratos-zap/zapkratos"
	"github.com/yylego/must"
	"github.com/yylego/rese"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/conf"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	loggergorm "gorm.io/gorm/logger"
)

var ProviderSet = wire.NewSet(NewData)

type Data struct {
	db *gorm.DB
}

func NewData(c *conf.Data, zapKratos *zapkratos.ZapKratos) (*Data, func(), error) {
	zapLog := zapKratos.SubZap()
	zapLog.SUG.Info("creating data resources")

	var db *gorm.DB
	switch must.Nice(c.Database.Driver) {
	case "sqlite3":
		dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
		db = rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: loggergorm.Default.LogMode(loggergorm.Info),
		}))
		must.Done(db.AutoMigrate(models.Objects()...))
	case "postgres":
		dsn := fmt.Sprintf("postgres://%s:%s@%s&TimeZone=UTC",
			must.Nice(c.Database.Username),
			must.Nice(c.Database.Password),
			must.Nice(c.Database.Source),
		)
		db = rese.P1(gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: loggergorm.Default.LogMode(loggergorm.Info),
		}))
		checkmigration.CheckMigrate(db, models.Objects())
	default:
		panic(fmt.Sprintf("UNSUPPORTED DATABASE DRIVER: %s", c.Database.Driver))
	}

	cleanup := func() {
		zapLog.SUG.Info("closing the data resources")
		must.Done(rese.P1(db.DB()).Close())
	}
	return &Data{db: db}, cleanup, nil
}

func (d *Data) DB() *gorm.DB {
	return d.db
}
