package products

import (
	"fmt"
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

type MockStorageRepository struct {
	dataStorage []Product

}

func (ms *MockStorageRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	prodList := ms.dataStorage
	if prodList == nil{
		return []Product{}, fmt.Errorf("empty storage") 
	}
	return prodList, nil
}

func TestGetAllBySeller(t *testing.T) {
	data := []Product{{"mock1", "FEX112AA", "Tablet", 123.55}, {"mock2", "FEX112AA", "Notebook", 400.99}}

	mockStorage := MockStorageRepository{dataStorage: data}

	service := NewService(&mockStorage)

	response, err := service.GetAllBySeller("FEX112AA")

	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataStorage, response)
}

func TestGetAllBySellerNil(t *testing.T) {
	mockStorage := MockStorageRepository{dataStorage: nil}
	expectedError := errors.New("empty storage") // Cual es la diferencia ? fmt.Errorf("empty storage") 

	service := NewService(&mockStorage)

	response, err := service.GetAllBySeller("FEX112AA")

	assert.Nil(t, response)
	assert.Equal(t, expectedError, err)
}