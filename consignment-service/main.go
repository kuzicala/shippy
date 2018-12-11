package main

import (
	"github.com/micro/go-micro"
	"log"
	"os"
	pb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	var host string
	if host = os.Getenv("DB_HOST"); host == "" {
		host = defaultHost
	}

	session, e := CreateSession(host)
	if e != nil {
		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, e)
	}
	server := micro.NewService(micro.Name("go_micro_srv_consignment"), micro.Version("latest"))
	server.Init()
	vesselClient := vesselPb.NewVesselServiceClient("go_micro_srv_vessel", server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &handler{session: session, vesselClient: vesselClient})

	if e := server.Run(); e != nil {
		log.Fatalf("failed to serve: %v", e)
	}
}
