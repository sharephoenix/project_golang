package config

type Config struct {
	Name string
	Host string
	Port string
	Redis Redis
	Auth Auth
}

type Redis struct {
	Addr string
	Password string
	DB string
}

type Auth struct {
	AccessSecret string
}
