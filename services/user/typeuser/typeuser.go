package typeuser

type User struct {
	NickName string `json:"nickName"`
	Mobile string `json:"mobile"`
	Email string `json:"email"`
	Age int64 `json:"age"`
	Address string `json:"address"`
	Avatar string `json:"avatar"`
	AccessToken interface{} `json:"accessToken,omitempty"`
}

type MobileCode struct {
	Code string `json:"code"`
}

const (
	JwtUserField = "usr"
	JwtVersionField = "ver"
)

const (
	NETERROR_NO_USER = "没有找到用户"
)