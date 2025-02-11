package user_mapper

//go:generate mockgen -source=user_auth.go -destination=./mock/user_auth.go
type UserAuthMapper interface {
	// DEFINE METHODS
}