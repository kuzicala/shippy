package main

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	pb "shippy/vessel-service/proto/vessel"
)

//实现服务端
func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	vessel, err := s.repo.FindAvailable(spec)

	if err != nil {
		return err
	}

	resp.Vessel = vessel

	return nil
}

func main() {
	//停留在港口的货船，先写死
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty Mcboatface", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{Vessels: vessels}

	newService := micro.NewService(micro.Name("go_micro_srv_vessel"), micro.Version("latest"))

	newService.Init()

	pb.RegisterVesselServiceHandler(newService.Server(), &service{repo})

	if err := newService.Run(); err != nil {
		log.Fatalf("fail to serve:%v", err)
	}
}
