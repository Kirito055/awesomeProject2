package main

import (
	"animeteam.com/avg"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	avg.UnimplementedAvgServiceServer
}

func (s *Server) Avg(req *avg.Request, stream avg.AvgService_AvgServer) error {
	number := req.GetNumber()
	var slice []int = FindAvg(double(number))

	for i := 0; i < len(slice); i++ {
		res := &avg.Response{Avg: int64(slice[i])}
		if err := stream.Send(res); err != nil {
			log.Fatalf("errir")
		}
		time.Sleep(time.Second)
	}
	return nil
}

func FindAvg(a []int) double {
	sum:=0
	for i := 0; i < len(a); i++{
		sum=sum+a[i]

	}
	return sum/len(a)
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:5555")
	if err != nil {
		log.Fatal("Failed to listen :%v", err)
	}
	s := grpc.NewServer()
	avg.RegisterAvgServiceServer(s, &Server{})
	log.Println("Server is running on port 5555 :")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

//Greet is an example of unary rpc call
