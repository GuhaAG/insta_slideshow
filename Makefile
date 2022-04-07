build:
	go build
build_fe:
	cd client && yarn && yarn build	
run:
	go run App.go