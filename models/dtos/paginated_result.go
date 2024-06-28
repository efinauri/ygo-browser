package dtos

import "yugioh-browser/models/entities"

type _PaginatedResult[T any] struct {
	Elements      []*T
	Page          int
	Size          int
	TotalElements int
}

type PaginatedCardResult _PaginatedResult[entities.Card]
