package patternmatcher

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"stack/src/entity"
)

func Load_json(filename string) ([]entity.InventorySecret, error) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
		return nil, err
	}

	var inventory []entity.InventorySecret
	err = json.Unmarshal(bytes, &inventory)
	if err != nil {
		log.Fatalf("Failed to unmarshal json: %v", err)
		return nil, err
	}
	return inventory, nil
}
