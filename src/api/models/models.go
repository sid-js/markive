package models

import (
    "gorm.io/gorm"
    "time"
)

type Author struct {
    gorm.Model
    AuthorID uint64 `gorm:"primaryKey"`
    Name     string `json:"name"`
    Bio      string `json:"bio"`
    Username string `json:"username" gorm:"unique"`
    Blogs    []Blog `json:"blogs"`
}

type Blog struct {
    gorm.Model
    BlogID    uint64    `gorm:"primaryKey"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    AuthorID  uint      `json:"-"`
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
