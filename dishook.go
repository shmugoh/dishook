package main

import (
	"fmt";
	"bytes";
	"os";
	"strings";
	"github.com/jochasinga/requests";
	"encoding/json"
)

func main() {
	// Obtains the arguments
	argsLen := len(os.Args)

	if argsLen == 1 {
		fmt.Println("Please insert the webhook link, the message and try again.")
		fmt.Println("Example: dishook https://discord.com/api/webhooks/.../.../ Message-Goes-Here")
		os.Exit(0)
		// Checks if there's no URL in first argument
	}
	if argsLen == 2 {
		fmt.Println("Please put in the message and try again.")
		os.Exit(0)
		// Check if there's no content in second+ argument
	}

	// Process arguments to variables
	url := os.Args[1]
	msg := os.Args[2:]

	// Process message argument
	var msgBuf string
	for i := 0; i < len(msg); i++ {
		msgBuf = msgBuf + " " + msg[i]
	}
	msgBuf = strings.TrimSpace(msgBuf)

	// Sends request to webhook
	sendMsg(url, msgBuf)
}

func sendMsg(url string, content string) {
	values := map[string]string{"content": content}
	jsonValue, _ := json.Marshal(values)
	// Turns content string into JSON
	
	resp, err := requests.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	fmt.Print(resp)
	manageError(err)
	// Sends request to webhook
}

func manageError(err error) {
	// just to calm down with the syntax
	if err != nil {
		fmt.Println("An unexpected error ocurred. Please try again.")
		panic(err)
	}
}