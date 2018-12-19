package main

import (
	"context"
	pb "github.com/kuzicala/shippy/user-service/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"log"
	"os"
)

func main() {
	cmd.Init()

	client := pb.NewUserServiceClient("go_micro_srv_user", microclient.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "You full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "You email",
			},

			cli.StringFlag{
				Name:  "password",
				Usage: "You password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "You company",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			emai := c.String("email")
			password := c.String("password")
			company := c.String("company")

			response, e := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    emai,
				Password: password,
				Company:  company,
			})

			if e != nil {
				log.Fatalf("Counld not create:%v", e)
			}

			log.Printf("Created:%t", response.User.Id)
			all, e := client.GetAll(context.Background(), &pb.Request{})
			if e != nil {
				log.Fatalf("Could not list users:%v", e)
			}

			for _, v := range all.Users {
				log.Println(v)
			}
			os.Exit(0)
		}),
	)

}
