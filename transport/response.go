package transport

import "book-catalog/entity"

type GeneralResponse struct {
	Message string `json:"message"`
}

type GetList struct {
	Count    int           `json: "count"`
	ListBook []entity.Book `json: "listBook"`
}

type GetBook struct {
	Data entity.Book `json: "data"`
}

type ValidateResponse struct {
	Message []string
	Status  int
}
