package service

import (
	"context"

	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/biz"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/enums"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"
)

type Svc岗位管理 struct {
	pb.UnimplementedPositionServiceServer

	uc      *biz.Uc岗位管理
	ucMatch *biz.Uc需求匹配
	ucChat  *biz.Uc沟通管理
}

func NewSvc岗位管理(uc *biz.Uc岗位管理, ucMatch *biz.Uc需求匹配, ucChat *biz.Uc沟通管理) *Svc岗位管理 {
	return &Svc岗位管理{uc: uc, ucMatch: ucMatch, ucChat: ucChat}
}

func (s *Svc岗位管理) CreatePosition(ctx context.Context, req *pb.CreatePositionReq) (*pb.PositionResp, error) {
	// validate match items before creating position — reject incomplete data upfront
	if len(req.MatchItems) == 0 {
		return nil, pb.ErrorMissingField("match_items is required when creating a position")
	}

	v岗位, ebz := s.uc.Xqt创建岗位(ctx, &biz.Req创建岗位{
		J岗位编号: req.JobId,
		T岗位名称: req.Title,
		C公司名称: req.Company,
		S薪资范围: req.SalaryRange,
		S薪资下限: req.SalaryMin,
		S薪资上限: req.SalaryMax,
		C城市名称: req.City,
		L岗位链接: req.Link,
		R招聘者:  req.Recruiter,
		E招聘者号: req.EncBossId,
		I猎头标记: req.IsHunter,
		S岗位状态: enums.Enum岗位状态映射表.GetByCode(req.Status).Basic(),
		D岗位职责: req.Duties,
		R岗位要求: req.Requirements,
		N备注信息: req.Notes,
	})
	if ebz != nil {
		return nil, ebz.Erk
	}

	// create match items together with position
	{
		items := make([]*biz.Req需求匹配项, 0, len(req.MatchItems))
		for _, item := range req.MatchItems {
			items = append(items, &biz.Req需求匹配项{
				R岗位要求: item.Requirement,
				M匹配状态: enums.Enum匹配状态映射表.GetByCode(item.MatchStatus).Basic(),
				R简历对应: item.ResumePoint,
				R补充说明: item.Remark,
				S排序序号: item.SortIndex,
			})
		}
		if _, ebz := s.ucMatch.Xqt批量设置(ctx, v岗位.J岗位编号, items); ebz != nil {
			return nil, ebz.Erk
		}
	}

	return toPositionResp(v岗位), nil
}

func (s *Svc岗位管理) UpdateSalary(ctx context.Context, req *pb.UpdateSalaryReq) (*pb.PositionResp, error) {
	if ebz := s.uc.Xqt更新薪资(ctx, &biz.Req更新薪资{ID: uint(req.Id), S薪资范围: req.SalaryRange, S薪资下限: req.SalaryMin, S薪资上限: req.SalaryMax}); ebz != nil {
		return nil, ebz.Erk
	}
	v岗位, ebz := s.uc.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toPositionResp(v岗位), nil
}

func (s *Svc岗位管理) UpdateStatus(ctx context.Context, req *pb.UpdateStatusReq) (*pb.PositionResp, error) {
	if ebz := s.uc.Xqt更新状态(ctx, &biz.Req更新状态{ID: uint(req.Id), S岗位状态: enums.Enum岗位状态映射表.GetByCode(req.Status).Basic()}); ebz != nil {
		return nil, ebz.Erk
	}
	v岗位, ebz := s.uc.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toPositionResp(v岗位), nil
}

func (s *Svc岗位管理) MarkSkipped(ctx context.Context, req *pb.MarkSkippedReq) (*pb.PositionResp, error) {
	if ebz := s.uc.Xqt标记跳过(ctx, &biz.Req标记跳过{ID: uint(req.Id), S跳过原因: req.Reason}); ebz != nil {
		return nil, ebz.Erk
	}
	v岗位, ebz := s.uc.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toPositionResp(v岗位), nil
}

func (s *Svc岗位管理) UpdateDuties(ctx context.Context, req *pb.UpdateDutiesReq) (*pb.PositionResp, error) {
	if ebz := s.uc.Xqt更新职责(ctx, &biz.Req更新职责{ID: uint(req.Id), D岗位职责: req.Duties}); ebz != nil {
		return nil, ebz.Erk
	}
	v岗位, ebz := s.uc.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toPositionResp(v岗位), nil
}

