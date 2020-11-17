package main

import "time"

type Queue struct {
	ID              int       `gorm:"column:id",json:"id"`
	UserId          int       `gorm:"column:user_id",json:"user_id"`
	Status          string    `gorm:"column:status",json:"status"`
	InviteTo        string    `gorm:"clumn:invite_to",json:"invite_to"`
	InviteCode      string    `gorm:"clumn:invite_code",json:"invite_code"`
	NodeWorkerId    string    `gorm:"clumn:node_worker_id",json:"node_worker_id"`
	GameInviteLevel int       `gorm:"column:game_invite_level",json:"game_invite_level"`
	TotalSpin       int       `gorm:"column:total_spin",json:"total_spin"`
	CurrentSpin     int       `gorm:"column:current_spin",json:"current_spin"`
	DoneAt          time.Time `gorm:"column:created_at",json:"done_at"`
	CreatedAt       time.Time `gorm:"column:created_at",json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at",json:"updated_at"`
}


