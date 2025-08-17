package installments

type InstallmentStatus int

const (
	OPEN InstallmentStatus = iota
	PAID
	CANCELLED
)

var statusName = map[InstallmentStatus]string{
	OPEN:      "OPEN",
	PAID:      "PAID",
	CANCELLED: "CANCELLED",
}

func (is InstallmentStatus) String() string {
	return statusName[is]
}

type Installment struct {
	Number int
	Status InstallmentStatus
}

func NewInstallment() *Installment {
	return &Installment{
		Number: 1,
		Status: OPEN,
	}
}
