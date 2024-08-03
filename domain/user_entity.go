package domain

type User struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	UserEmail string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
}
