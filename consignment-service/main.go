package main

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	pb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto/vessel"
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
	repo IRepository
	//consignment-service 作为客户端调用 vessel-service的函数
	vesselClient vesselPb.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	//检查是或否有何时的货轮

	spec := &vesselPb.Specification{Capacity: int32(len(req.Containers)), MaxWeight: req.Weight}

	response, e := s.vesselClient.FindAvailable(context.Background(), spec)
	if e != nil {
		return e
	}

	log.Printf("fount vessel: %s \n", response.Vessel.Name)
	req.VesselId = response.Vessel.Id
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
	repo := &Repository{}
	vesselClient := vesselPb.NewVesselServiceClient("go_micro_srv_vessel",server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo: repo,vesselClient:vesselClient})

	if e := server.Run(); e != nil {
		log.Fatalf("failed to serve: %v", e)
	}
}
