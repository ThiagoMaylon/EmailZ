package campaing

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)
var (
	name = "Nome x"
	content = "body"
	contacts = []string{"email1@teste.com", "email2@gteste.com"}
)
func Test_NewCampaing_CreateCampaing(t *testing.T) {
	assert := assert.New(t)
	
	campaing,_ := NewCampaing(name, content, contacts)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))

}

func Test_NewCampaing_IDIs(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.NotNil(campaing.ID)

}
func Test_NewCampaing_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.Greater(campaing.CreatedOn, now)

}
func Test_NewCampaing_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing("", content, contacts)

	assert.Equal("name is required min 5", err.Error())
}
func Test_NewCampaing_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing("12345678998765432112345698765", content, contacts)

	assert.Equal("name is required max 24", err.Error())
}
func Test_NewCampaing_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, "", contacts)

	assert.Equal("contet is required", err.Error())
}
func Test_NewCampaing_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, []string{})

	assert.Equal("contacts is required", err.Error())
}
