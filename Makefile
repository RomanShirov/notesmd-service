service-build:
	go build -o build/ cmd/app/app.go
	cp .env build/
	cd build && ./app

