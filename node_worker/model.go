package main

import "time"

type NodeWorker struct {
	ID          int       `gorm:"column:id",json:"id"`
	Label       string    `gorm:"column:label",json:"label"`
	TargetUrl   string    `gorm:"column:target_url",json:"target_url"`
	InTask      int       `gorm:"column:in_task",json:"in_task"`
	MaximumTask int       `gorm:"column:maximum_task",json:"maximum_task"`
	Enabled     bool      `gorm:"column:enabled",json:"enabled"`
	CreatedAt   time.Time `gorm:"column:created_at",json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at",json:"updated_at"`
}

func (*NodeWorker) TableName() string {
	return "node_workers"
}
