package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(text string) string {
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	userInput = strings.ReplaceAll(userInput, "\n", "")

	return userInput
}
