package restapi

type UserData struct {
	Username string `json:user`
	Password string `json:password`
	Email    string `json:email`
}

type Response struct {
	Message string `json:message`
}

type Login struct {
	Username string `json:user`
	Password string `json:password`
}
