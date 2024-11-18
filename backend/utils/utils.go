package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskQuestion(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	return strings.TrimSpace(answer)
}

func AskYesNo(prompt string) bool {
	for {
		answer := AskQuestion(prompt)
		answer = strings.ToLower(answer)
		if answer == "y" || answer == "yes" {
			return true
		} else if answer == "n" || answer == "no" {
			return false
		}
	}
}
