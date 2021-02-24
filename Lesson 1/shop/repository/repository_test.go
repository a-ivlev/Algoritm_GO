package repository

import (
	"shop/models"
	"testing"
)

func TestCreateItem(t *testing.T) {
	db := NewMapDB()
// Проверка добавления в базу позиции товара.
// Создадим позицию товара с названием someName и ценой 10 коппек.
	input := &models.Item{
		Name:  "someName",
		Price: 10,
	}
// Ожидаемый результат.
// В базе должена появиться новая позиция товар с ID=9, Name="someName", Price=10.
	expected := &models.Item{
		ID:    9,
		Name:  input.Name,
		Price: input.Price,
	}
// Вызовем функцию CreateItem и передадим в неё параметры указанные в input.
// Запишем полученный результат в переменную result и ошибку в err.
	result, err := db.CreateItem(input)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Сравним полученные результаты ID, Name и Price с ожидаемыми.
	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}
	// Вызовем функцию CreateItem и передадим в неё параметры указанные в input.
	// Запишем полученный результат в переменную result и ошибку в err.
	result, err = db.GetItem(expected.ID)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Сравним полученные результаты ID, Name и Price с ожидаемыми.
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
// Ожидаемый результат.
// В базе должена появиться новая позиция товар с ID=10, Name="someName2", Price=20.
	expected = &models.Item{
		ID:    10,
		Name:  input.Name,
		Price: input.Price,
	}
// Вызовем функцию CreateItem и передадим в неё параметры указанные в input.
// Запишем полученный результат в переменную result и ошибку в err.
	result, err = db.CreateItem(input)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Сравним полученные результаты ID, Name и Price с ожидаемыми.
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

func TestDeleteItem(t *testing.T) {
	db := NewMapDB()
// Проверка удаления определённой позиции товара.
// Удалим цену и название товара у позиции товара с ID=5.

// Создадим переменную inputID и запишем номер ID позиции которую хотим удалить.
	inputID := &models.Item{
		ID:  5,
	}
// Передаём номер ID в функцию DeleteItem и получаем из неё результат который запишем в переменную err.
// Вызовем функцию DeleteItem и передадим в неё ID позиции товара который хотим удалить.
// Запишем полученный результат в переменную err.
	err := db.DeleteItem(inputID.ID)
// Проверяем если err не равна nil, значит удаление не прошло и мы выводим ошибку.
// Если err == nil значит функция отработала корректно.
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Сделаем запрос позиции товара с ID=8 через функцию GetItem в базу.
	// Запишем полученный результат в переменную result и ошибку в err.
	result, err := db.GetItem(inputID.ID)
	// Проверяем что не получаем ошибку
	if err != ErrNotFound {
		t.Error("unexpected error: ", err)
	}
	// Проверяем что данная позиция отсутствует в базе.
	if result != nil {
		t.Error("unexpected error: ", result)
	}
}

func TestUpdateItem(t *testing.T) {
	db := NewMapDB()
// Проверка обновления определённой позиции товара.
// Обновим цену и название товара у позиции товара с ID=8.

// Установим новые значения имени и цены.
	inputUpdate := &models.Item{
		ID: 8,
		Name:  "Gigabyte X570 AORUS Elite",
		Price: 16000000,
	}
// Ожидаемый результат. У товара с ID=8 должны обновиться следующие значения
// ID=8, Name="GIGABYTE X570 AORUS ELITE" Price=16690000 на новые значения.
	expected := &models.Item{
		ID:    8,
		Name:  inputUpdate.Name,
		Price: inputUpdate.Price,
	}
	// Вызовем функцию UpdateItem и передадим в неё параметры указанные в inputUpdate.
	// Запишем полученный результат в переменную result и ошибку в err.
	result, err := db.UpdateItem(inputUpdate)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Сравним полученные результаты ID, Name и Price с ожидаемыми.
	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}
	// Сделаем запрос позиции товара с ID=8 через функцию GetItem в базу.
	// Запишем полученный результат в переменную result и ошибку в err.
	result, err = db.GetItem(expected.ID)
	// Проверим на наличи ошибок.
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Сравним полученные результаты ID, Name и Price с ожидаемыми.
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

