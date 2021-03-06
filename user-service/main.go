package main

import (
	pb "github.com/kuzicala/shippy/user-service/proto/user"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	db, e := CreateConnection()
	defer db.Close()
	if e != nil {
		log.Fatalf("Could not connect to DB:%v", e)
	}
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db: db}

	tokenService := &TokenService{repo}

	service := micro.NewService(micro.Name("go_micro_srv_user"), micro.Version("latest"))

	service.Init()

	pb.RegisterUserServiceHandler(service.Server(), &handler{repo, tokenService})

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
