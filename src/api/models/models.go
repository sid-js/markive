package models

import (
    "gorm.io/gorm"
    "time"
)

type Author struct {
    gorm.Model
    Name     string `json:"name"`
    Bio      string `json:"bio"`
    Username string `json:"username" gorm:"unique"`
}

type Blog struct {
    gorm.Model
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Author    Author    `json:"author"`
}

func (b *Blog) BeforeCreate(tx *gorm.DB) (err error) {
    b.CreatedAt = time.Now()
    b.UpdatedAt = time.Now()
    return nil
}

func (b *Blog) BeforeUpdate(tx *gorm.DB) (err error) {
    b.UpdatedAt = time.Now()
    return nil
}
