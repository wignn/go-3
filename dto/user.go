package dto


type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	IComplated bool `json:"isComplated"`
}