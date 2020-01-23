package models

type Config struct {
	Server struct {
		Address  string
	}
	Sql struct {
		Type     string
		Host     string
		User 	 string
		DbName 	 string
		Password string
		Port     uint
		Sslmode  string
	}
}
