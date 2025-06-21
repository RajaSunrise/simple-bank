package response

import "time"

type AuthResponse struct {
	UserID      string    `json:"user_id"`
	CustomerID  string    `json:"customer_id"`
	Username    string    `json:"username"`
	Role        string    `json:"role"`
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}
