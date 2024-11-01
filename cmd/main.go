package main

import (
	"fmt"
	"log"

	client "github.com/luisthieme/ProcessCube.Engine.Client.go/client"
)

func main() {
	// identity := &client.Identity{ Token: "ZHVtbXlfdG9rZW4=", UserId: "dummy_token"}
	client := client.NewClient("http://localhost:56100")

	res, err := client.ProcessModels.GetAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.TotalCount)
}
