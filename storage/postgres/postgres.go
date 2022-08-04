package postgres

import (
	"context"
	pb "crud_service/genproto/crud_service"
)

type CrudI interface {
	GetAll(ctx context.Context, req *pb.ListReq) (*pb.ListRespPost, error)
	Get(ctx context.Context, req *pb.ById) (*pb.Post, error)
	Update(ctx context.Context, req *pb.Post) error
	Delete(ctx context.Context, req *pb.ById) error
}
