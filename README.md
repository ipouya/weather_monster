Base on (uncle bob) clean architecture

## Technologies
* Golang 1.13 
* Clean Architecture (uncle bob)
* Postgresql
* Go Mod + vendor
* Gorm (golang sql ORM)
* Gin ( web framework)

## Getting Started

The easiest way to get started is run `go run main.go`

For running this application you must set the database config from config.yml in root folder 

## Running the tests

I made a testHelper package that can start test database and router

you must set database config in testHelper/testServer.go file

you can find some test in these files
1. temperature/delivery/http/temperature_test.go
2. webhook/delivery/http/webhook_test.go
3. city/delivery/http/city_test.go

for example testing forecast api you can find it in temperature_test.go file