func TestListItems(t *testing.T) {
	db := NewMapDB()
// Проверка условия Limit равный 3. Границы цен не установлены.
// PriceLeft == nil && PriceRight == nil
// Установим Limit=3 и посмотрим сколько у нас будет выведено в результате позиций товара.
// Limit п о умолчанию равен 5.

// Устанавливаем Limit равный 3.
// Это означает что мы получим только первые три позиции товаров.
	inputFiter := &ItemFilter{
		Limit : 3,
	}
// Создаём ожидаемый результат.
	expectedLimit := []*models.Item{
		1: &models.Item{
			ID:    1,
			Name:  "Intel i3 Core",
			Price: 1000,
		},
		2: &models.Item{
			ID:    2,
			Name:  "Intel i5 Core",
			Price: 2000,
		},
		3: &models.Item{
			ID:    3,
			Name:  "Intel i7 Core",
			Price: 3000,
		},
	}
// Вызовем функцию ListItems и передадим в неё параметры указанные в inputFiter.
// Запишем полученный результат в переменную result и ошибку в err.
	res, err := db.ListItems(inputFiter)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Инициализируем слайс и записываем результат в items.
	items := make([]*models.Item, 0, len(res))

	// Проитерируемся по полученному слайсу и сравним ID, Name и Price полученные с ожидаемыми.
	for i, result := range items {
		if expectedLimit[i].ID != result.ID {
			t.Errorf("unexpected name: expected %d result: %d", expectedLimit[i].ID, result.ID)
		}
		if expectedLimit[i].Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedLimit[i].Name, result.Name)
		}
		if expectedLimit[i].Price != result.Price {
			t.Errorf("unexpected name: expected %d result: %d", expectedLimit[i].Price, result.Price)
		}
	}// Если нет ошибок всё ОК, проверка условия пройдена.

/// Проверка условия, когда установлена нижняя граница цен.
// filter.PriceLeft != nil && filter.PriceRight == nil && item.Price >= *filter.PriceLeft
// Это означает что ниже этой границы товары в выборку попадать не должны.
// Верхний порог цен PriceRight и Limit при проверке данного условия мы не устанавливаем.
// Limit заданный по умолчанию равен 5, поэтому мы должны получить не более 5 позиций.

	// Устанавливаем нижнюю границу цен.
	var priceLeftInput int64 = 3000
	inputFiter = &ItemFilter{
		PriceLeft: &priceLeftInput,
	}
	// Создаём ожидаемый результат.
	expectedPriceLeft := []*models.Item {
		3:&models.Item{
			ID: 3,
			Name: "Intel i7 Core",
			Price: 3000,
		},
		4:&models.Item{
			ID: 4,
			Name: "GeForce GTX 1650",
			Price: 25000000,
		},
		5:&models.Item{
			ID: 5,
			Name: "GeForce GTX 1050 Ti",
			Price: 19000000,
		},
		6:&models.Item{
			ID: 6,
			Name: "GIGABYTE TRX40 AORUS PRO",
			Price: 36390000,
		},
		7:&models.Item{
			ID: 7,
			Name: "GIGABYTE Z490 AORUS ULTRA G2",
			Price: 30440000,
		},
	}
// Вызовем функцию ListItems и передадим в неё параметры указанные в priceLeftInput.
// Запишем полученный результат в переменную result и ошибку в err.
	res, err = db.ListItems(inputFiter)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Инициализируем слайс и записываем результат в items
	items = make([]*models.Item, 0, len(res))

	// Теперь проитерируемся по полученному слайсу и сравним ID, Name и Price полученные с ожидаемыми.
	for i, result := range items {
		if expectedPriceLeft[i].ID != result.ID {
			t.Errorf("unexpected name: expected %d result: %d", expectedPriceLeft[i].ID, result.ID)
		}
		if expectedPriceLeft[i].Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedPriceLeft[i].Name, result.Name)
		}
		if expectedPriceLeft[i].Price != result.Price {
			t.Errorf("unexpected name: expected %d result: %d", expectedPriceLeft[i].Price, result.Price)
		}
	}

// Проверка условия, когда установлена верхняя граница цены.
// filter.PriceLeft == nil && filter.PriceRight != nil && item.Price <= *filter.PriceRight.
// Это означает что выше этой границы товары в выборку попадать не должны.
// Нижний порог цен PriceLeft и Limit при проверке данного условия мы не устанавливаем.
// Limit заданный по умолчанию равен 5, поэтому мы должны получить не более 5 позиций.

	// Устанавливаем верхнюю границу цен.
	var priceRightInput int64 = 16690000
	inputFiter = &ItemFilter{
		PriceRight: &priceRightInput,
	}
	// Создаём ожидаемый результат.
	expectedPriceRight := []*models.Item {
		1: &models.Item{
			ID: 1,
			Name: "Intel i3 Core",
			Price: 1000,
		},
		2: &models.Item{
			ID: 2,
			Name: "Intel i5 Core",
			Price: 2000,
		},
		3: &models.Item{
			ID: 3,
			Name: "Intel i7 Core",
			Price: 3000,
		},
		8: &models.Item{
			ID: 8,
			Name: "GIGABYTE X570 AORUS ELITE",
			Price: 16690000,
		},
	}
