build: tidy
	go build -o bin/actividad-03 .
	docker compose up -d --build

tidy:
	go mod tidy

run:
	docker compose up -d

stop:
	docker compose down