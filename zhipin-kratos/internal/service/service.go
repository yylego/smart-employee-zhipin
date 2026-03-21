package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewSvc岗位管理, NewSvc沟通管理, NewSvc需求匹配, NewSvc黑名单管理, NewSvc管理面板)
