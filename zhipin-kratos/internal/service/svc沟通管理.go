package service

import (
	"context"

	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/biz"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/enums"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/pkg/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Svc沟通管理 struct {
	pb.UnimplementedCommunicationServiceServer

	uc *biz.Uc沟通管理
}

func NewSvc沟通管理(uc *biz.Uc沟通管理) *Svc沟通管理 {
	return &Svc沟通管理{uc: uc}
}

func (s *Svc沟通管理) RecordMsgSent(ctx context.Context, req *pb.RecordMsgSentReq) (*pb.CommunicationResp, error) {
	v, ebz := s.uc.Xqt记录事件(ctx, &biz.Req记录事件{P岗位主键: uint(req.PositionId), E事件类型: models.C事件类型_发消息, C消息内容: req.Content, D消息方向: 0})
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toCommunicationResp(v), nil
}

func (s *Svc沟通管理) RecordMsgReceived(ctx context.Context, req *pb.RecordMsgReceivedReq) (*pb.CommunicationResp, error) {
	v, ebz := s.uc.Xqt记录事件(ctx, &biz.Req记录事件{P岗位主键: uint(req.PositionId), E事件类型: models.C事件类型_收消息, C消息内容: req.Content, D消息方向: 1})
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toCommunicationResp(v), nil
}

func (s *Svc沟通管理) RecordResumeSent(ctx context.Context, req *pb.RecordResumeSentReq) (*pb.CommunicationResp, error) {
	v, ebz := s.uc.Xqt记录事件(ctx, &biz.Req记录事件{P岗位主键: uint(req.PositionId), E事件类型: models.C事件类型_发简历, C消息内容: req.ResumeVersion, D消息方向: 0})
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toCommunicationResp(v), nil
}

func (s *Svc沟通管理) RecordChatLimited(ctx context.Context, req *pb.RecordChatLimitedReq) (*pb.CommunicationResp, error) {
	v, ebz := s.uc.Xqt记录事件(ctx, &biz.Req记录事件{P岗位主键: uint(req.PositionId), E事件类型: models.C事件类型_开聊限制, D消息方向: 0})
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toCommunicationResp(v), nil
}

func (s *Svc沟通管理) RecordInterview(ctx context.Context, req *pb.RecordInterviewReq) (*pb.CommunicationResp, error) {
	v, ebz := s.uc.Xqt记录事件(ctx, &biz.Req记录事件{P岗位主键: uint(req.PositionId), E事件类型: models.C事件类型_安排面试, C消息内容: req.Notes, D消息方向: 0})
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toCommunicationResp(v), nil
}

func (s *Svc沟通管理) RecordRejection(ctx context.Context, req *pb.RecordRejectionReq) (*pb.CommunicationResp, error) {
	v, ebz := s.uc.Xqt记录事件(ctx, &biz.Req记录事件{P岗位主键: uint(req.PositionId), E事件类型: models.C事件类型_被拒绝, C消息内容: req.Content, D消息方向: 1})
	if ebz != nil {
		return nil, ebz.Erk
	}
	return toCommunicationResp(v), nil
}

func (s *Svc沟通管理) ListCommunications(ctx context.Context, req *pb.ListCommunicationsReq) (*pb.ListCommunicationsResp, error) {
	v记录们, ebz := s.uc.Get沟通列表(ctx, uint(req.PositionId))
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.CommunicationResp, 0, len(v记录们))
	for _, v := range v记录们 {
		items = append(items, toCommunicationResp(v))
	}
	return &pb.ListCommunicationsResp{Items: items}, nil
}

func toCommunicationResp(v *models.T沟通记录) *pb.CommunicationResp {
	return &pb.CommunicationResp{
		Id:         uint64(v.ID),
		PositionId: uint64(v.P岗位主键),
		EventType:  int32(enums.Enum事件类型映射表.MustGetByBasic(v.E事件类型).Proto()),
		EventTime:  timestamppb.New(v.E事件时间),
		Content:    v.C消息内容,
		Direction:  v.D消息方向,
	}
}
