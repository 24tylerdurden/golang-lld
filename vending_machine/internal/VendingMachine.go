package internal

import (
	"vending_machine/internal/contracts"
	"vending_machine/internal/models"
)

type VendingMachine struct {
	State       contracts.MachineState
	Items       []models.Item
	BalanceCash int
}

// NewVendingMachine creates a new VendingMachine instance
func NewVendingMachine(items []models.Item, initialBalance int) *VendingMachine {
	return &VendingMachine{
		Items:       items,
		BalanceCash: initialBalance,
	}
}

func (v *VendingMachine) SetState(nextState contracts.MachineState) {
	v.State = nextState
}

func (v *VendingMachine) AcceptCash(amount int) int {
	return v.State.AcceptCash(amount)
}

func (v *VendingMachine) Checkout() *models.Item {
	if v.State.ValidateRequest() {
		return v.State.DispenseProduct()
	}
	return nil
}

func (v *VendingMachine) GetBalanceCash() int {
	return v.BalanceCash
}

func (v *VendingMachine) SetBalanceCash(amount int) {
	v.BalanceCash = amount
}

func (v *VendingMachine) GetItems() []models.Item {
	return v.Items
}

func (v *VendingMachine) SetItems(items []models.Item) {
	v.Items = items
}
