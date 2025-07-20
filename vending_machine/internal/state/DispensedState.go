package state

import (
	"vending_machine/internal/contracts"
	"vending_machine/internal/models"
)

type DispenseState struct {
	vm contracts.VMContext
}

func NewDispenseState(vm contracts.VMContext) *DispenseState {
	return &DispenseState{vm: vm}
}

func (s *DispenseState) AcceptCash(amount int) int {
	return s.vm.GetBalanceCash()
}

func (s *DispenseState) ValidateRequest() bool {
	return false
}

func (s *DispenseState) DispenseProduct() *models.Item {
	return nil
}

func (s *DispenseState) DispenseCash() int {
	cash := s.vm.GetBalanceCash()
	s.vm.SetBalanceCash(0)
	// Transition back to IdleState
	s.vm.SetState(NewIdleState(s.vm))
	return cash
}

func (s *DispenseState) CancelRequest() {
	// Already dispensing, can't cancel
}

func (s *DispenseState) GetCash() int {
	return s.vm.GetBalanceCash()
}
