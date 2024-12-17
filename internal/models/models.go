package models

type User struct {
	UserID        int           `gorm:"column:user_id;primaryKey" json:"user_id"`
	FullName      string        `gorm:"column:full_name" json:"full_name"`
	Email         string        `gorm:"column:email" json:"email"`
	Mobile        string        `gorm:"column:mobile" json:"mobile"`
	Address       string        `gorm:"column:address" json:"address"`
	UserType      string        `gorm:"column:user_type" json:"user_type"`
	CompanyFeilds CompanyFields `gorm:"embedded;embeddedPrefix:company_" json:"company_feilds"`
	CreatedAt     int64         `gorm:"column:created_at" json:"created_at"`
}

type Provider struct {
	ProviderID    int           `gorm:"column:provider_id;primaryKey" json:"provider_id"`
	FullName      string        `gorm:"column:full_name" json:"full_name"`
	Email         string        `gorm:"column:email" json:"email"`
	Mobile        string        `gorm:"column:mobile" json:"mobile"`
	Address       string        `gorm:"column:address" json:"address"`
	ProviderType  string        `gorm:"column:provider_type" json:"provider_type"`
	CompanyFeilds CompanyFields `gorm:"embedded;embeddedPrefix:company_" json:"company_fields"`
	BusinessTaxID string        `gorm:"column:business_tax_id" json:"business_tax_id,omitempty"`
	CreatedAt     int64         `gorm:"column:created_at" json:"created_at"`
}



type CompanyFields struct {
	RepresentTativeName   string `gorm:"column:representative_name" json:"representative_name" validate:"required"`
	RepresentTativeEmail  string `gorm:"column:representative_email" json:"representative_email" validate:"required,email"`
	RepresentTativeMobile string `gorm:"column:representative_mobile" json:"representative_mobile" validate:"required,len=10,numeric"`
	CompanyName           string `gorm:"column:company_name" json:"company_name,omitempty" validate:"required"`
}

type Skill struct {
	SkillID      int     `gorm:"column:skill_id;primaryKey" json:"skill_id"`
	ProviderID   int     `gorm:"column:provider_id;index" json:"provider_id"`
	Category     string  `gorm:"column:category" json:"category"`
	Experience   int     `gorm:"column:experience" json:"experience"`
	NatureOfWork string  `gorm:"column:nature_of_work" json:"nature_of_work"`
	HourlyRate   float64 `gorm:"column:hourly_rate" json:"hourly_rate"`
	CreatedAt    int64   `gorm:"column:created_at" json:"created_at"`
}

type Task struct {
	TaskID        int     `gorm:"column:task_id;primaryKey" json:"task_id"`
	UserID        int     `gorm:"column:user_id;index" json:"user_id"`
	Category      string  `gorm:"column:category" json:"category"`
	TaskName      string  `gorm:"column:task_name" json:"task_name"`
	WorkingHours  int     `gorm:"column:working_hours" json:"working_hours"`
	Description   string  `gorm:"column:description" json:"description"`
	ExpectedStart string  `gorm:"column:expected_start" json:"expected_start"`
	HourlyRate    float64 `gorm:"column:hourly_rate" json:"hourly_rate"`
	RateCurrency  string  `gorm:"column:rate_currency" json:"rate_currency"`
	CreatedAt     string  `gorm:"column:created_at" json:"created_at"`
}

type Offer struct {
	OfferID    int    `gorm:"column:offer_id;primaryKey" json:"offer_id"`
	TaskID     int    `gorm:"column:task_id;index" json:"task_id"`
	ProviderID int    `gorm:"column:provider_id;index" json:"provider_id"`
	Status     string `gorm:"column:status" json:"status"`
	CreatedAt  int64  `gorm:"column:created_at" json:"created_at"`
}

type SqlCondition struct {
	Condition string
	Values    []interface{}
}
