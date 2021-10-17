package mysql

import (
	"arnold/internal/gym"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
)

// ExternalSessionRepository is a MySQL gym.ExternalSessionRepository implementation.
type ExternalSessionRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewExternalSession initializes a MySQL-based implementation of gym.ExternalSessionRepository.
func NewExternalSessionRepository(db *sql.DB, dbTimeout time.Duration) *ExternalSessionRepository {

	return &ExternalSessionRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Update implements the gym.ExternalSessionRepository interface.
func (r *ExternalSessionRepository) Update(ctx context.Context, externalSession gym.ExternalSession) error {

	externalSessionSQLStruct := sqlbuilder.NewStruct(new(sqlExternalSession))
	query, args := externalSessionSQLStruct.Update(sqlExternalSessionTable, sqlExternalSession{ // Revisar si el metodo de sql es correcto
		ID:           externalSession.ID().String(),
		UserID:       externalSession.UserID().String(),
		AccessToken:  externalSession.AccessToken().String(),
		RefreshToken: externalSession.RefreshToken().String(),
		Scope:        externalSession.Scope().String(),
		TokenType:    externalSession.TokenType().String(),
	}).Build()
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist session on database: %v", err)
	}

	return nil
}
