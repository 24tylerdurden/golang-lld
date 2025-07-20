package state

import (
	"vending_machine/internal/contracts"
	"vending_machine/internal/models"
)

type IdleState struct {
	vm contracts.VMContext
}

func NewIdleState(vm contracts.VMContext) *IdleState {
	return &IdleState{vm: vm}
}

func (s *IdleState) AcceptCash(amount int) int {
	s.vm.SetBalanceCash(s.vm.GetBalanceCash() + amount)
	// Transition to ProcessingState
	s.vm.SetState(NewProcessingState(s.vm))
	return s.vm.GetBalanceCash()
}

func (s *IdleState) ValidateRequest() bool {
	return false
}

func (s *IdleState) DispenseProduct() *models.Item {
	return nil
}

func (s *IdleState) DispenseCash() int {
	return 0
}

func (s *IdleState) CancelRequest() {
	// Reset to idle state
}

func (s *IdleState) GetCash() int {
	return s.vm.GetBalanceCash()
}
