package apiconsts

type contextKey string

const (
	RequestIDKey  contextKey = "request_id"
	RequestLogger contextKey = "request_logger"
	ID            contextKey = "id"
	UserIDKey     contextKey = "userID"
	UserRoleKey   contextKey = "userRole"
)
