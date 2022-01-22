package models

type User struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	DateOfBirth string `json:"dateOfBirth"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

type UserResponse struct {
	Success string `json:"success"`
	Users   []User `json:"users,omitempty"`
}
