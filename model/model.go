package model

type Profile struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Bio        string `json:"bio" db:"bio"`
	Pictureurl string `json:"picture_url" db:"picture_url"`
}

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
