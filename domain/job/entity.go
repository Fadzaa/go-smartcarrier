package job

type Job struct {
	//gorm.Model
	ID          int8   `json:"id" gorm:"primaryKey"`
	CompanyName string `json:"company_name" gorm:"type:varchar(255)"`
	Jobdesk     string `json:"jobdesk" gorm:"type:varchar(100)"`
	Location    string `json:"location" gorm:"type:varchar(100)"`
	//JobType      int8      `json:"job_type"`
	CompanyImage string `json:"company_image"`
	//CreatedAt    time.Time `json:"created_at"`
	//UpdatedAt    time.Time `json:"updated_at"`
}
