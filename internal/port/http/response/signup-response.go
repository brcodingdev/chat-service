package response

import (
	"github.com/brcodingdev/chat-service/internal/pkg/model"
)

// SignUpResponse ...
type SignUpResponse struct {
	User     model.User `json:"User"`
	JwtToken string     `json:"Token"`
}
