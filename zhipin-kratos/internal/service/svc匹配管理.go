package service

import (
	"context"

	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/biz"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/enums"
)

type Svc匹配管理 struct {
	pb.UnimplementedMatchItemServiceServer

	uc *biz.Uc匹配管理
}

func NewSvc匹配管理(uc *biz.Uc匹配管理) *Svc匹配管理 {
	return &Svc匹配管理{uc: uc}
}

func (s *Svc匹配管理) SetMatchItems(ctx context.Context, req *pb.SetMatchItemsReq) (*pb.SetMatchItemsResp, error) {
	items := make([]*biz.Req匹配项, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, &biz.Req匹配项{
			R岗位要求: item.Requirement,
			M匹配状态: enums.Enum匹配状态映射表.MustGetByCode(item.MatchStatus).Basic(),
			R简历对应: item.ResumePoint,
			R补充说明: item.Remark,
			S排序序号: item.SortIndex,
		})
	}
	count, ebz := s.uc.Xqt批量设置(ctx, uint(req.PositionId), items)
	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.SetMatchItemsResp{Count: count}, nil
}

func (s *Svc匹配管理) ListMatchItems(ctx context.Context, req *pb.ListMatchItemsReq) (*pb.ListMatchItemsResp, error) {
	v匹配们, ebz := s.uc.Get匹配列表(ctx, uint(req.PositionId))
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.MatchItemResp, 0, len(v匹配们))
	for _, v := range v匹配们 {
		items = append(items, &pb.MatchItemResp{
			Id:          uint64(v.ID),
			PositionId:  uint64(v.P岗位主键),
			Requirement: v.R岗位要求,
			MatchStatus: int32(enums.Enum匹配状态映射表.MustGetByBasic(v.M匹配状态).Proto()),
			ResumePoint: v.R简历对应,
			Remark:      v.R补充说明,
			SortIndex:   v.S排序序号,
		})
	}
	return &pb.ListMatchItemsResp{Items: items}, nil
}
