package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Client struct {
	threads    []handler
	serverAddr string
	protocol   string
	Err        error
}

//TODO Separate Each Client By Protocol
func (e *Client) SetConnection() {
	jsonFile, err := os.Open("connect.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	e.serverAddr = fmt.Sprintf("%v", result["serverAddress"])
	e.protocol = fmt.Sprintf("%v", result["protocol"])
	log.Print(e)
}

func (e *Client) CreateThreads(values [][]string) {
	for _, value := range values {
		thread := handler{}
		thread.Create(e.serverAddr, value)
		e.threads = append(e.threads, thread)
		log.Print(thread)
	}
}
func (e *Client) MakeTest(messages []string) {
	for _, thread := range e.threads {
		thread.test(messages)
	}
}

func (e *Client) MakeRequest(Message string) {
	for _, thread := range e.threads {
		ch := make(chan string, 10)
		time.Sleep(1 * time.Second)
		thread.MakeRequest(Message, ch)
		select {
		case send := <-ch:
			log.Print(send)
		default:
			log.Print("default")
		}
	}
}

func (e *Client) MakeCommunication(Messages []string) {
	for _, Message := range Messages {
		e.MakeRequest(Message)
	}
}
