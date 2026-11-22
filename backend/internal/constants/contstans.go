package constants

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleOwner UserRole = "owner"
)

type TaskStatus string

const (
	StatusNew     TaskStatus = "new"
	StatusPending TaskStatus = "pending"
	StatusRewiev  TaskStatus = "review"
	StatusClosed  TaskStatus = "closed"
)
