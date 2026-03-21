package service

import (
	"context"

	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/biz"
)

type Svc沟通管理 struct {
	pb.UnimplementedCommunicationServiceServer

	uc *biz.Uc沟通管理
}

func NewSvc沟通管理(uc *biz.Uc沟通管理) *Svc沟通管理 {
	return &Svc沟通管理{uc: uc}
}

func (s *Svc沟通管理) SyncChat(ctx context.Context, req *pb.SyncChatReq) (*pb.SyncChatResp, error) {
	items := make([]*biz.Req聊天消息, 0, len(req.Messages))
	for _, msg := range req.Messages {
		items = append(items, &biz.Req聊天消息{
			D消息方向: msg.Direction,
			C消息内容: msg.Content,
			T消息时间: msg.Timestamp,
			B简历消息: msg.IsResume,
			R简历版本: msg.ResumeVersion,
		})
	}
	res, ebz := s.uc.Xqt同步聊天(ctx, uint(req.PositionId), items)
	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.SyncChatResp{
		LastCommAt:   res.L最后沟通,
		LastCommDir:  res.L最后方向,
		LastResume:   res.L简历版本,
		MessageCount: res.N消息数量,
	}, nil
}

func (s *Svc沟通管理) GetChat(ctx context.Context, req *pb.GetChatReq) (*pb.GetChatResp, error) {
	v记录们, ebz := s.uc.Get聊天记录(ctx, uint(req.PositionId))
	if ebz != nil {
		return nil, ebz.Erk
	}
	messages := make([]*pb.ChatMessage, 0, len(v记录们))
	for _, v := range v记录们 {
		messages = append(messages, &pb.ChatMessage{
			Direction:     v.D消息方向,
			Content:       v.C消息内容,
			Timestamp:     v.T消息时间,
			IsResume:      v.B简历消息,
			ResumeVersion: v.R简历版本,
		})
	}
	return &pb.GetChatResp{Messages: messages}, nil
}
