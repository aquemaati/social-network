package model

type Auth struct {
	Id       string `json:"Id"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Register struct {
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	BirthDate       string `json:"BirthDate"`

	// OPTIONNAL
	ProfilePicture any    `json:"ProfilePicture"`
	Username       string `json:"Username"`
	AboutMe        string `json:"AboutMe"`
	Gender         string `json:"Gender"`
}
