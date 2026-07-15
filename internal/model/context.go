package model

type AuthContext struct {
	UserID     string
	MerchantID string
	Scopes     []string
}
