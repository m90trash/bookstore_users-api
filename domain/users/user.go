package users

type User struct {
	Id         int    `json:"id"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	CreateDate string `json:"create_date"`
}
