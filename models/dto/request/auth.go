package request

type RegisterUser struct {
	CustomerID string `json:"customer_id" validate:"required,uuid4"`
	Username   string `json:"username" validate:"required,alphanum,min=3"`
	Password   string `json:"password" validate:"required,min=8"`
	Role       string `json:"role" validate:"oneof=CUSTOMER TELLER MANAGER ADMIN"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
