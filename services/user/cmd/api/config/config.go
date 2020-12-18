package config

type Config struct {
	Name  string
	Host  string
	Port  string
	Redis Redis
	Mongo Mongo
	Auth  Auth
}

type Redis struct {
	Addr     string
	Password string
	DB       string
}

type Mongo struct {
	Addr       string
	DB         string
	Collection string
}

type Mysql struct {
	Addr     string
	Password string
	DB       string
	desc     string
}

type Auth struct {
	AccessSecret string
}
