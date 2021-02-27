package main

import (

	"animeteam.com/avg"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:5555", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := avg.NewAvgServiceClient(conn)
	PrimeNumberDecomposition(c)
}

func PrimeNumberDecomposition(c avg.AvgServiceClient) {

	req := &avg.Request{
		Number:120,
	}
	stream,err:=c.Avg(context.Background(),req)
	if err!=nil{
		log.Fatalf("error")
	}
	go func() {
			log.Printf("Send %v",req)
			err:=stream.SendMsg(req)
			if err!=nil{
				log.Fatalf("error ")

			}
			time.Sleep(time.Second)

	}()

}
