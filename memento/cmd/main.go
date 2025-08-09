package main

import "fmt"

func main() {
	careTacker := NewCareTacker()
	currState := NewOriginator(careTacker)

	currState.AddText("Adding First State : Pavan Illa ")
	currState.SavePoint()

	currState.AddText("Adding Second State : Teja Illa ")
	currState.SavePoint()

	currState.AddText("Adding Third State : Prasad Rao")
	// currState.SavePoint()

	currState.RestoreSavePoint()
	currState.RestoreSavePoint()
	// currState.RestoreSavePoint()

	fmt.Println("The curr text state is : \n", currState.GetState())
}

// we will test the memento design pattern
