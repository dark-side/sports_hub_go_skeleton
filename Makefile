.PHONY: run
.PHONY: vendor

run:
	docker compose up --build

vendor:
	go mod vendor