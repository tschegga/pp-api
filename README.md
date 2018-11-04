# API for the CoH-Pottpokal application

## Prerequisites

1. The application is written in Go. You can get it [here](https://golang.org).
2. You need a MySQL database (e.g. MySQL, MariaDB, PostgreSQL) running on your local machine or somewhere in the internet. There is a docker-compose file in the deployment folder which starts a MariaDB alongside phpMyAdmin. You might also be interested in the Vagrantfile which exposes ports to both services if you are running on Windows.
3. Your database should contain a database called 'pottpokal'. You can use the MySQL workbench file in /resources folder to create a SQL script that feeds your database with all the necessary databases and tables. If you want to have some testdata you can use the createDummyData.sql in the /deployment folder.

## Running the application

First move to the /bin folder. It will not exist if you clone the repository for the first time. In this case you need to create it. Then you can run the application with:
```bash
go run ../main.go
```

## Building the application

You can also build the application before running it. Similarly to above, navigate to the /bin folder and run:
```bash
go build ../main.go
./main
```

## Configuration

You can define the port, the URI of the database and the loglevel in /resources/config.yml.