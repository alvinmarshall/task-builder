start:
	docker-compose up -d

stop:
	docker-compose down

mock-port:
	mockgen -destination internal/mock/user_repository.go  taskbuilder/internal/core/port UserRepository
	mockgen -destination internal/mock/user_service.go  taskbuilder/internal/core/port UserService

unit-test:
	go test -v ./...

unit-cov:
	go test -v -cover ./...