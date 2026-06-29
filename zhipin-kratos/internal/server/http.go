package server

import (
	"github.com/go-kratos/kratos/v3/middleware/logging"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/yylego/kratos-swaggo/swaggokratos"
	"github.com/yylego/kratos-swaggo/swaggokratos/swaggogin"
	"github.com/yylego/kratos-zap/zapkratos"
	zhipin_kratos "github.com/yylego/smart-employee-zhipin/zhipin-kratos"
	adminpb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/admin"
	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/conf"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/service"
)

func NewHTTPServer(
	c *conf.Server,
	position *service.Svc岗位管理,
	communication *service.Svc沟通管理,
	matchItem *service.Svc需求匹配,
	blacklist *service.Svc黑名单管理,
	admin *service.Svc管理面板,
	zapKratos *zapkratos.ZapKratos,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(zapKratos.GetLogger("http-request")),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Address != "" {
		opts = append(opts, http.Address(c.Http.Address))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	pb.RegisterPositionServiceHTTPServer(srv, position)
	pb.RegisterCommunicationServiceHTTPServer(srv, communication)
	pb.RegisterRequirementItemServiceHTTPServer(srv, matchItem)
	pb.RegisterBlacklistServiceHTTPServer(srv, blacklist)
	adminpb.RegisterAdminServiceHTTPServer(srv, admin)

	serveSwaggerHttpDocument(c, srv, zapKratos)
	return srv
}

func serveSwaggerHttpDocument(c *conf.Server, srv *http.Server, zapKratos *zapkratos.ZapKratos) {
	zapLog := zapKratos.SubZap()
	zapLog.SUG.Infoln("准备添加接口文档")

	swaggokratos.RegisterSwaggoHTTPServer(srv, "/doc/", []*swaggogin.Param{
		{
			SwaggerPath: "/swagger/a/*any",
			ExplorePath: "/abc/openapi-a.yaml",
			ContentData: zhipin_kratos.GetOpenapiContent("smart-employee-zhipin"),
		},
	})

	zapLog.SUG.Infoln("[DOC]", "(http://127.0.0.1:"+swaggokratos.MustGetPortNum(c.Http.Address)+"/doc/swagger/a/index.html)")
	zapLog.SUG.Infoln("接口文档添加成功")
}
