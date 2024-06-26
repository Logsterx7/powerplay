package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s session) GetRosters() ([]models.Roster, error) {
	rosters := make([]models.Roster, 0)

	err := s.connection.Preload("Players").Preload("Captain").Find(&rosters)

	return resultsOrError(rosters, err)
}

func (s session) CreateRoster(roster *models.Roster) error {
	result := s.connection.Create(roster)

	return result.Error
}

func (s session) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := s.connection.Find(&user, "email = ?", email)

	return resultOrError(user, err)
}

func (s session) GetUserByEmails(emails []string) ([]*models.User, error) {
	users := make([]*models.User, 0)

	err := s.connection.Where("email IN (?)", emails).Find(&users)

	return resultsOrError(users, err)
}
