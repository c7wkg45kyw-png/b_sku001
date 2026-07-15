package usecase

import "bsku001/backend/internal/model"

type SummaryRepositoryPort interface {
	Get(merchantID string) (model.SummaryResponse, error)
}

type SummaryUsecase struct{ repo SummaryRepositoryPort }

func NewSummaryUsecase(repo SummaryRepositoryPort) *SummaryUsecase {
	return &SummaryUsecase{repo: repo}
}

func (u *SummaryUsecase) Get(auth model.AuthContext) (model.SummaryResponse, error) {
	return u.repo.Get(auth.MerchantID)
}
