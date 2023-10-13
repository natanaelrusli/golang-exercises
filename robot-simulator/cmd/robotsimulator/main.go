package main

import (
	"fmt"

	"github.com/natanaelrusli/robot-simulator/internal/cli"
	"github.com/natanaelrusli/robot-simulator/internal/robotsimulator"
)

func main() {
	robot := robotsimulator.NewRobot()

	initialPosition := cli.GetUserInput("User Input: ")
	movements := cli.GetUserInput("Movements: ")

	robot.GetInitialPosition(initialPosition)
	robot.MoveRobot(movements)

	result := robot.GetFinalPosition()
	fmt.Println(result)
}