func (s *Svc岗位管理) UpdateRequirements(ctx context.Context, req *pb.UpdateRequirementsReq) (*pb.PositionResp, error) {
	if ebz := s.uc.Xqt更新要求(ctx, &biz.Req更新要求{ID: uint(req.Id), R岗位要求: req.Requirements}); ebz != nil {
		return nil, ebz.Erk
	}
	v岗位, ebz := s.uc.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toPositionResp(v岗位), nil
}

func (s *Svc岗位管理) UpdateMatchRate(ctx context.Context, req *pb.UpdateMatchRateReq) (*pb.PositionResp, error) {
	if ebz := s.uc.Xqt更新匹配度(ctx, &biz.Req更新匹配度{ID: uint(req.Id), M匹配度: req.MatchRate}); ebz != nil {
		return nil, ebz.Erk
	}
	v岗位, ebz := s.uc.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toPositionResp(v岗位), nil
}

func (s *Svc岗位管理) UpdateNotes(ctx context.Context, req *pb.UpdateNotesReq) (*pb.PositionResp, error) {
	if ebz := s.uc.Xqt更新备注(ctx, &biz.Req更新备注{ID: uint(req.Id), N备注信息: req.Notes}); ebz != nil {
		return nil, ebz.Erk
	}
	v岗位, ebz := s.uc.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toPositionResp(v岗位), nil
}

func (s *Svc岗位管理) UpdateRecruiter(ctx context.Context, req *pb.UpdateRecruiterReq) (*pb.PositionResp, error) {
	if ebz := s.uc.Xqt更新招聘者(ctx, &biz.Req更新招聘者{ID: uint(req.Id), R招聘者: req.Recruiter}); ebz != nil {
		return nil, ebz.Erk
	}
	v岗位, ebz := s.uc.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toPositionResp(v岗位), nil
}

func (s *Svc岗位管理) GetPosition(ctx context.Context, req *pb.GetPositionReq) (*pb.PositionDetailResp, error) {
	v岗位, ebz := s.uc.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.PositionDetailResp{Position: toPositionResp(v岗位)}, nil
}

func (s *Svc岗位管理) GetPositionByJobId(ctx context.Context, req *pb.GetPositionByJobIdReq) (*pb.PositionDetailResp, error) {
	v岗位, ebz := s.uc.Get按编号查(ctx, req.JobId)
	if ebz != nil {
		return nil, ebz.Erk
	}
	if v岗位 == nil {
		return nil, pb.ErrorPositionNotFound("job_id=%s", req.JobId)
	}
	return &pb.PositionDetailResp{Position: toPositionResp(v岗位)}, nil
}

func (s *Svc岗位管理) ListPositions(ctx context.Context, req *pb.ListPositionsReq) (*pb.ListPositionsResp, error) {
	var dbStatus models.E岗位状态
	if req.Status > 0 {
		dbStatus = enums.Enum岗位状态映射表.GetByCode(req.Status).Basic()
	}
	res, ebz := s.uc.Get岗位列表(ctx, &biz.Req岗位列表{S岗位状态: dbStatus, Page: req.Page, PageSize: req.PageSize})
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.PositionResp, 0, len(res.Items))
	for _, v := range res.Items {
		items = append(items, toPositionResp(v))
	}
	return &pb.ListPositionsResp{Items: items, Total: int32(res.Total)}, nil
}

func (s *Svc岗位管理) CheckJobId(ctx context.Context, req *pb.CheckJobIdReq) (*pb.CheckJobIdResp, error) {
	res, ebz := s.uc.Get检查编号(ctx, req.JobId)
	if ebz != nil {
		return nil, ebz.Erk
	}
	if res == nil {
		return &pb.CheckJobIdResp{Exists: false}, nil
	}
	return &pb.CheckJobIdResp{
		Exists:     res.Exists,
		PositionId: uint64(res.ID),
		Status:     int32(enums.Enum岗位状态映射表.GetByBasic(res.S岗位状态).Proto()),
	}, nil
}

