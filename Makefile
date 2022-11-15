service-build:
	cp .env build/
	go build -o build/ cmd/app/app.go
	cd build && ./service

