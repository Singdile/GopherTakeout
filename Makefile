#定义变量
DB_URL_TEST="user=postgres password=123456 dbname=takeout_test port=5434 sslmode=disable"
MIGRATIONS_DIR=./migrations

#定义下列的名称为动作名
.PHONY: dev-up dev-down dev-logs dev-reset test-up test-down test-logs test-reset dev-psql test-psql

# ============================================
# Docker Database Commands
# ============================================
# Development database
dev-up:
	docker compose -f docker/docker-compose.dev.yml up -d
	@echo "Development database started on port 5433"
dev-down:
	docker compose -f docker/docker-compose.dev.yml down
dev-logs:
	docker compose -f docker/docker-compose.dev.yml logs -f
dev-reset:
	docker compose -f docker/docker-compose.dev.yml down -v
	docker compose -f docker/docker-compose.dev.yml up -d

# Test database (port 5433)
test-up:
	docker compose -f docker/docker-compose.test.yml up -d
	@echo "Test database started on port 5434"
test-down:
	docker compose -f docker/docker-compose.test.yml down
test-logs:
	docker compose -f docker/docker-compose.test.yml logs -f
test-reset:
	docker compose -f docker/docker-compose.test.yml down -v
	docker compose -f docker/docker-compose.test.yml up -d

# Interactive database access
dev-psql:
	docker exec -it gopher_takeout_db psql -U postgres -d go_takeout

test-psql:
	docker exec -it gopher_takeout_test psql -U postgres -d gopher_takeout_test

# Both databases
db-up: dev-up test-up
	@echo "All databases started"
db-down: dev-down test-down
	@echo "All databases stopped"
db-reset: dev-reset test-reset
	@echo "All databases reset"
