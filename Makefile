start:
	docker-compose -f docker-compose.yaml exec andrade  go run cmd/main.go --bind 0.0.0.0 -m andrade
	exit 0

bash:
	docker-compose -f docker-compose.yaml exec api sh

up:
	docker-compose -f docker-compose.yaml up -d

down:
	docker-compose -f docker-compose.yaml down --timeout 0

deps:
	go mod tidy
