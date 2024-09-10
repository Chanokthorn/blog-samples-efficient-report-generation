package domain

type Job struct {
	ID     string `json:"id"`
	Done   bool   `json:"done"`
	Report Report `json:"report"`
}
