package matchbook

type LoginResponse struct {
	SessionToken string `json:"session-token"`
	UserId string `json:"user-id"`
	Role string `json:"role"`
	Account AccountDetails `json:"account"`
}

type AccountDetails struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Name NameDetails `json:"name"`
}

type NameDetails struct {
	FirstName string `json:"first"`
	LastName string `json:"last"`
}