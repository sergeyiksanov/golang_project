package dto

import "github.com/sergeyiksanov/AuthService/internal/entity"

type CredentialsDto struct {
	ID       int64  `gorm:"column_id:id,primaryKey"`
	Email    string `gorm:"column_id:email,unique"`
	Password string `gorm:"column_id:password"`
}

func (CredentialsDto) TableName() string {
	return "credentials"
}

func (c CredentialsDto) ToCredentialsEntity() entity.Credentials {
	return entity.Credentials{
		ID:       c.ID,
		Email:    c.Email,
		Password: c.Password,
	}
}
