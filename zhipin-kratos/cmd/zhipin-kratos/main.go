package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/go-kratos/kratos/contrib/encoding/json/v3"
	"github.com/go-kratos/kratos/v3"
	"github.com/go-kratos/kratos/v3/config"
	"github.com/go-kratos/kratos/v3/config/file"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/yylego/done"
	"github.com/yylego/kratos-errors/errorskratos/newerk"
	"github.com/yylego/kratos-zap/zapkratos"
	"github.com/yylego/must"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/rese"
	zhipin_kratos "github.com/yylego/smart-employee-zhipin/zhipin-kratos"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/conf"
	"github.com/yylego/tern/zerotern"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")

	// Configure JSON field naming style for HTTP responses
	// 配置 HTTP 响应的 JSON 字段命名风格，使用小写驼峰命名确保跨语言兼容性
	json.MarshalOptions.UseProtoNames = false

	// Set UseEnumNumbers to true to serialize enums as numbers instead of strings
	// 设置 UseEnumNumbers 为 true 使枚举序列化为数字而非字符串
	json.MarshalOptions.UseEnumNumbers = true

	// Set metadata field name to pass numeric enum value to frontend
	// 设置 metadata 字段名用于传递枚举数值给前端
	newerk.SetReasonCodeFieldName("numeric_reason_code_enum")
}

func newApp(gs *grpc.Server, hs *http.Server, zapKratos *zapkratos.ZapKratos) *kratos.App {
	return kratos.New(
		kratos.ID(done.VCE(os.Hostname()).Omit()),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(zapKratos.NewLogger("zhipin-kratos")),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()

	{
		rootBin := filepath.Join(zhipin_kratos.SourceRoot(), "bin")
		must.Done(os.MkdirAll(rootBin, 0755))
		path1 := filepath.Join(rootBin, "log-newest.log")
		path2 := filepath.Join(rootBin, "log-oldest.log")

		if osmustexist.IsFile(path1) {
			must.Done(os.Truncate(path1, 0))
		}

		zaplog.SetLog(rese.P1(zaplog.NewZapLog(zaplog.NewConfig().
			AddOutputPaths(path1, path2))).With(
			zap.String("service", zerotern.VF(Name, func() string {
				return filepath.Base(zhipin_kratos.SourceRoot())
			})),
			zap.String("version", zerotern.VV(Version, "v0.0.0")),
		))
	}

	zapKratos := zapkratos.NewZapKratos(zaplog.LOGGER, zapkratos.NewOptions())
	zapLog := zapKratos.SubZap()
	zapLog.LOG.Info("application starting...")
	zapLog.LOG.Info("reading config", zap.String("path", flagconf))

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer rese.F0(c.Close)

	must.Done(c.Load())

	var cfg conf.Bootstrap
	must.Done(c.Scan(&cfg))

	app, cleanup := rese.V2(wireApp(cfg.Server, cfg.Data, zapKratos))
	defer cleanup()

	// start and wait for stop signal
	must.Done(app.Run())
}
