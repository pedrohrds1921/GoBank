package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command("tern", "migrate", "--migrations", "./internal/repository/migrations", "--config", "./tern.conf")

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Comand execution failed:", err)
		fmt.Println("Output", string(output))
		panic(err)

	}
	fmt.Println("Comand execution sucessfully", string(output))
}
