package main

import (
	"fmt"

	"github.com/gofp/pkg/singleton"
)

func main() {
	store := singleton.NewSecret()

	fmt.Printf("Secret %s = %s\n", "Osaka", store.Value("Osaka"))
}