func (s *Svc岗位管理) BatchCheckJobIds(ctx context.Context, req *pb.BatchCheckJobIdsReq) (*pb.BatchCheckJobIdsResp, error) {
	found, ebz := s.uc.Get批量检查编号(ctx, req.JobIds)
	if ebz != nil {
		return nil, ebz.Erk
	}
	results := make([]*pb.JobIdCheckResult, 0, len(req.JobIds))
	for _, jobId := range req.JobIds {
		v, ok := found[jobId]
		if ok {
			results = append(results, &pb.JobIdCheckResult{
				JobId:      jobId,
				Exists:     true,
				PositionId: uint64(v.ID),
				Status:     int32(enums.Enum岗位状态映射表.GetByBasic(v.S岗位状态).Proto()),
			})
		} else {
			results = append(results, &pb.JobIdCheckResult{
				JobId:  jobId,
				Exists: false,
			})
		}
	}
	return &pb.BatchCheckJobIdsResp{Results: results}, nil
}

func toPositionResp(v *models.T岗位) *pb.PositionResp {
	return &pb.PositionResp{
		Id:           uint64(v.ID),
		JobId:        v.J岗位编号,
		Title:        v.T岗位名称,
		Company:      v.C公司名称,
		SalaryRange:  v.S薪资范围,
		SalaryMin:    v.S薪资下限,
		SalaryMax:    v.S薪资上限,
		City:         v.C城市名称,
		Link:         v.L岗位链接,
		Recruiter:    v.R招聘者,
		EncBossId:    v.E招聘者号,
		IsHunter:     v.I猎头标记,
		Status:       int32(enums.Enum岗位状态映射表.GetByBasic(v.S岗位状态).Proto()),
		SkipReason:   v.S跳过原因,
		MatchRate:    v.M匹配度,
		Notes:        v.N备注信息,
		LastCommAt:   v.L最后沟通,
		LastCommDir:  v.L最后方向,
		LastResume:   v.L简历版本,
		Duties:       v.D岗位职责,
		Requirements: v.R岗位要求,
	}
}

func (s *Svc岗位管理) ListStalePositions(ctx context.Context, req *pb.ListStalePositionsReq) (*pb.ListPositionsResp, error) {
	res, ebz := s.uc.Get待跟进(ctx, req.StaleHours)
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.PositionResp, 0, len(res.Items))
	for _, v := range res.Items {
		items = append(items, toPositionResp(v))
	}
	return &pb.ListPositionsResp{Items: items, Total: int32(res.Total)}, nil
}

func (s *Svc岗位管理) ListNeedReplyPositions(ctx context.Context, req *pb.ListNeedReplyPositionsReq) (*pb.ListPositionsResp, error) {
	res, ebz := s.uc.Get待回复(ctx)
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.PositionResp, 0, len(res.Items))
	for _, v := range res.Items {
		items = append(items, toPositionResp(v))
	}
	return &pb.ListPositionsResp{Items: items, Total: int32(res.Total)}, nil
}

func (s *Svc岗位管理) ListNeedResendPositions(ctx context.Context, req *pb.ListNeedResendPositionsReq) (*pb.ListPositionsResp, error) {
	res, ebz := s.uc.Get待补发(ctx, req.LatestResumeVersion)
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.PositionResp, 0, len(res.Items))
	for _, v := range res.Items {
		items = append(items, toPositionResp(v))
	}
	return &pb.ListPositionsResp{Items: items, Total: int32(res.Total)}, nil
}

func (s *Svc岗位管理) UpdateEncBossId(ctx context.Context, req *pb.UpdateEncBossIdReq) (*pb.PositionResp, error) {
	if ebz := s.uc.Xqt更新招聘者号(ctx, &biz.Req更新招聘者号{ID: uint(req.Id), E招聘者号: req.EncBossId}); ebz != nil {
		return nil, ebz.Erk
	}
	v岗位, ebz := s.uc.Get获取岗位(ctx, uint(req.Id))
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toPositionResp(v岗位), nil
}

