package main

type MyEvent struct {
	Username string `json:"username"`
}

func HandleRequest(event MyEvent) error{
	if event.Username == "" {

	}
	return nil
}

func main() {

}