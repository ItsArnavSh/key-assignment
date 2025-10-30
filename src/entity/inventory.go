package entity

// InventorySecret represents each secret token entry
type InventorySecret struct {
	Provider    string `json:"provider"`
	TokenType   string `json:"token_type"`
	Value       string `json:"value"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
}