func (s *Svc岗位管理) SyncPosition(ctx context.Context, req *pb.SyncPositionReq) (*pb.SyncPositionResp, error) {
	// strict validation — all fields must be present
	if len(req.JobId) != 28 {
		return nil, pb.ErrorInvalidJobId("jobId must be 28 chars, got %d", len(req.JobId))
	}
	if req.Title == "" {
		return nil, pb.ErrorMissingField("title is required")
	}
	if req.Company == "" {
		return nil, pb.ErrorMissingField("company is required")
	}
	if req.SalaryRange == "" {
		return nil, pb.ErrorMissingField("salaryRange is required")
	}
	if req.City == "" {
		return nil, pb.ErrorMissingField("city is required")
	}
	if req.Recruiter == "" {
		return nil, pb.ErrorMissingField("recruiter is required")
	}
	if req.Status == 0 {
		return nil, pb.ErrorMissingField("status is required")
	}
	if req.Duties == "" {
		return nil, pb.ErrorMissingField("duties is required")
	}
	if req.Requirements == "" {
		return nil, pb.ErrorMissingField("requirements is required")
	}
	if len(req.MatchItems) == 0 {
		return nil, pb.ErrorMissingField("matchItems is required and must not be empty")
	}
	if len(req.ChatMessages) == 0 {
		return nil, pb.ErrorMissingField("chatMessages is required — this endpoint is designed for syncing from chat page")
	} else if req.EncBossId == "" {
		return nil, pb.ErrorMissingField("encBossId is required when chatMessages is present — chat requires encBossId")
	}

	// upsert position: check if exists, create or update
	reqBiz := &biz.Req创建岗位{
		J岗位编号: req.JobId,
		T岗位名称: req.Title,
		C公司名称: req.Company,
		S薪资范围: req.SalaryRange,
		S薪资下限: req.SalaryMin,
		S薪资上限: req.SalaryMax,
		C城市名称: req.City,
		L岗位链接: req.Link,
		R招聘者:  req.Recruiter,
		E招聘者号: req.EncBossId,
		I猎头标记: req.IsHunter,
		S岗位状态: enums.Enum岗位状态映射表.GetByCode(req.Status).Basic(),
		D岗位职责: req.Duties,
		R岗位要求: req.Requirements,
		N备注信息: req.Notes,
	}

	v岗位, ebz := s.uc.Get按编号查(ctx, req.JobId)
	if ebz != nil {
		return nil, ebz.Erk
	}
	if v岗位 == nil {
		v岗位, ebz = s.uc.Xqt创建岗位(ctx, reqBiz)
		if ebz != nil {
			return nil, ebz.Erk
		}
	} else {
		if ebz := s.uc.Xqt全量更新(ctx, v岗位.ID, reqBiz); ebz != nil {
			return nil, ebz.Erk
		}
	}

	// replace matchItems (full replacement)
	matchItems := make([]*biz.Req需求匹配项, 0, len(req.MatchItems))
	for _, item := range req.MatchItems {
		matchItems = append(matchItems, &biz.Req需求匹配项{
			R岗位要求: item.Requirement,
			M匹配状态: enums.Enum匹配状态映射表.GetByCode(item.MatchStatus).Basic(),
			R简历对应: item.ResumePoint,
			R补充说明: item.Remark,
			S排序序号: item.SortIndex,
		})
	}
	if _, ebz := s.ucMatch.Xqt批量设置(ctx, req.JobId, matchItems); ebz != nil {
		return nil, ebz.Erk
	}

	// replace chat messages (full replacement)
	chatItems := make([]*biz.Req聊天消息, 0, len(req.ChatMessages))
	for _, msg := range req.ChatMessages {
		chatItems = append(chatItems, &biz.Req聊天消息{
			D消息方向: msg.Direction,
			C消息内容: msg.Content,
			T消息时间: msg.Timestamp,
			B简历消息: msg.IsResume,
			R简历版本: msg.ResumeVersion,
		})
	}
	chatRes, ebz := s.ucChat.Xqt同步聊天(ctx, req.JobId, chatItems)
	if ebz != nil {
		return nil, ebz.Erk
	}

	return &pb.SyncPositionResp{
		Id:                uint64(v岗位.ID),
		JobId:             req.JobId,
		MatchItemsCount:   int32(len(matchItems)),
		ChatMessagesCount: chatRes.N消息数量,
		LastCommAt:        chatRes.L最后沟通,
		LastCommDir:       chatRes.L最后方向,
		LastResume:        chatRes.L简历版本,
	}, nil
}
