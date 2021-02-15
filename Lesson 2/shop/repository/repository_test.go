package repository

import (
	"gb-go-architecture/lesson-1/shop_new/models"
	"testing"
)

func TestCreateItem(t *testing.T) {
	db := NewMapDB()

	input := &models.Item{
		Name:  "someName",
		Price: 10,
	}
	expected := &models.Item{
		ID:    1,
		Name:  input.Name,
		Price: input.Price,
	}

	result, err := db.CreateItem(input)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}

	result, err = db.GetItem(expected.ID)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}

	input = &models.Item{
		Name:  "someName2",
		Price: 20,
	}
	expected = &models.Item{
		ID:    2,
		Name:  input.Name,
		Price: input.Price,
	}

	result, err = db.CreateItem(input)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}
}
