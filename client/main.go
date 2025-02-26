package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/nabind47/go_47/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCoffeeShopClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	menuStream, err := c.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatalf("Failed to get menu: %v", err)
	}

	var items []*pb.Item
	for {
		resq, err := menuStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive menu: %v", err)
		}
		items = append(items, resq.Items...)
		log.Printf("Received items: %v", resq.Items)
	}

	receipt, err := c.PlaceOrder(ctx, &pb.Order{Items: items})
	if err != nil {
		log.Fatalf("Failed to place order: %v", err)
	}
	log.Printf("Order placed: %v", receipt)

	orderStatus, err := c.GetOrderStatus(ctx, &pb.Receipt{Id: receipt.Id})
	if err != nil {
		log.Fatalf("Failed to get order status: %v", err)
	}
	log.Printf("Order status: %v", orderStatus)
}
