package main

import (
	"fmt"
	"vending_machine/internal"
	"vending_machine/internal/enums"
	"vending_machine/internal/models"
	"vending_machine/internal/state"
)

func main() {
	// Create items
	items := []models.Item{
		{ItemId: 1, Type: enums.Beverages, Price: 30},
		{ItemId: 2, Type: enums.Eatables, Price: 100},
		{ItemId: 3, Type: enums.Beverages, Price: 50},
	}

	// Create a new vending machine using the constructor
	vm := internal.NewVendingMachine(items, 0)

	// Start in IdleState
	vm.SetState(state.NewIdleState(vm))

	fmt.Println("=== Vending Machine Demo ===")
	fmt.Printf("Initial items: %d\n", len(vm.GetItems()))
	fmt.Printf("Initial balance: $%d\n", vm.GetBalanceCash())

	// Insert cash
	fmt.Println("\n1. Inserting $50...")
	balance := vm.AcceptCash(50)
	fmt.Printf("Current balance: $%d\n", balance)

	// Checkout
	fmt.Println("\n2. Checking out...")
	item := vm.Checkout()
	if item != nil {
		fmt.Printf("Dispensed item: ID=%d, Type=%v\n", item.ItemId, item.Type)
	}

	// Dispense cash (change)
	fmt.Println("\n3. Dispensing change...")
	change := vm.State.DispenseCash()
	fmt.Printf("Returned change: $%d\n", change)

	fmt.Printf("\nFinal items: %d\n", len(vm.GetItems()))
	fmt.Printf("Final balance: $%d\n", vm.GetBalanceCash())
}
