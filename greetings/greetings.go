package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// 返回在消息中嵌入名称的问候语.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
