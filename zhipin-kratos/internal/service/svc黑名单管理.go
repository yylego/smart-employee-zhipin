package service

import (
	"context"

	pb "github.com/yylego/smart-employee-zhipin/zhipin-kratos/api/zhipin"
	"github.com/yylego/smart-employee-zhipin/zhipin-kratos/internal/biz"
)

type Svc黑名单管理 struct {
	pb.UnimplementedBlacklistServiceServer

	uc *biz.Uc黑名单管理
}

func NewSvc黑名单管理(uc *biz.Uc黑名单管理) *Svc黑名单管理 {
	return &Svc黑名单管理{uc: uc}
}

func (s *Svc黑名单管理) AddBlacklist(ctx context.Context, req *pb.AddBlacklistReq) (*pb.BlacklistResp, error) {
	v, ebz := s.uc.Xqt添加黑名单(ctx, &biz.Req添加黑名单{C公司名称: req.Company, R拉黑原因: req.Reason})
	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.BlacklistResp{Id: uint64(v.ID), Company: v.C公司名称, Reason: v.R拉黑原因}, nil
}

func (s *Svc黑名单管理) CheckBlacklist(ctx context.Context, req *pb.CheckBlacklistReq) (*pb.CheckBlacklistResp, error) {
	return &pb.CheckBlacklistResp{Blacklisted: s.uc.Get检查黑名单(ctx, req.Company)}, nil
}

func (s *Svc黑名单管理) ListBlacklist(ctx context.Context, req *pb.ListBlacklistReq) (*pb.ListBlacklistResp, error) {
	v黑名单们, ebz := s.uc.Get黑名单列表(ctx)
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.BlacklistResp, 0, len(v黑名单们))
	for _, v := range v黑名单们 {
		items = append(items, &pb.BlacklistResp{Id: uint64(v.ID), Company: v.C公司名称, Reason: v.R拉黑原因})
	}
	return &pb.ListBlacklistResp{Items: items}, nil
}

func (s *Svc黑名单管理) RemoveBlacklist(ctx context.Context, req *pb.RemoveBlacklistReq) (*pb.RemoveBlacklistResp, error) {
	if ebz := s.uc.Xqt移除黑名单(ctx, uint(req.Id)); ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.RemoveBlacklistResp{}, nil
}
