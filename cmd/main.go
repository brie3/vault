package main

import (
	"fmt"
	"os"
	"time"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	vaultAddr := os.Getenv("VAULT_ADDR")
	dbSecretPath := os.Getenv("DB_SECRET_PATH")

	config := &vault.Config{Address: vaultAddr, Timeout: 20 * time.Second}

	vaultClient, err := vault.NewClient(config)
	if err != nil {
		fmt.Println("can't create a new vault client instance: %w", err)
		return
	}

	var ok bool
	var dbURL map[string]interface{}
	for attempt := time.Duration(1); attempt < 6; attempt++ {
		dbSecret, err := vaultClient.Logical().Read(dbSecretPath)
		if err != nil {
			fmt.Printf("can't read secret from vault: %v", err)
			return
		}

		if dbURL, ok = dbSecret.Data["data"].(map[string]interface{}); ok {
			break
		}
		fmt.Printf("attempt #%d - can't get data from vault client: %v\n", attempt, dbSecret)
		time.Sleep(attempt * time.Second)
	}

	if !ok {
		fmt.Printf("can't find db_url key in secret")
	} else {
		fmt.Printf("secret from vault is: %v", dbURL["db_url"])
	}
}