// Вызовем функцию ListItems и передадим в неё параметры указанные в priceRightInput.
// Запишем полученный результат в переменную result и ошибку в err.
	res, err = db.ListItems(inputFiter)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Инициализируем слайс и записываем результат в items.
	items = make([]*models.Item, 0, len(res))

	// Теперь проитерируемся по полученному слайсу и сравним ID, Name и Price полученные с ожидаемыми.
	for i, result := range items {
		if expectedPriceRight[i].ID != result.ID {
			t.Errorf("unexpected name: expected %d result: %d", expectedPriceRight[i].ID, result.ID)
		}
		if expectedPriceRight[i].Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedPriceRight[i].Name, result.Name)
		}
		if expectedPriceRight[i].Price != result.Price {
			t.Errorf("unexpected name: expected %d result: %d", expectedPriceRight[i].Price, result.Price)
		}
	}

// Проверка условия, когда одновременно установлены нижняя и верзняя границы цены.
// filter.PriceLeft != nil && filter.PriceRight != nil &&
// item.Price >= *filter.PriceLeft  && item.Price <= *filter.PriceRight.
// Это означает что товары с ценой ниже PriceLeft в выборку попадать не должны и
// товары с ценой выше PriceRight в выборку попадать не должны.
// Limit при проверке данного условия мы не устанавливаем. Limit заданный по умолчанию равен 5,
// поэтому мы должны получить не более 5 позиций.

	// Устанавливаем верхнюю и нижнюю границу цен.
	priceLeftInput  = int64(3000)
	priceRightInput  = int64(16690000)

	inputFiter = &ItemFilter{
		PriceLeft: &priceLeftInput,
		PriceRight: &priceRightInput,
	}
	// Создаём ожидаемый результат.
	expectedAll := []*models.Item {
		7:&models.Item{
			ID: 7,
			Name: "GIGABYTE Z490 AORUS ULTRA G2",
			Price: 30440000,
		},
		8: &models.Item{
			ID: 8,
			Name: "GIGABYTE X570 AORUS ELITE",
			Price: 16690000,
		},
	}
// Вызовем функцию ListItems и передадим в неё параметры указанные в inputFiter.
// Запишем полученный результат в переменную result и ошибку в err.
	res, err = db.ListItems(inputFiter)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Инициализируем слайс и записываем результат в items.
	items = make([]*models.Item, 0, len(res))

	// Теперь проитерируемся по полученному слайсу и сравним ID, Name и Price полученные с ожидаемыми.
	for i, result := range items {
		if expectedAll[i].ID != result.ID {
			t.Errorf("unexpected name: expected %d result: %d", expectedAll[i].ID, result.ID)
		}
		if expectedAll[i].Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedAll[i].Name, result.Name)
		}
		if expectedAll[i].Price != result.Price {
			t.Errorf("unexpected name: expected %d result: %d", expectedAll[i].Price, result.Price)
		}
	}

// Проверим работу Offset и Limit. Получим список позиций товаров для 4 страницы.
// Limit - это колличество товаров на странице.
// Offset - это смещение.
// Если у нас в выборке всего 8 позиций товаров, то какие позиции попадут на 4 страницу мы може узнать
// исходя из следующей формулы Offset/Limit + 1.
// Если Offset = 0, а Limit = 2 то тогда у нас получается 1 страница и 2 позиции товара 1 и 2.
// Если	Offset = 6, а Limit = 2 то тогда мы получим 4 страницу (6/2=3+1) и 2 позиции товара 7 и 8.

	// Устанавливаем Offset=6 Limit=2
	inputFiter = &ItemFilter{
		Limit : 2,
		Offset: 6,
	}
	// Создаём ожидаемый результат.
	expectedOffset := []*models.Item {
		3: &models.Item{
			ID: 3,
			Name: "Intel i7 Core",
			Price: 3000,
		},
		8: &models.Item{
			ID: 8,
			Name: "GIGABYTE X570 AORUS ELITE",
			Price: 16690000,
		},
	}
// Вызовем функцию ListItems и передадим в неё параметры указанные в inputFiter.
// Запишем полученный результат в переменную result и ошибку в err.
	res, err = db.ListItems(inputFiter)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	// Инициализируем слайс и записываем результат в items.
	items = make([]*models.Item, 0, len(res))

	// Теперь проитерируемся по полученному слайсу и сравним ID, Name и Price полученные с ожидаемыми.
	for i, result := range items {
		if expectedOffset[i].ID != result.ID {
			t.Errorf("unexpected name: expected %d result: %d", expectedOffset[i].ID, result.ID)
		}
		if expectedOffset[i].Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedOffset[i].Name, result.Name)
		}
		if expectedOffset[i].Price != result.Price {
			t.Errorf("unexpected name: expected %d result: %d", expectedOffset[i].Price, result.Price)
		}
	}
}