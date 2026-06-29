// 数据库迁移工具 - 智能求职数据库
// 使用说明请查看: README.md
// 常用命令请查看: Makefile
// 执行迁移: make MIGRATE-ALL 或 go run main.go migrate all
package main

import (
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	pgmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/cobra"
	"github.com/yylego/go-migrate/cobramigration"
	"github.com/yylego/go-migrate/migrationparam"
	"github.com/yylego/go-migrate/migrationstate"
	"github.com/yylego/go-migrate/newmigrate"
	"github.com/yylego/must"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/rese"
	"github.com/yylego/runpath"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/exports/expconfig"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/exports/expmodels"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := expconfig.DatabaseDSN()

	scriptsInRoot := osmustexist.ROOT(runpath.PARENT.Join("scripts"))
	zaplog.LOG.Debug("migrate", zap.String("root", scriptsInRoot))

	newDB := func() *gorm.DB {
		return rese.P1(gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             200 * time.Millisecond,
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
					ParameterizedQueries:      true,
					Colorful:                  true,
				},
			),
		}))
	}

	newMigration := func(database *gorm.DB) *migrate.Migrate {
		return rese.P1(newmigrate.NewWithScriptsAndDatabase(
			&newmigrate.ScriptsAndDatabaseParam{
				ScriptsInRoot:    scriptsInRoot,
				DatabaseName:     "postgres",
				DatabaseInstance: rese.V1(pgmigrate.WithInstance(rese.P1(database.DB()), &pgmigrate.Config{})),
			},
		))
	}

	param := migrationparam.NewMigrationParam(newDB, newMigration)

	objects := expmodels.Objects()

	var debugMode bool
	var rootCmd = &cobra.Command{
		Use:   "main",
		Short: "zhipin-migrate",
		Long:  "zhipin database migration tool",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			migrationparam.SetDebugMode(debugMode)
		},
	}
	rootCmd.PersistentFlags().BoolVar(&debugMode, "debug", false, "enable debug mode")
	rootCmd.AddCommand(cobramigration.NewMigrateCmd(param))
	rootCmd.AddCommand(migrationstate.NewStatusCmd(&migrationstate.Config{
		Param:       param,
		ScriptsPath: scriptsInRoot,
		Objects:     objects,
	}))

	must.Done(rootCmd.Execute())
}
