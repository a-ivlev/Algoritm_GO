package repository

import (
	"shop/models"
	"time"
)

var initMapDBitems = map[int32]*models.Item{
						1: &models.Item{
								ID: 1,
								Name: "Intel i3 Core",
								Price: 1000,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
						},
						2: &models.Item{
								ID: 2,
								Name: "Intel i5 Core",
								Price: 2000,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
						},
						3: &models.Item{
								ID: 3,
								Name: "Intel i7 Core",
								Price: 3000,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
						},
						4: &models.Item{
								ID: 4,
								Name: "GeForce GTX 1650",
								Price: 25000000,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
						},
						5: &models.Item{
								ID: 5,
								Name: "GeForce GTX 1050 Ti",
								Price: 19000000,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
						},
						6: &models.Item{
								ID: 6,
								Name: "GIGABYTE TRX40 AORUS PRO",
								Price: 36390000,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
						},
						7: &models.Item{
								ID: 7,
								Name: "GIGABYTE Z490 AORUS ULTRA G2",
								Price: 30440000,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
						},
						8: &models.Item{
								ID: 8,
								Name: "GIGABYTE X570 AORUS ELITE",
								Price: 16690000,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
						},
}