package models

import "time"

type Todo struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`   //主键，自增
	Title     string    `gorm:"type:varchar(100);not null"` //非空
	Status    bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
type TodoResponse struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 转换函数，将Todo转换成TodoResponse
func (t *Todo) ToResponse() TodoResponse {
	return TodoResponse{
		ID:     t.ID,
		Title:  t.Title,
		Status: t.Status,
	}
}
