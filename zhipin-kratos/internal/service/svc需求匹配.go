package service

import (
	"context"

	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/biz"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/enums"
)

type Svc需求匹配 struct {
	pb.UnimplementedRequirementItemServiceServer

	uc *biz.Uc需求匹配
}

func NewSvc需求匹配(uc *biz.Uc需求匹配) *Svc需求匹配 {
	return &Svc需求匹配{uc: uc}
}

func (s *Svc需求匹配) SetRequirementItems(ctx context.Context, req *pb.SetRequirementItemsReq) (*pb.SetRequirementItemsResp, error) {
	items := make([]*biz.Req需求匹配项, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, &biz.Req需求匹配项{
			R岗位要求: item.Requirement,
			M匹配状态: enums.Enum匹配状态映射表.GetByCode(item.MatchStatus).Basic(),
			R简历对应: item.ResumePoint,
			R补充说明: item.Remark,
			S排序序号: item.SortIndex,
		})
	}
	count, ebz := s.uc.Xqt批量设置(ctx, req.JobId, items)

	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.SetRequirementItemsResp{Count: count}, nil
}

func (s *Svc需求匹配) ListRequirementItems(ctx context.Context, req *pb.ListRequirementItemsReq) (*pb.ListRequirementItemsResp, error) {
	v匹配们, ebz := s.uc.Get匹配列表(ctx, req.JobId)
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.RequirementItemResp, 0, len(v匹配们))
	for _, v := range v匹配们 {
		items = append(items, &pb.RequirementItemResp{
			Id:          uint64(v.ID),
			JobId:       v.J岗位编号,
			Requirement: v.R岗位要求,
			MatchStatus: int32(enums.Enum匹配状态映射表.GetByBasic(v.M匹配状态).Proto()),
			ResumePoint: v.R简历对应,
			Remark:      v.R补充说明,
			SortIndex:   v.S排序序号,
		})
	}
	return &pb.ListRequirementItemsResp{Items: items}, nil
}
