package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		log.Println(scanner.Text())
		line := scanner.Bytes()

		var msg Message
		if err := json.Unmarshal(line, &msg); err != nil {
			log.Fatal(err)
		}

		switch msg.Method {
		case MethodNone:
			var res Response
			if err := json.Unmarshal(line, &res); err != nil {
				log.Fatal(err)
			}

			log.Println("Response", res)
			handleResponse(&res)
		case MethodPing:
			log.Println("ping")
		case MethodPingFromEditor:
			log.Println("ping_from_editor", msg.Params)
			send(&Request{0, MethodNLines, nil}) // []struct{}{}})
		}
	}
}

func send(req *Request) {
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	data = append(data, []byte("\n")...)
	os.Stdout.Write(data)
}

func handleResponse(res *Response) {
	// ...
}
