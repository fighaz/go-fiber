package request

type BookRequest struct {
	Title  string `json:"name"  validate:"required"`
	Author string `json:"author" validate:"required"`
	Cover  string `json:"cover"`
}
