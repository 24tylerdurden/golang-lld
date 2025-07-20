package contracts

import "vending_machine/internal/models"

type MachineState interface {
	AcceptCash(int) int
	ValidateRequest() bool
	DispenseProduct() *models.Item
	DispenseCash() int
	CancelRequest()
	GetCash() int
}
