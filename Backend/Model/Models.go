package model

type RegisterContextKey string

const (
	RegisterCtx RegisterContextKey = "register"
)

type Auth struct {
	Id              string `json:"Id"`
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
}

type Register struct {
	Auth      Auth   `json:"Auth"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	BirthDate string `json:"BirthDate"`

	// OPTIONNAL
	ProfilePicture any    `json:"ProfilePicture"`
	Username       string `json:"Username"`
	AboutMe        string `json:"AboutMe"`
	Gender         string `json:"Gender"`
}
