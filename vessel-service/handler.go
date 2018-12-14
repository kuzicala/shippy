package main

import "gopkg.in/mgo.v2"

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() *VesselRepository {
	return &VesselRepository{s.session.Clone()}
}



