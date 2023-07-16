.PHONY: dev
dev:
	air

.PHONY: gen
gen:
	sqlc generate

.PHONY: migrate
migrate:
	dbmate migrate
