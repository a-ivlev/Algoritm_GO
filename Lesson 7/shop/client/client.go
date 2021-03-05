package main

import (
	pb "shop/client/api/proto"
	"context"
	"log"

	"google.golang.org/grpc"
)

func NewItemRepositoryClient(addr string) (pb.ItemRepositoryClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return pb.NewItemRepositoryClient(conn), nil
}

func main()  {
	itemRepository, err := NewItemRepositoryClient("localhost:9094")
	if err != nil {
		log.Fatal(err)
	}

	createItemReq := &pb.CreateItemRequest{
		Name: "test item 1",
		Price: 25000000,
	}
	item, err := itemRepository.CreateItem(context.Background(), createItemReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("item created", *item)

	updateItemReq := &pb.UpdateItemRequest{
		Id: 1,
		Name: "test item update",
		Price: 25500000,
	}
	updItem, err := itemRepository.UpdateItem(context.Background(), updateItemReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("item 1 updated", *updItem)

	listItemsReq := &pb.ListItemsRequest{
		Limit: 1,
	}
	resp, err := itemRepository.ListItems(context.Background(), listItemsReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("list items", resp.Items[0])
}
