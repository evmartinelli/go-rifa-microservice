package usecase

import (
	"context"
	"fmt"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
)

// TranslationUseCase -.
type RaffleUseCase struct {
	repo RaffleRepo
}

// New -.
func NewRaffleUseCase(r RaffleRepo) *RaffleUseCase {
	return &RaffleUseCase{
		repo: r,
	}
}

func (uc *RaffleUseCase) Create(ctx context.Context, model entity.Raffle) error {
	return uc.repo.Create(ctx, model)
}

func (uc *RaffleUseCase) GetAvailableRaffle(ctx context.Context) ([]entity.Raffle, error) {
	res, err := uc.repo.GetAvailableRaffle(ctx)
	if err != nil {
		return []entity.Raffle{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
	}

	return res, nil
}
