Prepare DB:

start db docker:
docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=restream -d -p 5433:5433 postgres:latest

restore test data:
docker exec -i postgres pg_restore --dbname=restream --verbose --clean < ./DBscript

=========================================

Start Server:

go run main.go

=========================================

