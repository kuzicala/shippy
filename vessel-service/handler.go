package main

import (
	"context"
	"gopkg.in/mgo.v2"
	"shippy/vessel-service/proto/vessel"
)

type service struct {
	session *mgo.Session
}

func (s *service) FindAvailable(ctx context.Context, req *go_micro_srv_vessel.Specification, res *go_micro_srv_vessel.Response) error {
	defer s.GetRepo().Close()
	vessel, err := s.GetRepo().FindAvailable(req)

	if err != nil {
		return err
	}

	res.Vessel = vessel

	return nil
}

func (s *service) Create(ctx context.Context, req *go_micro_srv_vessel.Vessel, res *go_micro_srv_vessel.Response) error {
	defer s.GetRepo().Close()
	if err := s.GetRepo().Create(req); err != nil {
		return err
	}

	res.Vessel = req
	res.Created = true
	return nil
}

func (s *service) GetRepo() *VesselRepository {
	return &VesselRepository{s.session.Clone()}
}
