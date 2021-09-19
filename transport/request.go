package transport

type CreateBook struct {
	Name    string `json:"name" validate:"required"`
	Creator string `json:"creator" validate:"required"`
}

type UpdateBook struct {
	Id      string `json:"id" validate:"required"`
	Name    string `json:"name"`
	Creator string `json:"creator"`
}
