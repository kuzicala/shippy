package main

import (
	"context"
	"github.com/kuzicala/shippy/user-service/proto/user"
)

type handler struct {
	repo         Repository
	tokenService Authable
}

func (h *handler) Create(ctx context.Context, req *go_micro_srv_user.User, resp *go_micro_srv_user.Response) error {
	if err := h.repo.Create(req); err != nil {
		return err
	}

	resp.User = req
	return nil
}

func (h *handler) Get(ctx context.Context, req *go_micro_srv_user.User, resp *go_micro_srv_user.Response) error {
	user, e := h.repo.Get(req.Id)
	if e != nil {
		return e
	}

	resp.User = user
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *go_micro_srv_user.Request, resp *go_micro_srv_user.Response) error {
	users, e := h.repo.GetAll()

	if e != nil {
		return e
	}

	resp.Users = users
	return nil
}

func (h *handler) Auth(ctx context.Context, req *go_micro_srv_user.User, resp *go_micro_srv_user.Token) error {
	_, e := h.repo.GetByEmailAndPassword(req)
	if e != nil {
		return e
	}

	resp.Token = "testingabc"
	return nil
}

func (handler) ValidateToken(context.Context, *go_micro_srv_user.Token, *go_micro_srv_user.Token) error {
	return nil
}
