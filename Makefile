build:
	docker-compose build app

run:
	docker-compose up app

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	$GOPATH/bin/swag init -g cmd/main.go
	
run-go:
	go run cmd/main.go
