build:
	go build
run:
	go run App.go
make build_fe:
	cd client && yarn && yarn build