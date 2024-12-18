### FOR LOCAL TESTING
```bash
## To bring up the SQL docker 
MariaDB: docker run --name test-mariadb -e MYSQL_ROOT_PASSWORD=12345 -p 3306:3306 -d docker.io/library/mariadb:10.3

## To login to the docker 
MariaDB: docker exec -it test-mariadb bash

## Once logged enter mysql using
mysql -u root -p12345

## Once inside sql, create a database named test using command
create database test;

## To start the server
go run runner/main.go

## To use the cli tool start the server and use the main.go at the root directory level
