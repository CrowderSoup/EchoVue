run: build-app run-server

run-server:
	go run cmd/server/main.go

build-app:
	cd app && yarn build

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -c reflex.conf
