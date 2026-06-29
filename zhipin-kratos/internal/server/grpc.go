package server

import (
	"github.com/go-kratos/kratos/v3/middleware/logging"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/yylego/kratos-zap/zapkratos"
	adminpb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/admin"
	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/conf"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/service"
)

func NewGRPCServer(
	c *conf.Server,
	position *service.Svc岗位管理,
	communication *service.Svc沟通管理,
	matchItem *service.Svc需求匹配,
	blacklist *service.Svc黑名单管理,
	admin *service.Svc管理面板,
	zapKratos *zapkratos.ZapKratos,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(zapKratos.GetLogger("grpc-request")),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Address != "" {
		opts = append(opts, grpc.Address(c.Grpc.Address))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterPositionServiceServer(srv, position)
	pb.RegisterCommunicationServiceServer(srv, communication)
	pb.RegisterRequirementItemServiceServer(srv, matchItem)
	pb.RegisterBlacklistServiceServer(srv, blacklist)
	adminpb.RegisterAdminServiceServer(srv, admin)
	return srv
}
