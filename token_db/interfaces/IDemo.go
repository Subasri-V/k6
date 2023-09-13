package interfaces

import "k6/token_db/models"


type IDemo interface {
	CreateToken(customer *models.Sample) (*models.DBResponse, error)
	StoreData(customer *models.Sample)(*models.DBResponse, error)
}
