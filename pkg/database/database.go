package database

import "fmt"

func Save(key int) string {
	message := fmt.Sprintf("Saved with the key: %v", key)
	return message
}
