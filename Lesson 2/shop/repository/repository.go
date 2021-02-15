package repository

import (
	"errors"
	"sort"
	"sync"
	"time"

	"shop/models"
)

var (
	ErrNotFound = errors.New("not found")
)

type Repository interface {
	CreateItem(item *models.Item) (*models.Item, error)
	UpdateItem(item *models.Item) (*models.Item, error)
	DeleteItem(itemID int32) error
	GetItem(itemID int32) (*models.Item, error)
	ListItems(filter *ItemFilter) ([]*models.Item, error)

	CreateOrder(order *models.Order) (*models.Order, error)
	ListOrders(filter *OrderFilter) ([]*models.Order, error)
}

type mapDB struct {
	mu          *sync.RWMutex
	itemsTable  *itemsTable
	ordersTable *ordersTable
}

type ordersTable struct {
	orders map[int32]*models.Order
	maxID  int32
}

type itemsTable struct {
	items map[int32]*models.Item
	maxID int32
}

func NewMapDB() Repository {
	return &mapDB{
		mu: &sync.RWMutex{},
		itemsTable: &itemsTable{
			items: initMapDBitems,
			maxID: 8,
		},
		ordersTable: &ordersTable{
			orders: make(map[int32]*models.Order),
			maxID:  0,
		},
	}
}

func (m *mapDB) CreateOrder(order *models.Order) (*models.Order, error) {
	m.ordersTable.maxID++

	timeNow := time.Now().UTC()

	newOrder := &models.Order{
		ID:            m.ordersTable.maxID,
		CustomerName:  order.CustomerName,
		CustomerPhone: order.CustomerPhone,
		CustomerEmail: order.CustomerEmail,
		ItemIDs:       order.ItemIDs,
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
	}

	m.mu.Lock()
	m.ordersTable.orders[newOrder.ID] = newOrder
	m.mu.Unlock()

	return &models.Order{
		ID:            newOrder.ID,
		CustomerName:  newOrder.CustomerName,
		CustomerPhone: newOrder.CustomerPhone,
		CustomerEmail: newOrder.CustomerEmail,
		ItemIDs:       newOrder.ItemIDs,
		CreatedAt:     newOrder.CreatedAt,
		UpdatedAt:     newOrder.UpdatedAt,
	}, nil
}

func (m *mapDB) ListOrders(filter *OrderFilter) ([]*models.Order, error) {
	var res []*models.Order

	m.mu.RLock()
	orderSlice := make([]*models.Order, 0, len(m.ordersTable.orders))
	for _, order := range m.ordersTable.orders {
		orderSlice = append(orderSlice, order)
	}
	m.mu.RUnlock()

	sort.Slice(orderSlice, func(i, j int) bool {
		return orderSlice[i].ID < orderSlice[j].ID
	})

	resFiltered := orderSlice
	for idx, order := range res {
		if len(resFiltered) == filter.Limit {
			break
		}
		if idx < filter.Offset {
			continue
		}
		resFiltered = append(resFiltered, order)
	}
	return resFiltered, nil
}

func (m *mapDB) CreateItem(item *models.Item) (*models.Item, error) {
	m.itemsTable.maxID++

	timeNow := time.Now().UTC()

	newItem := &models.Item{
		ID:        m.itemsTable.maxID,
		Price:     item.Price,
		Name:      item.Name,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	m.mu.Lock()
	m.itemsTable.items[newItem.ID] = newItem
	m.mu.Unlock()

	return &models.Item{
		ID:        newItem.ID,
		Name:      newItem.Name,
		Price:     newItem.Price,
		CreatedAt: newItem.CreatedAt,
		UpdatedAt: newItem.UpdatedAt,
	}, nil
}

func (m *mapDB) ListItems(filter *ItemFilter) ([]*models.Item, error) {
	var res []*models.Item

	m.mu.RLock()
	itemSlice := make([]*models.Item, 0, len(m.itemsTable.items))
	for _, item := range m.itemsTable.items {
		itemSlice = append(itemSlice, item)
	}
	m.mu.RUnlock()
	sort.Slice(itemSlice, func(i, j int) bool {
		return itemSlice[i].ID < itemSlice[j].ID
	})

	if filter.PriceLeft == nil && filter.PriceRight == nil {
		res = itemSlice
	} else {
		for _, item := range itemSlice {
			switch {
			case filter.PriceLeft != nil && filter.PriceRight == nil:
				if item.Price >= *filter.PriceLeft {
					res = append(res, item)
				}
			case filter.PriceRight != nil && filter.PriceLeft == nil:
				if item.Price <= *filter.PriceRight {
					res = append(res, item)
				}
			case filter.PriceRight != nil && filter.PriceRight != nil:
				if item.Price <= *filter.PriceRight && item.Price >= *filter.PriceLeft {
					res = append(res, item)
				}
			}
		}
	}

	resFiltered := make([]*models.Item, 0, len(res))
	for idx, item := range res {
		if len(resFiltered) == filter.Limit {
			break
		}
		if idx < filter.Offset {
			continue
		}
		resFiltered = append(resFiltered, item)
	}
	return resFiltered, nil
}

func (m *mapDB) GetItem(ID int32) (*models.Item, error) {
	m.mu.RLock()
	item, ok := m.itemsTable.items[ID]
	if !ok {
		return nil, ErrNotFound
	}
	m.mu.RUnlock()

	return &models.Item{
		ID:        item.ID,
		Name:      item.Name,
		Price:     item.Price,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}, nil
}

func (m *mapDB) DeleteItem(ID int32) error {
	m.mu.RLock()
	_, ok := m.itemsTable.items[ID]
	m.mu.RUnlock()
	if !ok {
		return ErrNotFound
	}
	m.mu.Lock()
	delete(m.itemsTable.items, ID)
	m.mu.Unlock()
	return nil
}

func (m *mapDB) UpdateItem(item *models.Item) (*models.Item, error) {
	m.mu.RLock()
	updateItem, ok := m.itemsTable.items[item.ID]
	m.mu.RUnlock()
	if !ok {
		return nil, ErrNotFound
	}
	updateItem.Name = item.Name
	updateItem.Price = item.Price
	updateItem.UpdatedAt = time.Now().UTC()

	return &models.Item{
		ID:        updateItem.ID,
		Name:      updateItem.Name,
		Price:     updateItem.Price,
		CreatedAt: updateItem.CreatedAt,
		UpdatedAt: updateItem.UpdatedAt,
	}, nil
}
