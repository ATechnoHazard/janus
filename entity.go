package janus

// Account represents a user account
type Account struct {
	OrganizationID uint   `json:"org_id" gorm:"primary_key"`
	CacheKey       string `json:"key" gorm:"primary_key"`
	Role           string `json:"role"`
}
