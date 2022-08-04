package service

import (
	"context"
	"crud_service/pkg/helper"
	"crud_service/pkg/logger"
	"crud_service/storage"

	pb "crud_service/genproto/crud_service"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type crudService struct {
	logger  logger.Logger
	storage storage.StorageI
	pb.UnimplementedCrudServiceServer
}

func NewCrudService(db *sqlx.DB, log logger.Logger) *crudService {
	return &crudService{
		logger:  log,
		storage: storage.NewStoragePg(db),
	}
}

func (s *crudService) GetAll(ctx context.Context, req *pb.ListReq) (*pb.ListRespPost, error) {
	s.logger.Info("---GetAll Posts--->", logger.Any("req", req))

	resp, err := s.storage.Crud().GetAll(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "Error GetAll Post", req, codes.Internal)
	}
	s.logger.Info("---GetAll Post--->", logger.Any("resp", resp))
	return resp, nil
}

func (s *crudService) Get(ctx context.Context, req *pb.ById) (*pb.Post, error) {
	s.logger.Info("---Get Post--->", logger.Any("req", req))

	resp, err := s.storage.Crud().Get(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "Error get Post", req, codes.Internal)
	}

	s.logger.Info("---Get Post--->", logger.Any("resp", resp))
	return resp, nil
}

func (s *crudService) Update(ctx context.Context, req *pb.Post) (*emptypb.Empty, error) {
	s.logger.Info("---Update Post--->", logger.Any("req", req))

	err := s.storage.Crud().Update(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "Error update Post", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}

func (s *crudService) Delete(ctx context.Context, req *pb.ById) (*emptypb.Empty, error) {
	s.logger.Info("---Delete Post--->", logger.Any("req", req))

	err := s.storage.Crud().Delete(ctx, req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "Error Delete Post", req, codes.Internal)
	}

	return &emptypb.Empty{}, nil
}
