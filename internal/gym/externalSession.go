package gym

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// - VALUE OBJECTS ------

// Sesion ID ------------
var ErrInvalidSessionID = errors.New("invalid Session ID")

// ExternalSessionID represents the session unique identifier.
type ExternalSessionID struct {
	value string
}

// NewExternalSessionID instantiate the VO for SessionID
func NewExternalSessionID(value string) (ExternalSessionID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return ExternalSessionID{}, fmt.Errorf("%w: %s", ErrInvalidSessionID, value)
	}

	return ExternalSessionID{
		value: v.String(),
	}, nil

}

// String type converts the ExternalSessionID into string.
func (id ExternalSessionID) String() string {
	return id.value
}

// ACCES TOKEN------------
var ErrEmptyAccessToken = errors.New("the field Access Token can not be empty")

type AccessToken struct {
	value string
}

func NewAccessToken(value string) (AccessToken, error) {
	if value == "" {
		return AccessToken{}, ErrEmptyAccessToken
	}

	return AccessToken{
		value: value,
	}, nil

}

func (v AccessToken) String() string {
	return v.value
}

// ErrEmptyRefreshToken REFRESH TOKEN------------
var ErrEmptyRefreshToken = errors.New("the field Refresh Token can not be empty")

type RefreshToken struct {
	value string
}

func NewRefreshToken(value string) (RefreshToken, error) {
	if value == "" {
		return RefreshToken{}, ErrEmptyRefreshToken
	}

	return RefreshToken{
		value: value,
	}, nil

}

func (v RefreshToken) String() string {
	return v.value
}

// SCOPE ------------
var ErrEmptyScope = errors.New("the field Scope can not be empty")

type Scope struct {
	value string
}

func NewScope(value string) (Scope, error) {
	if value == "" {
		return Scope{}, ErrEmptyScope
	}

	return Scope{
		value: value,
	}, nil

}

func (v Scope) String() string {
	return v.value
}

// TOKEN TYPE------------
var ErrEmptyTokenType = errors.New("the field Token Type can not be empty")

type TokenType struct {
	value string
}

func NewTokenType(value string) (TokenType, error) {
	if value == "" {
		return TokenType{}, ErrEmptyTokenType
	}

	return TokenType{
		value: value,
	}, nil

}

func (v TokenType) String() string {
	return v.value
}

// Domain Object  ---------

type ExternalSession struct {
	id           ExternalSessionID
	userID       UserID
	accessToken  AccessToken
	refreshToken RefreshToken
	scope        Scope
	tokenType    TokenType
}

// ExternalSessionRepository defines the expected behaviour from a external session storage.
type ExternalSessionRepository interface {
	Update(ctx context.Context, externalSession ExternalSession) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ExternalSessionRepository

// ExternalSessionClient defines the expected behaviour from a external api conection.
type ExternalSessionClient interface {
	GetToken(user User) (ExternalSession, error)
}

func NewExternalSession(id, userID, accessToken, refreshToken, scope, tokenType string) (ExternalSession, error) {

	idVO, err := NewExternalSessionID(id)
	if err != nil {
		return ExternalSession{}, err
	}

	userIdVO, err := NewUserID(userID)
	if err != nil {
		return ExternalSession{}, err
	}

	accessTokenVO, err := NewAccessToken(accessToken)
	if err != nil {
		return ExternalSession{}, err
	}

	refreshTokenVO, err := NewRefreshToken(refreshToken)
	if err != nil {
		return ExternalSession{}, err
	}

	scopeVO, err := NewScope(scope)
	if err != nil {
		return ExternalSession{}, err
	}

	tokenTypeVO, err := NewTokenType(tokenType)
	if err != nil {
		return ExternalSession{}, err
	}

	externalSession := ExternalSession{
		id:           idVO,
		userID:       userIdVO,
		accessToken:  accessTokenVO,
		refreshToken: refreshTokenVO,
		scope:        scopeVO,
		tokenType:    tokenTypeVO,
	}

	return externalSession, nil
}

// GETTERS

func (s ExternalSession) ID() ExternalSessionID {
	return s.id
}

func (s ExternalSession) UserID() UserID {
	return s.userID
}

func (s ExternalSession) AccessToken() AccessToken {
	return s.accessToken
}

func (s ExternalSession) RefreshToken() RefreshToken {
	return s.refreshToken
}

func (s ExternalSession) Scope() Scope {
	return s.scope
}

func (s ExternalSession) TokenType() TokenType {
	return s.tokenType
}
