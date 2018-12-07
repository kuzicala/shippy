package main

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	pb "shippy/consignment-service/proto/consignment"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

//服务

type service struct {
	repo Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {

	server := micro.NewService(micro.Name("go_micro_srv_consignment"), micro.Version("latest"))
	server.Init()
	repo := Repository{}
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo: repo})

	if e := server.Run(); e != nil {
		log.Fatalf("failed to serve: %v", e)
	}
}
