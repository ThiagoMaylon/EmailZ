package campaing

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/domain/internal-Errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaing *Campaing) error {
	args := r.Called(campaing)
	return args.Error(0)
}

var (
	newCampaing = contract.NewCampaing{
		Name:    "Teste X",
		Content: "Body",
		Emails:  []string{"teste1@teste.com"},
	}
	service = Service{}
	repositoryM   = new(repositoryMock)
)

func Test_Create_Campaing(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(newCampaing)

	assert.NotNil(id)
	assert.Nil(err)
}
func Test_Create_ValidateDoaminError(t *testing.T) {
	assert := assert.New(t)
	newCampaing.Name = ""
	_, err := service.Create(newCampaing)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
}
func Test_Create_SaveCampaing(t *testing.T) {
	repositoryM.On("Save", mock.MatchedBy(func(campaing *Campaing) bool {
		if campaing.Name != newCampaing.Name ||
			campaing.Content != newCampaing.Content ||
			len(campaing.Contacts) != len(newCampaing.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service := Service{Repository: repositoryM}
	service.Create(newCampaing)

	repositoryM.AssertExpectations(t)

}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryM.On("Save", mock.Anything).Return(internalerrors.ErrInternal)
	service.Repository = repositoryM

	_, err := service.Create(newCampaing)

	assert.True(errors.Is(internalerrors.ErrInternal, err))

}

