package enums

import (
	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"

	"github.com/yylego/protoenum"
	"github.com/yylego/rese"
)

var Enum匹配状态映射表 = rese.P1(protoenum.NewEnums(
	protoenum.NewEnum(pb.MatchStatus_MATCH_STATUS_UNKNOWN, models.C匹配状态_状态未知),
	protoenum.NewEnum(pb.MatchStatus_MATCH_STATUS_MATCHED, models.C匹配状态_匹配),
	protoenum.NewEnum(pb.MatchStatus_MATCH_STATUS_PARTIAL, models.C匹配状态_部分匹配),
	protoenum.NewEnum(pb.MatchStatus_MATCH_STATUS_MISSING, models.C匹配状态_不匹配),
))
