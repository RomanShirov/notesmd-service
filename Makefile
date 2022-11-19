service-deploy:
	rm -r build/
	cd internal/web && git clone https://github.com/RomanShirov/notesmd-app
	cd internal/web/notesmd-app/frontend && npm install && npm run build
	go build -o build/ cmd/app/app.go
	cp .env build/
	mkdir build/assets && cp -r internal/web/notesmd-app/frontend/dist/. build/assets
	yes | rm -r internal/web
	cd build && ./app


