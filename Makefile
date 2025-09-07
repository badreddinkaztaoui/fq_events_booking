create-db:
	docker-compose up -d db

run-server:
	air run -c .air.toml

.PHONY: create-db run-server
