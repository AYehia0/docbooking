package model

import (
	"docbooking/pkg/uuid"
	"time"
)

type Slot struct {
	ID         uuid.UUID `json:"id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	IsReserved bool      `json:"is_reserved"`
	DoctorID   uuid.UUID `json:"doctor_id"`
	DoctorName string    `json:"doctor_name"`
	Cost       float64   `json:"cost"`
}
