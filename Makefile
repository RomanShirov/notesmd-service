frontend:
	clear
	cd internal/web && git clone https://github.com/RomanShirov/notesmd-app

build:
	clear
	cd internal/web/notesmd-app/frontend && npm install && npm run build
	clear

run:
	docker-compose --env-file .env up -d

run-docker-build:
	docker-compose --env-file .env up --build -d

clear:
	yes | rm -r internal/web/notesmd-app
	clear

reset:
	docker-compose rm -f
	docker-compose down --remove-orphans
	make clear

