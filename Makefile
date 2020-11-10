compile:
	env GOOS=linux go build -ldflags="-s -w" -o bin/atm_withdrawal *.go
	chmod 0777 bin/* -v
run:
	make compile
	./bin/atm_withdrawal
clear:
	rm -rf ./bin -v
test:
	go test -coverpkg=./... ./...
test-report:
	go test -coverpkg=./... -coverprofile=coverage.out -covermode=count ./...
	go tool cover -html=coverage.out
run-docker:
	docker-compose up -d
run-docker-clean:
	docker-compose build --no-cache
	make run-docker