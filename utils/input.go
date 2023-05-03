package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func GetInput(msg string) (string, error) {
	if msg == "" {
		return "", errors.New("invalid msg")
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	v, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return v, nil
}
