package user_mapper

//go:generate mockgen -source=user_entry.go -destination=./mock/user_entry.go
type UserEntryMapper interface {
	// DEFINE METHODS
}