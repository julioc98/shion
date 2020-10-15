test:
	go test ./... -v -coverprofile fmtcoverage.html fmt

test/docker:
	docker-compose run shion make test

run/api:
	go run cmd/api/main.go

run/db:
	docker-compose up shiondb

run/docker:
	docker-compose up --build