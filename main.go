package main // import "github.com/hashicorp/vault"

import (
	"os"
	"fmt"
	"github.com/hashicorp/vault/cli"
)

func main() {
	fmt.Println("Hello")
	os.Exit(cli.Run(os.Args[1:]))
}
