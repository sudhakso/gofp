package main

import (
	"fmt"

	"github.com/gofp/pkg/singleton"
)

func main() {
	store := singleton.NewSecret()

	// Voilating DIP - but Ok,
	fmt.Printf("Secret %s = %s\n", "Osaka", store.Value("Osaka"))

	// With DIP considered
	singleton.GetSecretValue(store, "Osaka")
}
