package config

type Config struct {
	Name string //`json:"name"`
	Host string //`json:"host"`
	Port string //`json:"port"`
	Redis Redis //`json:"redis"`
}

type Redis struct {
	Addr string //`json:"addr"`
	Password string //`json:"password"`
	DB string //`json:"DB"`
}

// redis mongodb