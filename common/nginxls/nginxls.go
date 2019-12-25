package nginxls

type jwtSetting struct {
	enabled    bool
	secret     string
	prevSecret string
}

type Nginxls struct {
	Jwt func(secrect string)
}