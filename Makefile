include .env
GOBASH = docker exec -it go_server /bin/sh
DB_URI = 'postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable'

migrateup:
		$(GOBASH) -c 'migrate -path ./migrations -database $(DB_URI) -verbose up'

migratedown:
		$(GOBASH) -c 'migrate -path ./migrations -database $(DB_URI) -verbose down'

serve:
		go run cmd/main.go

gobash:
	$(GOBASH)

fresh:
	$(MAKE) init-network
	$(MAKE) force-build
	$(MAKE) start

force-build:
	docker compose -f docker-compose.yml build --no-cache

start:
	docker compose -f docker-compose.yml up


init-network:
		docker network inspect docker-network > /dev/null 2>&1 || docker network create docker-network --subnet=172.20.0.0/16 --gateway=172.20.0.1
