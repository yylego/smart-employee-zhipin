package service

import (
	"context"
	"time"

	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/admin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/biz"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/enums"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"
)

type Svc管理面板 struct {
	pb.UnimplementedAdminServiceServer

	uc岗位 *biz.Uc岗位管理
	uc匹配 *biz.Uc需求匹配
	uc沟通 *biz.Uc沟通管理
}

func NewSvc管理面板(uc岗位 *biz.Uc岗位管理, uc匹配 *biz.Uc需求匹配, uc沟通 *biz.Uc沟通管理) *Svc管理面板 {
	return &Svc管理面板{uc岗位: uc岗位, uc匹配: uc匹配, uc沟通: uc沟通}
}

func (s *Svc管理面板) ListTodayPositions(ctx context.Context, req *pb.ListTodayPositionsReq) (*pb.AdminPositionListResp, error) {
	// find positions contacted today (last 24 hours)
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// get all positions and filter by today's lastCommAt
	res, ebz := s.uc岗位.Get岗位列表(ctx, &biz.Req岗位列表{Page: 1, PageSize: 10000})
	if ebz != nil {
		return nil, ebz.Erk
	}
	cutoff := todayStart.Unix()
	items := make([]*pb.AdminPositionItem, 0)
	for _, v := range res.Items {
		if v.L最后沟通 >= cutoff {
			items = append(items, toAdminPositionItem(v))
		}
	}
	return &pb.AdminPositionListResp{Items: items, Total: int32(len(items))}, nil
}

func (s *Svc管理面板) ListAllPositions(ctx context.Context, req *pb.ListAllPositionsReq) (*pb.AdminPositionListResp, error) {
	var dbStatus models.E岗位状态
	if req.Status > 0 {
		dbStatus = enums.Enum岗位状态映射表.MustGetByCode(req.Status).Basic()
	}
	res, ebz := s.uc岗位.Get岗位列表(ctx, &biz.Req岗位列表{S岗位状态: dbStatus, Page: req.Page, PageSize: req.PageSize})
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.AdminPositionItem, 0, len(res.Items))
	for _, v := range res.Items {
		items = append(items, toAdminPositionItem(v))
	}
	return &pb.AdminPositionListResp{Items: items, Total: int32(res.Total)}, nil
}

func (s *Svc管理面板) GetPositionDetail(ctx context.Context, req *pb.GetPositionDetailReq) (*pb.AdminPositionDetailResp, error) {
	v岗位, ebz := s.uc岗位.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	v匹配们, ebz := s.uc匹配.Get匹配列表(ctx, v岗位.J岗位编号)
	if ebz != nil {
		return nil, ebz.Erk
	}
	v沟通们, ebz := s.uc沟通.Get聊天记录(ctx, v岗位.J岗位编号)
	if ebz != nil {
		return nil, ebz.Erk
	}

	matchItems := make([]*pb.AdminMatchItem, 0, len(v匹配们))
	for _, v := range v匹配们 {
		matchItems = append(matchItems, &pb.AdminMatchItem{
			Requirement: v.R岗位要求,
			MatchStatus: int32(enums.Enum匹配状态映射表.MustGetByBasic(v.M匹配状态).Proto()),
			ResumePoint: v.R简历对应,
			Remark:      v.R补充说明,
		})
	}

	chatMessages := make([]*pb.AdminChatMessage, 0, len(v沟通们))
	for _, v := range v沟通们 {
		chatMessages = append(chatMessages, &pb.AdminChatMessage{
			Direction:     v.D消息方向,
			Content:       v.C消息内容,
			Timestamp:     v.T消息时间,
			IsResume:      v.B简历消息,
			ResumeVersion: v.R简历版本,
		})
	}

	return &pb.AdminPositionDetailResp{
		Position:     toAdminPositionItem(v岗位),
		MatchItems:   matchItems,
		ChatMessages: chatMessages,
		Duties:         v岗位.D岗位职责,
		Requirements:   v岗位.R岗位要求,
		Notes:          v岗位.N备注信息,
		SkipReason:     v岗位.S跳过原因,
	}, nil
}

func (s *Svc管理面板) GetStats(ctx context.Context, req *pb.GetStatsReq) (*pb.GetStatsResp, error) {
	res, ebz := s.uc岗位.Get岗位列表(ctx, &biz.Req岗位列表{Page: 1, PageSize: 10000})
	if ebz != nil {
		return nil, ebz.Erk
	}
	countMap := make(map[int32]int32)
	for _, v := range res.Items {
		pbStatus := int32(enums.Enum岗位状态映射表.MustGetByBasic(v.S岗位状态).Proto())
		countMap[pbStatus]++
	}
	statusCounts := make([]*pb.StatusCount, 0, len(countMap))
	for status, count := range countMap {
		statusCounts = append(statusCounts, &pb.StatusCount{Status: status, Count: count})
	}
	return &pb.GetStatsResp{Total: int32(res.Total), StatusCounts: statusCounts}, nil
}

func toAdminPositionItem(v *models.T岗位) *pb.AdminPositionItem {
	return &pb.AdminPositionItem{
		Id:          uint64(v.ID),
		JobId:       v.J岗位编号,
		Title:       v.T岗位名称,
		Company:     v.C公司名称,
		SalaryRange: v.S薪资范围,
		SalaryMin:   v.S薪资下限,
		SalaryMax:   v.S薪资上限,
		City:        v.C城市名称,
		Recruiter:   v.R招聘者,
		IsHunter:    v.I猎头标记,
		Status:      int32(enums.Enum岗位状态映射表.MustGetByBasic(v.S岗位状态).Proto()),
		MatchRate:   v.M匹配度,
		LastResume:  v.L简历版本,
		LastCommAt:  v.L最后沟通,
		LastCommDir: v.L最后方向,
		Link:        v.L岗位链接,
	}
}
