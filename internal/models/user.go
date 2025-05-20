package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser() error {
	return nil
}
func FindById(id int) (*User, error) {
	return &User{}, nil
}
