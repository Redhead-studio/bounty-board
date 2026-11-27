.PHONY: fmt migrate-up migrate-down migrate-status

fmt:
	cd backend && go fmt ./...

migrate-up:
	cd backend && goose -dir migrations postgres "$(DATABASE_DSN)" up

migrate-down:
	cd backend && goose -dir migrations postgres "$(DATABASE_DSN)" down

migrate-status:
	cd backend && goose -dir migrations postgres "$(DATABASE_DSN)" status