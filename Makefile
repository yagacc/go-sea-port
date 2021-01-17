.PHONY: start
start:
	go mod vendor && docker-compose up -d

stop:
	docker-compose down --remove-orphans

clean:
	go clean && go mod tidy & go mod vendor

format:
	go fmt ./...

vet:
	go vet ./...
