package request

type UserRequest struct {
	Name     string `json:"name"  validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Address  string `json:"address"`
	Phone    string `json:"phone" validate:"required"`
}
type UserUpdateRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
