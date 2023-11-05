build:
	go build
test:
	go test -v ./test/unit
dev:
	air serve
create-new-migration:
	go run main.go migration create $(name)
migration-up:
	go run main.go migration up
migration-down:
	go run main.go migration down
gen-docs:
	swag init --parseDependency