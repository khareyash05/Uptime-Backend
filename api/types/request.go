package types

type RequestUser struct {
	UserId string `json:"user_id"`
	URL    string `json:"url"`
}

type RequestUser2 struct{
	UserId string `json:"user_id"`
}
