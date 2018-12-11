package main

import (
	"context"
	"gopkg.in/mgo.v2"
	"log"
	"shippy/consignment-service/proto/consignment"
	vesselProto "shippy/vessel-service/proto/vessel"
)

type handler struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

func (s *handler) CreateConsignment(ctx context.Context, req *go_micro_srv_consignment.Consignment, resp *go_micro_srv_consignment.Response) error {
	defer s.GetRepo().Close()
	spec := &vesselProto.Specification{Capacity: int32(len(req.Containers)), MaxWeight: req.Weight}

	response, e := s.vesselClient.FindAvailable(ctx, spec)
	if e != nil {
		return e
	}
	log.Printf("fount vessel: %s \n", response.Vessel.Name)
	req.VesselId = response.Vessel.Id

	e = s.GetRepo().Create(req)
	if e != nil {
		return e
	}

	resp.Created = true
	resp.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *go_micro_srv_consignment.GetRequest, resp *go_micro_srv_consignment.Response) error {
	defer s.GetRepo().Close()
	consignments, e := s.GetRepo().GetAll()

	if e != nil {
		return e
	}
	resp.Consignments = consignments

	return nil
}

func (s *handler) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}
