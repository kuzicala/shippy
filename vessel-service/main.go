package main

import (
	"github.com/micro/go-micro"
	"log"
	"os"
	pb "shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func CreateDummyData(repository Repository)  {
	defer  repository.Close()

	//停留在港口的货船，先写死
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty Mcboatface", MaxWeight: 200000, Capacity: 500},
	}

	for _,v := range vessels{
		repository.Create(v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")

	if host == ""{
		host = defaultHost
	}

	session, e := CreateSession(host)
	defer session.Close()

	if e != nil{
		log.Fatalf("Error connecting to datastore:%v",e)
	}

	repo := &VesselRepository{session.Copy()}
	CreateDummyData(repo)

	newService := micro.NewService(micro.Name("go_micro_srv_vessel"), micro.Version("latest"))

	newService.Init()

	pb.RegisterVesselServiceHandler(newService.Server(), &service{session:session})

	if err := newService.Run(); err != nil {
		log.Fatalf("fail to serve:%v", err)
	}
}
