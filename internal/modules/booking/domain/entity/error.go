package entity

import "errors"

// the errors that can be returned by the appointment entity.
var (
	ErrAppointmentNotFound      = errors.New("appointment not found")
	ErrAppointmentAlreadyExists = errors.New("appointment already exists")
)
