package service

import (
	"context"
	"fmt"
	"tc_backend/queries"
	"tc_backend/repository"
)

type TmpService struct {
	repository *repository.Repository
}

func NewTmpService() *TmpService {
	return &TmpService{
		repository: repository.NewRepository(),
	}
}

func (t *TmpService) GetAnswer(ctx context.Context) []queries.Tmp {
	rows, err := t.repository.Q.GetAllTMP(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return rows
}
