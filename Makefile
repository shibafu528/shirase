.PHONY: dev
dev:
	air

.PHONY: gen
gen:
	sqlc generate
