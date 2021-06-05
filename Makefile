start:
	docker-compose up -d

stop:
	docker-compose down

mock-port:
	mockgen -destination internal/mock/user_repository.go  --source= taskbuilder/internal/core/port UserRepository
	mockgen -destination internal/mock/user_service.go  --source= taskbuilder/internal/core/port UserService

unit-test:
	go test -v ./...

unit-cov:
	go test -v -cover ./...