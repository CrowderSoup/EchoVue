dev: dev_server dev_app

dev_server:
	go run cmd/server/main.go

dev_app:
	cd app && yarn serve
