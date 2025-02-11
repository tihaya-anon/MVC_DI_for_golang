package user_service

//go:generate mockgen -source=user_auth.go -destination=./mock/user_auth.go
type UserAuthService interface {
	// DEFINE METHODS
}