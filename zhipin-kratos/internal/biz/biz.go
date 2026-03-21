package biz

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUc岗位管理, NewUc需求匹配, NewUc沟通管理, NewUc黑名单管理)
