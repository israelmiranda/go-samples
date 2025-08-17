package installments

type InstallmentUseCase struct {
	repository InstallmentRepositoryInterface
}

func NewIntallmentUseCase(repository InstallmentRepositoryInterface) *InstallmentUseCase {
	return &InstallmentUseCase{repository}
}

func (uc *InstallmentUseCase) GetInstallment(number int) (*Installment, error) {
	return uc.repository.GetInstallment(number)
}
