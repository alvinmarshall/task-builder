build:
	docker build -t taskbuilder:v1.0 .
start:
	docker-compose up -d

stop:
	docker-compose down

run: build start

mock-port:
	mockgen -package mock_port -destination internal/mock/user_repository.go  --source= taskbuilder/internal/core/port UserRepository
	mockgen -package mock_port -destination internal/mock/user_service.go  --source= taskbuilder/internal/core/port UserService
	mockgen -package mock_port -destination internal/mock/jwt_service.go  --source= taskbuilder/internal/core/service JwtService
	mockgen -package mock_port -destination internal/mock/auth_service.go  --source= taskbuilder/internal/core/service AuthService
	mockgen -package mock_port -destination internal/mock/encryption_service.go  --source= taskbuilder/internal/core/service EncryptionService

unit-test:
	go test -v ./...

unit-cov:
	go test -v -cover ./...