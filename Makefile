run:
	go run cmd/api/main.go

hot-reload:
	nodemon --exec "go run cmd/api/main.go" --signal SIGTERM --ext go,mod --watch .