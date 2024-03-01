package service

import (
	"strconv"
	"test/database"
	"test/model"

	"github.com/jmoiron/sqlx"
)

var db = database.Connection()

//add profile function

func AddProfile(userProfile model.Profile) (model.Profile, error) {
	_, err := db.Exec("INSERT INTO tb_user_profile (name, bio) VALUES ($1, $2)", userProfile.Name, userProfile.Bio)
	if err != nil {
		return model.Profile{}, err
	}
	return userProfile, err
}

func GetAllProfile() ([]model.Profile, error) {
	getProfiles := []model.Profile{}
	rows, err := db.Queryx("SELECT * FROM tb_user_profile")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		place := model.Profile{}
		rows.StructScan(&place)
		getProfiles = append(getProfiles, place)
	}
	return getProfiles, nil
}

// get all profile functions
func GetProfile(id int) (model.Profile, error) {
	var userProfile model.Profile
	idStr := strconv.Itoa(id)
	err := db.Get(&userProfile, "SELECT * FROM tb_user_profile WHERE id = $1", idStr)
	if err != nil {
		return model.Profile{}, err
	}
	return userProfile, nil
}

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) UpdatePictureURL(userID int, pictureURL string) error {
	query := "UPDATE tb_user_profile SET picture_url = $1 WHERE id = $2"
	_, err := s.db.Exec(query, pictureURL, userID)
	return err
}
