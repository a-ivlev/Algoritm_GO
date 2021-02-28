package repository

type baseFilter struct {
	Limit  int
	Offset int
}

type ItemFilter struct {
	baseFilter

	ItemIDs []int32

	PriceLeft  *int64
	PriceRight *int64
}

type OrderFilter struct {
	baseFilter
}
