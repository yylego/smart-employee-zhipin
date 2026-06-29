package expconfig

import (
	"fmt"
	"path/filepath"

	"github.com/go-kratos/kratos/v3/config"
	"github.com/go-kratos/kratos/v3/config/file"
	"github.com/yylego/must"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/rese"
	zhipin_kratos "github.com/yylego/smart-employee-zhipin/zhipin-kratos"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/conf"
)

// LoadBootstrap loads the Bootstrap config from bin/configs if present, otherwise from configs
func LoadBootstrap() *conf.Bootstrap {
	root := zhipin_kratos.SourceRoot()
	configPath := filepath.Join(root, "bin", "configs")
	if !osmustexist.IsRoot(configPath) {
		configPath = filepath.Join(root, "configs")
	}
	osmustexist.ROOT(configPath)

	c := config.New(
		config.WithSource(
			file.NewSource(configPath),
		),
	)
	must.Done(c.Load())
	defer rese.F0(c.Close)

	var cfg conf.Bootstrap
	must.Done(c.Scan(&cfg))
	return &cfg
}

// DatabaseDSN returns the PostgreSQL DSN from config
func DatabaseDSN() string {
	return BuildDSN(LoadBootstrap().Data.Database)
}

// BuildDSN builds the PostgreSQL DSN from database config
func BuildDSN(db *conf.Data_Database) string {
	must.Same(db.Driver, "postgres")
	return fmt.Sprintf("postgres://%s:%s@%s&TimeZone=UTC",
		must.Nice(db.Username),
		must.Nice(db.Password),
		must.Nice(db.Source),
	)
}
