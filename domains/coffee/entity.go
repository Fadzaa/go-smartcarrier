package coffee

import "time"

type Coffee struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Origin    string    `json:"origin"`
	CreatedAt time.Time `json:"created_at"`
}
