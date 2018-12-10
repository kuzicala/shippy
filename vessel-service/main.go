package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	"log"
	pb "shippy/vessel-service/proto/vessel"
)

type Repository interface {
	FindAvailable(vessel *pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	Vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, v := range repo.Vessels {
		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
			return v, nil
		}
	}

	return nil, errors.New("no vessel can't be use")

}

type service struct {
	repo Repository
}

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
