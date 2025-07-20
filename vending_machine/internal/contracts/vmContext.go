package contracts

import "vending_machine/internal/models"

type VMContext interface {
	SetState(MachineState)
	GetBalanceCash() int
	SetBalanceCash(int)
	GetItems() []models.Item
	SetItems([]models.Item)
}
