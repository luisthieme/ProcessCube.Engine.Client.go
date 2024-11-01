package main

import client "github.com/luisthieme/ProcessCube.Engine.Client.go"

func main() {
	// identity := &client.Identity{ Token: "ZHVtbXlfdG9rZW4=", UserId: "dummy_token"}
	client := client.NewClient("http://localhost:56100")

	client.ProcessModels.GetAll()
}
