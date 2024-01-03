package campaing

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/domain/internal-Errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaing contract.NewCampaing) (string, error) {

	campaing, err := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaing)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaing.ID, nil
}
