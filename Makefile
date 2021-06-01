run: dev-app dev-server

dev-server:
	go run cmd/server/main.go

dev-app:
	cd app && yarn build

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -t 500ms -s -r '\.go$$' make run
