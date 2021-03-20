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
	url := os.Args[1]
	msg := os.Args[2:]

	if argsLen == 1 {
		fmt.Println("Please provide webhook link, the message and try again.")
		fmt.Println("Example: dishook https://discord.com/api/webhooks/.../.../ Hello World!")
		os.Exit(0)
		// Checks if there's no URL in first argument
	}
	if argsLen == 2 {
		isDiscordLink(url)
		fmt.Println("Please put in the message and try again.")
		os.Exit(0)
		// Check if there's no content in second+ argument
	}

	// Checks if webhook is valid and has no errors whatsoever.
	isDiscordLink(url)
	isTokenInvalid(url)

	// Process message argument
	var msgBuf string
	for i := 0; i < len(msg); i++ {
		msgBuf = msgBuf + " " + msg[i]
	}
	msgBuf = strings.TrimSpace(msgBuf)
	// fmt.Println(msgBuf)

	// Sends request to webhook
	sendMsg(url, msgBuf)
}

// functions time
func sendMsg(url string, content string) {
	values := map[string]string{"content": content}
	jsonValue, _ := json.Marshal(values)
	// Turns content string into JSON
	
	resp, err := requests.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	fmt.Println("Yeahh")
	fmt.Print(resp)
	manageError(err)
	// Sends request to webhook
}

func isDiscordLink(url string) bool {
	if url[0:33] == "https://discord.com/api/webhooks/" {
		return true
	} else {
		fmt.Println("Please provide a valid webhook URL and try again.")
		os.Exit(0)
	}
	return false
}

func isMsgMax(msg string) bool {
	msgLen := len(msg)
	msgLimit := 2000
	if msgLen < msgLimit {
		return false
	} else {
		msgToShort := msgLen - msgLimit
		fmt.Printf("Your message's length (%d) surpasses Discord's limit (%d). Please make it %d characters shorter and try again.", 
		msgLen, msgLimit, msgToShort)

		return true
	}
}

func isTokenInvalid(url string) {
	url_r, err := requests.Get(url)
	manageError(err)

	url_code := url_r.StatusCode // thank you discord for putting the invalid error in the url's status code
	if url_code == 401 {
		fmt.Println("Invalid Webhook Token. Please try again")
		os.Exit(0)
	}
}

func manageError(err error) {
	// just to calm down with the syntax
	if err != nil {
		fmt.Println("An unexpected error ocurred. Please try again.")
		panic(err)
	}
}