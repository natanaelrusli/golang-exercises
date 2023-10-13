package robotsimulator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type RobotPosition struct {
	X      int
	Y      int
	Facing int
}

func (robot *RobotPosition) GetInitialPosition(initialPosition string) RobotPosition {
	facingDirections := map[string]int{
		"n": 1,
		"e": 2,
		"s": 3,
		"w": 4,
	}

	trimmedInitialPosition := strings.ReplaceAll(initialPosition, " ", "")
	position := strings.Split(trimmedInitialPosition, "")

	xPosition, err := strconv.Atoi(position[0])

	if err != nil {
		log.Fatalln("error converting string to int", err.Error())
	}

	yPosition, err := strconv.Atoi(position[1])

	if err != nil {
		log.Fatalln("error converting string to int", err.Error())
	}

	facing := strings.ToLower(position[2])

	robot.X = xPosition
	robot.Y = yPosition
	robot.Facing = facingDirections[facing]

	return *robot
}

func (robot *RobotPosition) MoveRobot(movements string) RobotPosition {
	movementArr := strings.Split(movements, "")

	for _, v := range movementArr {
		switch v {
		case "A":
			switch robot.Facing {
			case 1:
				robot.Y += 1
			case 2:
				robot.X += 1
			case 3:
				robot.Y -= 1
			case 4:
				robot.X -= 1
			}

		case "R":
			if robot.Facing == 4 {
				robot.Facing = 1
			} else {
				robot.Facing += 1
			}

		case "L":
			if robot.Facing == 1 {
				robot.Facing = 4
			} else {
				robot.Facing -= 1
			}
		}
	}

	return *robot
}

func (robot *RobotPosition) GetFinalPosition() string {
	facingDirectionsInt := map[int]string{
		1: "N",
		2: "E",
		3: "S",
		4: "W",
	}

	result := fmt.Sprintf("%v %v %s", robot.X, robot.Y, facingDirectionsInt[robot.Facing])

	return result
}
