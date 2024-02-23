package auth

type User struct {
	ID       int8   `json:"id"`
	Name     string `json:"name" gorm:"size:255" gorm:"not null"`
	Email    string `json:"email" gorm:"size:100"`
	Password string `json:"password"`
	//PhoneNumber string `json:"phone_number" gorm:"size:50"`
	//Role        string `json:"role" gorm:"size:50"`
	//Gender      int8   `json:"gender"`
	//Province    string `json:"province"`
	//City        string `json:"city"`
}
