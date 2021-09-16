package main

import (
	"context"
	"fmt"
	"grpcexp/server"
	"grpcexp/src/exp/proto/exp"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	go runServer()

	time.Sleep(time.Second * 1)
	conn, err := grpc.Dial("localhost:50045", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := exp.NewExpServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := client.GetExp(ctx, &exp.GetExpReq{
		Count: 3,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, ee := range res.Exps {
		fmt.Println(ee.Name)
	}
	time.Sleep(time.Second)
}

func runServer() {
	lis, err := net.Listen("tcp", ":50045")
	if err != nil {
		log.Fatal(err)
	}
	e := &server.ExpServ{}
	serv := grpc.NewServer()
	exp.RegisterExpServiceServer(serv, e)
	serv.Serve(lis)
}
