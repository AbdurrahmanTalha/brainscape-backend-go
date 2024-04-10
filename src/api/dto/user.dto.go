package dto

type UpdateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password []byte `gorm:"type:bytea" json:"password"`
	Role     string `json:"role"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
