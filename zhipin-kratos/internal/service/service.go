package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewSvc岗位管理, NewSvc沟通管理, NewSvc匹配管理, NewSvc黑名单管理, NewSvc管理面板)
