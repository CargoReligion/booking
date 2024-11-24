package service

import "fmt"

type ErrSlotNotFound struct {
	SlotID string
}

func (e *ErrSlotNotFound) Error() string {
	return fmt.Sprintf("slot with ID %s not found", e.SlotID)
}

type ErrSlotAlreadyBooked struct {
	SlotID string
}

func (e *ErrSlotAlreadyBooked) Error() string {
	return fmt.Sprintf("slot with ID %s is already booked", e.SlotID)
}

type ErrPastSlot struct {
	SlotID string
}

func (e *ErrPastSlot) Error() string {
	return fmt.Sprintf("slot with ID %s is in the past and cannot be booked", e.SlotID)
}

type ErrNotStudent struct {
	UserID string
}

func (e *ErrNotStudent) Error() string {
	return fmt.Sprintf("user with ID %s is not a student and cannot book slots", e.UserID)
}

type ErrOverlappingBooking struct {
	StudentID string
}

func (e *ErrOverlappingBooking) Error() string {
	return fmt.Sprintf("student with ID %s has an overlapping booking", e.StudentID)
}

type ErrNotCoach struct {
	UserID string
}

func (e *ErrNotCoach) Error() string {
	return fmt.Sprintf("user with ID %s is not a coach and cannot view upcoming slots", e.UserID)
}
