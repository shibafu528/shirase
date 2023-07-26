package repo

import (
	"context"
	"fmt"

	"github.com/shibafu528/shirase/db"
	"github.com/shibafu528/shirase/entity"
)

type StatusRepository struct{}

func NewStatusRepository() *StatusRepository {
	return &StatusRepository{}
}

func (r *StatusRepository) GetStatusesByAccountID(ctx context.Context, aid int64) ([]entity.Status, error) {
	_, q := db.Open()
	rows, err := q.GetStatusesByAccountID(ctx, aid)
	if err != nil {
		return nil, fmt.Errorf("[repo.StatusRepository.GetStatusesByAccountID] failed to query: %w", err)
	}

	var statuses []entity.Status
	for _, row := range rows {
		statuses = append(statuses, entity.Status{
			ID:            row.ID,
			AccountID:     row.AccountID,
			ActivityPubID: row.ActivityPubID,
			Text:          row.Text,
			CreatedAt:     row.CreatedAt,
			UpdatedAt:     row.UpdatedAt,
		})
	}

	return statuses, nil
}
