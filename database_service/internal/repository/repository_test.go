package repository

import (
	"db_service/internal/repository/test/mocks"
	"db_service/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InterfaceRepositoryCreate(t *testing.T) {
	t.Run("успешное создание", func(t *testing.T) {
		mockrepo := new(mocks.InterfaceRepository)

		mockrepo.On("Create", "test arg").Return(nil)

		err := mockrepo.Create("test arg")

		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
	})
}

func Test_InterfaceRepositoryList(t *testing.T) {
	t.Run("успешное чтение", func(t *testing.T) {
		mockrepo := new(mocks.InterfaceRepository)

		m := make([]models.DatabaseModels, 0, 50)
		for i := 0; i < 50; i++ {
			testModel := models.DatabaseModels{
				ID:    1,
				Title: "title",
			}

			m = append(m, testModel)
		}

		mockrepo.On("List").Return(m, nil)

		res, err := mockrepo.List()

		assert.NoError(t, err)
		assert.Len(t, res, 50)
		mockrepo.AssertExpectations(t)
	})
}

func Test_InterfaceRepositoryDelete(t *testing.T) {
	t.Run("успешное создание", func(t *testing.T) {
		mockrepo := new(mocks.InterfaceRepository)

		mockrepo.On("Create", "test arg").Return(nil)

		err := mockrepo.Create("test arg")

		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
	})
}

func Test_InterfaceRepositoryDone(t *testing.T) {
	t.Run("успешное создание", func(t *testing.T) {
		mockrepo := new(mocks.InterfaceRepository)

		mockrepo.On("Create", "test arg").Return(nil)

		err := mockrepo.Create("test arg")

		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
	})
}
