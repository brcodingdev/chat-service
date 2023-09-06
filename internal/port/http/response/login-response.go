package response

import (
	"github.com/brcodingdev/chat-service/internal/pkg/model"
)

// LoginResponse ...
type LoginResponse struct {
	User     model.User `json:"User"`
	JwtToken string     `json:"Token"`
}
