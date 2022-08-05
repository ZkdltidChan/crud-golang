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

func GetAllUsers(user *models.User, pagination *models.Pagination) (*[]models.User, *models.Pagination, error) {
	// var paginationResp models.PaginationResp
	db, conErr := utils.GetDatabaseConnection()
	if conErr != nil {
		log.Err(conErr).Msg("Error occurred while getting a DB connection from the connection pool")
		return nil, nil, conErr
	}
	var users []models.User
	var totalRows int64 = 0
	queryBuider := db.Limit(pagination.GetLimit()).Offset(pagination.GetOffset()).Order(pagination.Sort)
	result := queryBuider.Model(&models.User{}).Where(user).Find(&users)

	if result.Error != nil {
		msg := result.Error
		return nil, &models.Pagination{}, msg
	}
	errCount := db.Model(&models.User{}).Count(&totalRows).Error
	if errCount != nil {
		return nil, &models.Pagination{}, errCount
	}
	// paginationResp.PageIndex = pagination.GetOffset()
	// paginationResp.TotalPage = models.GetTotalPages(totalRows, pagination.GetSize())
	// paginationResp.Size = pagination.GetLimit()
	return &users, &models.Pagination{}, nil
}
