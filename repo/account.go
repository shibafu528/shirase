package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/shibafu528/shirase/db"
	"github.com/shibafu528/shirase/entity"
)

type AccountRepository struct{}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, id int64) (*entity.Account, error) {
	_, q := db.Open()
	row, err := q.GetAccount(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrRecordNotFound
	} else if err != nil {
		return nil, fmt.Errorf("[repo.AccountRepository.GetAccountByID] failed to query: %w", err)
	}

	return &entity.Account{
		ID:            row.ID,
		Username:      row.Username,
		Domain:        row.Domain,
		DisplayName:   row.DisplayName,
		PrivateKey:    row.PrivateKey,
		PublicKey:     row.PublicKey,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
		ActivityPubID: row.ActivityPubID,
		Description:   row.Description,
	}, nil
}

func (r *AccountRepository) GetAccountByUsername(ctx context.Context, username string) (*entity.Account, error) {
	_, q := db.Open()
	row, err := q.GetAccountByUsername(ctx, username)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrRecordNotFound
	} else if err != nil {
		return nil, fmt.Errorf("[repo.AccountRepository.GetAccountByUsername] failed to query: %w", err)
	}

	return &entity.Account{
		ID:            row.ID,
		Username:      row.Username,
		Domain:        row.Domain,
		DisplayName:   row.DisplayName,
		PrivateKey:    row.PrivateKey,
		PublicKey:     row.PublicKey,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
		ActivityPubID: row.ActivityPubID,
		Description:   row.Description,
	}, nil
}

func (r *AccountRepository) GetAccountByActivityPubID(ctx context.Context, apid string) (*entity.Account, error) {
	_, q := db.Open()
	row, err := q.GetAccountByActivityPubID(ctx, sql.NullString{String: apid, Valid: true})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrRecordNotFound
	} else if err != nil {
		return nil, fmt.Errorf("[repo.AccountRepository.GetAccountByActivityPubID] failed to query: %w", err)
	}

	return &entity.Account{
		ID:            row.ID,
		Username:      row.Username,
		Domain:        row.Domain,
		DisplayName:   row.DisplayName,
		PrivateKey:    row.PrivateKey,
		PublicKey:     row.PublicKey,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
		ActivityPubID: row.ActivityPubID,
		Description:   row.Description,
	}, nil
}

func (r *AccountRepository) GetAccountIDByActivityPubID(ctx context.Context, apid string) (int64, error) {
	_, q := db.Open()
	id, err := q.GetAccountIDByActivityPubID(ctx, sql.NullString{String: apid, Valid: true})
	if errors.Is(err, sql.ErrNoRows) {
		return 0, ErrRecordNotFound
	} else if err != nil {
		return 0, fmt.Errorf("[repo.AccountRepository.GetAccountIDByActivityPubID] failed to query: %w", err)
	}

	return id, nil
}
