package model

type PaginationOpts struct {
	Skip  uint
	Limit uint
}

type Filter struct {
	IsActive bool
	Value    string
}
