package model

type Page struct {
	Current int `json:"current"`
	Size    int `json:"size"`
}

type PageResponse[T any] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}
