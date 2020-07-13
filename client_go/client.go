package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"pancake/maker/gen/api"
)

var client = getClient()

func getClient() api.PancakeBakerServiceClient {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return api.NewPancakeBakerServiceClient(conn)
}

func bakePancake(menu api.Pancake_Menu) (*api.BakeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &api.BakeRequest{
		Menu: menu,
	}
	r, err := client.Bake(ctx, req)
	return r, err
}

func getReport() (*api.ReportResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &api.ReportRequest{}
	r, err := client.Report(ctx, req)
	return r, err
}

func main() {
	res, _ := bakePancake(api.Pancake_CLASSIC)
	fmt.Printf("Chef %s served %s pancake.\n", res.Pancake.ChefName, res.Pancake.Menu)

	res, _ = bakePancake(api.Pancake_MIX_BERRY)
	fmt.Printf("Chef %s served %s pancake.\n", res.Pancake.ChefName, res.Pancake.Menu)

	res, _ = bakePancake(api.Pancake_CLASSIC)
	fmt.Printf("Chef %s served %s pancake.\n", res.Pancake.ChefName, res.Pancake.Menu)

	reportRes, _ := getReport()
	for _, bakeCount := range reportRes.Report.BakeCounts {
		fmt.Println(bakeCount.Menu, bakeCount.Count)
	}
}
