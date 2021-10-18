package external_login

import (
	"arnold/internal/gym"
	"context"
)

type ExternalLooginService struct {
	ExternalSessionRepository gym.ExternalSessionRepository
	ExternalSessionClient     gym.ExternalSessionClient
}

func NewExternalLooginService(externalSessionRepository gym.ExternalSessionRepository, externalSessionClient gym.ExternalSessionClient) ExternalLooginService {

	return ExternalLooginService{
		ExternalSessionRepository: externalSessionRepository,
		ExternalSessionClient:     externalSessionClient,
	}
}

func (s ExternalLooginService) Loggin(ctx context.Context, userID string) error {

	id := "9b56c21f-c85d-485a-9aa3-c2b4137db90a"
	accessToken := "0d91220f-1dd3-42cb-a195-b9de525b753d"
	refreshToken := "ae7a60b5-8b43-41a0-8360-b6f66f680dc7"
	scope := "read write"
	tokenType := "bearer"

	externalSession, err := gym.NewExternalSession(id, userID, accessToken, refreshToken, scope, tokenType)
	if err != nil {
		return err
	}

	user, err := gym.NewUser(userID, "mstefanutti24@gmail.com", "holamundo")
	if err != nil {
		return err
	}
	_, err2 := s.ExternalSessionClient.getToken(user) // porq no reconoce el metodo??
	if err != nil {
		return err2
	}

	// err = s.ExternalSessionRepository.Update(ctx, externalSession)
	updateError := s.ExternalSessionRepository.Update(ctx, externalSession)

	return updateError
}
