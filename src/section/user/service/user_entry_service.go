package user_service

//go:generate mockgen -source=user_entry.go -destination=./mock/user_entry.go
type UserEntryService interface {
	// DEFINE METHODS
}