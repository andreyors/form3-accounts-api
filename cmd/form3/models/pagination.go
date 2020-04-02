package models

type Pagination struct {
	Number int
	Size   int
}

func NewPagination(page int, size int) Pagination {
	return Pagination{page, size}
}
