package interfaces

import "k6/eg2/models"

type IDemo interface {
	CreateToken(customer *models.Sample) (*models.DBResponse, error)
}
