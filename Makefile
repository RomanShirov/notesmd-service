build-frontend:
	clear
	cd internal/web && git clone https://github.com/RomanShirov/notesmd-app
	cd internal/web/notesmd-app/frontend && npm install && npm run build

service-build:
	clear
	make build-frontend
	go build -o build/ cmd/app/app.go
	cp .env build/
	mkdir build/assets && cp -r internal/web/notesmd-app/frontend/dist/. build/assets
	clear

service-run:
	docker-compose --env-file .env up -d
	cd build && clear && ./app


service-clear:
	rm -r build/
	yes | rm -r internal/web/notesmd-app
	clear

service-reset:
	docker-compose down --remove-orphans
	make service-clear

make rebuild:
	make service-reset
	make service-build
	make service-run


