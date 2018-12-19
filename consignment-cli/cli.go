package main

import (
	"context"
	"encoding/json"
	pb "github.com/kuzicala/shippy/consignment-service/proto/consignment"
	"github.com/micro/go-micro"
	"io/ioutil"
	"log"
	"os"
)

const (
	ADDRESS           = "localhost:50051"
	DEFAULT_INFO_FILE = "consignment.json"
)

func pareseFile(filename string) (*pb.Consignment, error) {
	bytes, e := ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}

	var consignment *pb.Consignment

	e = json.Unmarshal(bytes, &consignment)
	if e != nil {
		return nil, e
	}

	return consignment, nil
}

func main() {

	service := micro.NewService(micro.Name("go_micro_srv_consignment"))
	service.Init()

	client := pb.NewShippingServiceClient("go_micro_srv_consignment", service.Client())

	// 在命令行中指定新的货物信息 json 文件
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	// 解析货物信息
	consignment, err := pareseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}

	log.Printf("created: %t", resp.Created)
	log.Printf("resp: %v", resp)

	response, e := client.GetConsignments(context.Background(), &pb.GetRequest{})

	if e != nil {
		log.Fatalf("failed to list consignments: %v", e)
	}

	for _, c := range response.Consignments {
		log.Printf("%+v", c)
	}

}
