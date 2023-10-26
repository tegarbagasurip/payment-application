package dto

type PaginationQueryParam struct {
	Page, Offset, Limit int
}

type PaginationReturn struct {
	CurrentPage, LimitRows, StartIndex int
}

type PaginationResponse struct {
	Page        int `json:"currentPage"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalRows   int `json:"totalRows"`
	TotalPages  int `json:"totalPages"`
}
