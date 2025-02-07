package models

import "time"

type Guest struct {
	ID          int    `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	Nationality string `json:"nationality"`
	NationalID  string `json:"nationalID"`
	CountryFlag string `json:"countryFlag"`
}

type Cabin struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	MaxCapacity  int     `json:"maxCapacity"`
	RegularPrice float64 `json:"regularPrice"`
	Discount     float64 `json:"discount"`
	Image        string  `json:"image"`
	Description  string  `json:"description"`
}

type Booking struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	CabinID      int       `json:"cabinId"`
	GuestID      int       `json:"guestId"`
	HasBreakfast bool      `json:"hasBreakfast"`
	Observations string    `json:"observations"`
	IsPaid       bool      `json:"isPaid"`
	NumGuests    int       `json:"numGuests"`
}

type Settings struct {
	ID                 int     `json:"id"`
	MinBookingLength   int     `json:"minBookingLength"`
	MaxBookingLength   int     `json:"maxBookingLength"`
	MaxGuestPerBooking int     `json:"maxGuestPerBooking"`
	BreakfastPrice     float64 `json:"breakfastPrice"`
}
