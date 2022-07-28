package repo

import (
	// "crud-golang/config"
	"crud-golang/models"
	"crud-golang/utils"

	"github.com/rs/zerolog/log"
)

// func AuthHandler(c *gin.Context) {
// 	var user models.User
// }

func CreateUser(user *models.User) (*models.User, error) {
	db, conErr := utils.GetDatabaseConnection()
	if conErr != nil {
		log.Err(conErr).Msg("Error occurred while getting a DB connection from the connection pool")
		return nil, conErr
	}

	result := db.Create(&user)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return user, nil
}

func GetAllUsers(user *models.User, pagination *models.Pagination) (*[]models.User, int64, error) {
	db, conErr := utils.GetDatabaseConnection()
	if conErr != nil {
		log.Err(conErr).Msg("Error occurred while getting a DB connection from the connection pool")
		return nil, 0, conErr
	}
	var users []models.User
	var totalRows int64 = 0
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	// queryBuider := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.User{}).Where(user).Find(&users)
	if result.Error != nil {
		msg := result.Error
		return nil, totalRows, msg
	}
	// errCount := config.DB.Model(&models.User{}).Count(&totalRows).Error
	errCount := db.Model(&models.User{}).Count(&totalRows).Error
	if errCount != nil {
		return nil, totalRows, errCount
	}
	return &users, totalRows, nil
}
