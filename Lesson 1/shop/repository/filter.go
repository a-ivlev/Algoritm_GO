package repository

type ItemFilter struct {
	PriceLeft  *int64
	PriceRight *int64
	Limit      int
	Offset     int
}
