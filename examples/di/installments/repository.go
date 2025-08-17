package installments

import (
	"database/sql"
)

type InstallmentRepositoryInterface interface {
	GetInstallment(number int) (*Installment, error)
}

type InstallmentRepository struct {
	db *sql.DB
}

func NewInstallmentRepository(db *sql.DB) *InstallmentRepository {
	return &InstallmentRepository{db}
}

func (ir *InstallmentRepository) GetInstallment(number int) (*Installment, error) {
	return NewInstallment(), nil
}
