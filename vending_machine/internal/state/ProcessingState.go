package state

import (
	"vending_machine/internal/contracts"
	"vending_machine/internal/models"
)

type ProcessingState struct {
	vm contracts.VMContext
}

func NewProcessingState(vm contracts.VMContext) *ProcessingState {
	return &ProcessingState{vm: vm}
}

func (s *ProcessingState) AcceptCash(amount int) int {
	s.vm.SetBalanceCash(s.vm.GetBalanceCash() + amount)
	return s.vm.GetBalanceCash()
}

func (s *ProcessingState) ValidateRequest() bool {
	// For demo, always return true if we have items
	items := s.vm.GetItems()
	return len(items) > 0
}

func (s *ProcessingState) DispenseProduct() *models.Item {
	items := s.vm.GetItems()
	if len(items) == 0 {
		return nil
	}
	
	// Get the first item
	item := items[0]
	// Remove the first item
	s.vm.SetItems(items[1:])
	
	// Transition to DispenseState
	s.vm.SetState(NewDispenseState(s.vm))
	
	return &item
}

func (s *ProcessingState) DispenseCash() int {
	return 0
}

func (s *ProcessingState) CancelRequest() {
	// Transition back to IdleState
	s.vm.SetState(NewIdleState(s.vm))
}

func (s *ProcessingState) GetCash() int {
	return s.vm.GetBalanceCash()
}
