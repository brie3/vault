package main

import (
	"fmt"
	"os"
	"time"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	vaultAddr := os.Getenv("VAULT_ADDR")
	vaultToken := os.Getenv("VAULT_TOKEN")
	dbSecretPath := os.Getenv("DB_SECRET_PATH")

	fmt.Printf("addr: %s, token: %s, path: %s\n", vaultAddr, vaultToken, dbSecretPath)

	config := vault.Config{Timeout: 5 * time.Second}
	fmt.Println(config.Address)
}
