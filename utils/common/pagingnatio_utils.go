package common

import (
	"payment-application/model/dto"
	"log"
	"math"
	"os"
	"strconv"
)

func CreatePaginationFromQueryParams(queryParams dto.PaginationQueryParam) dto.PaginationReturn {
	// read .env FILE
	err := LoadENV()
	if err != nil {
		log.Fatalln(err)
	}

	var (
		currentPage, limitRows, startIndex int
	)

	if queryParams.Page > 0 {
		currentPage = queryParams.Page
	} else {
		currentPage = 1
	}

	if queryParams.Limit == 0 {
		defaultLimit, _ := strconv.Atoi(os.Getenv("DEFAULT_ROWS_PER_PAGE"))
		limitRows = defaultLimit
	} else {
		limitRows = queryParams.Limit
	}

	startIndex = (currentPage - 1) * limitRows

	return dto.PaginationReturn{
		CurrentPage: currentPage,
		LimitRows:   limitRows,
		StartIndex:  startIndex,
	}
}


func CreatePaginationResponse(page, limit, totalRows int) dto.PaginationResponse {
	return dto.PaginationResponse{
		Page:        page,
		RowsPerPage: limit,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(limit))),
	}
}
