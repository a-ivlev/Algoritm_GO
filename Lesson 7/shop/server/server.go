package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	pb "shop/server/api/proto"
	"sort"
	"sync"
	"time"
)

type Item struct {
	ID        int32
	Name      string
	Price     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ItemRepositoryService struct {
	mu          *sync.RWMutex
	itemsTable  *itemsTable

	pb.UnimplementedItemRepositoryServer
}

type itemsTable struct {
	items map[int32]*Item
	maxID int32
}

func (s *ItemRepositoryService) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.Item, error) {
	s.itemsTable.maxID++

	timeNow := time.Now().UTC()

	newItem := &Item{
		ID:        s.itemsTable.maxID,
		Price:     req.Price,
		Name:      req.Name,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	s.mu.Lock()
	s.itemsTable.items[newItem.ID] = newItem
	s.mu.Unlock()

	return &pb.Item{
		Id:        newItem.ID,
		Name:      newItem.Name,
		Price:     newItem.Price,
	}, nil
}

func (s *ItemRepositoryService) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.Item, error) {
	s.mu.RLock()
	updateItem, ok := s.itemsTable.items[req.Id]
	s.mu.RUnlock()
	if !ok {
		return nil, status.Error(codes.NotFound, "item not found")
	}
	updateItem.Name = req.Name
	updateItem.Price = req.Price
	updateItem.UpdatedAt = time.Now().UTC()

	return &pb.Item{
		Id:        updateItem.ID,
		Name:      updateItem.Name,
		Price:     updateItem.Price,
	}, nil
}

func (s *ItemRepositoryService) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.Message, error) {
	s.mu.RLock()
	_, ok := s.itemsTable.items[req.Id]
	s.mu.RUnlock()
	if !ok {
		return &pb.Message{Text: "item not found"}, status.Error(codes.NotFound, "item not found")
	}
	s.mu.Lock()
	delete(s.itemsTable.items, req.Id)
	s.mu.Unlock()
	return &pb.Message{Text: "item deleted"}, nil
}

func (s *ItemRepositoryService) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.Item, error) {
	s.mu.RLock()
	item, ok := s.itemsTable.items[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "item not found")
	}
	s.mu.RUnlock()

	return &pb.Item{
		Id:        item.ID,
		Name:      item.Name,
		Price:     item.Price,
	}, nil
}

func (s *ItemRepositoryService) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
	s.mu.RLock()
	itemSlice := make([]*Item, 0, len(s.itemsTable.items))
	for _, item := range s.itemsTable.items {
		itemSlice = append(itemSlice, item)
	}
	s.mu.RUnlock()
	sort.Slice(itemSlice, func(i, j int) bool {
		return itemSlice[i].ID < itemSlice[j].ID
	})

	resFiltered := make([]*pb.Item, 0, len(itemSlice))
	for idx, item := range itemSlice {
		if int32(len(resFiltered)) == req.Limit {
			break
		}
		if int32(idx) < req.Offset {
			continue
		}

		itemPb := &pb.Item{
			Id: item.ID,
			Name: item.Name,
			Price: item.Price,
		}

		resFiltered = append(resFiltered, itemPb)
	}
	return &pb.ListItemsResponse{Items: resFiltered}, nil
}

func NewMapDB() pb.ItemRepositoryServer {
	return &ItemRepositoryService{
		mu: &sync.RWMutex{},
		itemsTable: &itemsTable{
			items: map[int32]*Item{},//initMapDBitems,
			maxID: 8,
		},
	}
}

func NewItemRepositoryServerStart(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := NewMapDB()
	serv := grpc.NewServer()

	log.Println("starting grpc server at", addr)

	pb.RegisterItemRepositoryServer(serv, s)
	if err = serv.Serve(lis); err != nil {
		return err
	}
	return nil
}

func main()  {
	NewItemRepositoryServerStart("localhost:9094")
}