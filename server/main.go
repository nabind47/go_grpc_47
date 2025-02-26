package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/nabind47/go_47/generated"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(r *pb.MenuRequest, srv grpc.ServerStreamingServer[pb.Menu]) error {
	items := []*pb.Item{
		{Id: "1", Name: "Espresso"},
		{Id: "2", Name: "Latte"},
		{Id: "3", Name: "Cappuccino"},
		{Id: "4", Name: "Americano"},
		{Id: "5", Name: "Mocha"},
		{Id: "6", Name: "Macchiato"},
		{Id: "7", Name: "Flat White"},
		{Id: "8", Name: "Affogato"},
		{Id: "9", Name: "Irish Coffee"},
		{Id: "10", Name: "Turkish Coffee"},
		{Id: "11", Name: "Cortado"},
		{Id: "12", Name: "Ristretto"},
		{Id: "13", Name: "Lungo"},
	}

	for i := range items {
		srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
		// time.Sleep(time.Second)
	}
	return nil
}

func (s *server) PlaceOrder(ctx context.Context, r *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{Id: "1"}, nil
}

func (s *server) GetOrderStatus(ctx context.Context, r *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{OrderId: "1", Status: "Progress"}, nil
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	resp, err := handler(ctx, req)
	log.Printf("Request - Method:%s Duration:%s Error:%v\n",
		info.FullMethod,
		time.Since(start),
		err)
	return resp, err
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
