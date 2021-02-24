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
}

type mapDB struct {
	mu         *sync.Mutex
	itemsTable *itemsTable
}

type itemsTable struct {
	items map[int32]*models.Item
	maxID int32
}

func NewMapDB() Repository {
	return &mapDB{
		itemsTable: &itemsTable{
			items: initMapDBitems,
			maxID: 8,
		},
	}
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

	m.itemsTable.items[newItem.ID] = newItem

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

	itemSlice := make([]*models.Item, 0, len(m.itemsTable.items))
	for _, item := range m.itemsTable.items {
		itemSlice = append(itemSlice, item)
	}
	sort.Slice(itemSlice, func(i, j int) bool {
		return itemSlice[i].ID < itemSlice[j].ID
	})

	for _, item := range itemSlice {
		if filter.PriceLeft == nil && filter.PriceRight == nil {
			res = itemSlice
			break
		}
		if filter.PriceLeft != nil && filter.PriceRight == nil && item.Price >= *filter.PriceLeft ||
			filter.PriceLeft == nil && filter.PriceRight != nil && item.Price <= *filter.PriceRight {
			res = append(res, item)
		}
		if filter.PriceLeft != nil && filter.PriceRight != nil &&
			item.Price >= *filter.PriceLeft  && item.Price <= *filter.PriceRight {
				res = append(res, item)
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
	//return res[filter.Offset : filter.Offset+filter.Limit - 1], nil
}

func (m *mapDB) GetItem(ID int32) (*models.Item, error) {
	item, ok := m.itemsTable.items[ID]
	if !ok {
		return nil, ErrNotFound
	}

	return &models.Item{
		ID:        item.ID,
		Name:      item.Name,
		Price:     item.Price,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}, nil
}

func (m *mapDB) DeleteItem(ID int32) error {
	_, ok := m.itemsTable.items[ID]
	if !ok {
		return ErrNotFound
	}

	delete(m.itemsTable.items, ID)
	return nil
}

func (m *mapDB) UpdateItem(item *models.Item) (*models.Item, error) {
	updateItem, ok := m.itemsTable.items[item.ID]
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