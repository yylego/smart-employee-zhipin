package enums

import (
	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"

	"github.com/yylego/protoenum"
)

var Enum事件类型映射表 = protoenum.NewEnums(
	protoenum.NewEnum(pb.EventType_EVENT_TYPE_UNKNOWN, models.C事件类型_类型未知),
	protoenum.NewEnum(pb.EventType_EVENT_TYPE_MSG_SENT, models.C事件类型_发消息),
	protoenum.NewEnum(pb.EventType_EVENT_TYPE_MSG_RECEIVED, models.C事件类型_收消息),
	protoenum.NewEnum(pb.EventType_EVENT_TYPE_RESUME_SENT, models.C事件类型_发简历),
	protoenum.NewEnum(pb.EventType_EVENT_TYPE_CHAT_LIMITED, models.C事件类型_开聊限制),
	protoenum.NewEnum(pb.EventType_EVENT_TYPE_INTERVIEW, models.C事件类型_安排面试),
	protoenum.NewEnum(pb.EventType_EVENT_TYPE_OFFER, models.C事件类型_收到邀请),
	protoenum.NewEnum(pb.EventType_EVENT_TYPE_REJECTED, models.C事件类型_被拒绝),
)
