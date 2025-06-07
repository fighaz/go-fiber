package request

type UserRequest struct {
	Name    string `json:"name"  validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address"`
	Phone   string `json:"phone" validate:"required"`
}
type UserUpdateRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
