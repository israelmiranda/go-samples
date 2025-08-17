//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/israelmiranda/go-samples/examples/di/installments"
)

var setRepositoryDependency = wire.NewSet(
	installments.NewInstallmentRepository,
	wire.Bind(
		new(installments.InstallmentRepositoryInterface),
		new(*installments.InstallmentRepository),
	),
)

func NewUseCase(db *sql.DB) *installments.InstallmentUseCase {
	wire.Build(
		setRepositoryDependency,
		installments.NewIntallmentUseCase,
	)
	return &installments.InstallmentUseCase{}
}
