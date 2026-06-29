package enums

import (
	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"

	"github.com/yylego/protoenum"
	"github.com/yylego/rese"
)

var Enum岗位状态映射表 = rese.P1(protoenum.NewEnums(
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_UNKNOWN, models.C岗位状态_状态未知),
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_PENDING, models.C岗位状态_待处理),
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_SKIPPED, models.C岗位状态_已跳过),
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_CHAT_LIMITED, models.C岗位状态_开聊限制),
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_MSG_SENT, models.C岗位状态_已发消息),
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_REPLIED, models.C岗位状态_已回复),
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_RESUME_SENT, models.C岗位状态_已发简历),
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_INTERVIEWING, models.C岗位状态_面试中),
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_OFFERED, models.C岗位状态_已拿到),
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_REJECTED, models.C岗位状态_已拒绝),
	protoenum.NewEnum(pb.PositionStatus_POSITION_STATUS_NO_CONTACT, models.C岗位状态_不再联系),
))
