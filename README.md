## Tech demo for Pismo interview

This README provides instructions on how to set up and run the project.

#### Requirements to start the project

- Docker
- docker-compose
- air
- golang ~v1.21.1

#### How to Set Up
1. Clone the repository locally:
```bash
git clone git@github.com:Kintuda/tech-challenge-pismo.git
```

2. Create a .env file with your local settings. You can use the .env.example file as a template to start.

* Install the required dependencies using Go modules:
```
go mod tidy
```

3. Start docker-compose and spin the Postgres container
```bash
docker-compose up -d
```

4. Run the migration files
```bash
make migration-up
```

5. Start the API
```bash
go run main.go serve
OR
make dev
```

#### Creating new migrations
```
make create-new-migration name="operation type seed"
```

#### Running test
```
make test
```
