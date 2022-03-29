package service

import (
	"context"
	"kratos-project/app/policy/internal/biz"

	pb "kratos-project/api/policy/v1"
)

type PolicyService struct {
	pu  *biz.PolicyUsecase
	ipu *biz.ImportPolicyUsecase
	pb.UnimplementedPolicyServer
}

func NewPolicyService(pu *biz.PolicyUsecase, ipu *biz.ImportPolicyUsecase) *PolicyService {
	return &PolicyService{
		pu:  pu,
		ipu: ipu,
	}
}

func (s *PolicyService) CreatePolicy(ctx context.Context, req *pb.PolicyStruct) (*pb.PolicyStruct, error) {
	return &pb.PolicyStruct{}, nil
}
func (s *PolicyService) GetPolicy(ctx context.Context, req *pb.GetPolicyRequest) (*pb.PolicyStruct, error) {
	return &pb.PolicyStruct{}, nil
}
func (s *PolicyService) ListPolicy(ctx context.Context, req *pb.ListPolicyRequest) (*pb.ListPolicyReply, error) {
	return &pb.ListPolicyReply{}, nil
}
func (s *PolicyService) UpdatePolicy(ctx context.Context, req *pb.PolicyStruct) (*pb.PolicyStruct, error) {
	return &pb.PolicyStruct{}, nil
}
func (s *PolicyService) DeletePolicy(ctx context.Context, req *pb.DeletePolicyRequest) (*pb.PolicyStruct, error) {
	return &pb.PolicyStruct{}, nil
}
