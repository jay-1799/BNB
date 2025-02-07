package repository

import (
	"database/sql"
	"errors"
	"time"

	"backend/internal/models"
)

// Repository holds a reference to the DB connection.
type Repository struct {
	DB *sql.DB
}

// ---------------------
// Guest Repository Functions
// ---------------------

func (r *Repository) CreateGuest(g *models.Guest) error {
	query := `
		INSERT INTO guests (full_name, email, nationality, national_id, country_flag)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	return r.DB.QueryRow(query, g.FullName, g.Email, g.Nationality, g.NationalID, g.CountryFlag).Scan(&g.ID)
}

func (r *Repository) GetGuestByID(id int) (*models.Guest, error) {
	g := &models.Guest{}
	query := `
		SELECT id, full_name, email, nationality, national_id, country_flag
		FROM guests
		WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&g.ID, &g.FullName, &g.Email, &g.Nationality, &g.NationalID, &g.CountryFlag)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("guest not found")
		}
		return nil, err
	}
	return g, nil
}

func (r *Repository) GetAllGuests() ([]models.Guest, error) {
	query := `SELECT id, full_name, email, nationality, national_id, country_flag FROM guests`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guests []models.Guest
	for rows.Next() {
		var g models.Guest
		if err := rows.Scan(&g.ID, &g.FullName, &g.Email, &g.Nationality, &g.NationalID, &g.CountryFlag); err != nil {
			return nil, err
		}
		guests = append(guests, g)
	}
	return guests, nil
}

// ---------------------
// Cabin Repository Functions
// ---------------------

func (r *Repository) CreateCabin(c *models.Cabin) error {
	query := `
		INSERT INTO cabins (name, max_capacity, regular_price, discount, image, description)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	return r.DB.QueryRow(query, c.Name, c.MaxCapacity, c.RegularPrice, c.Discount, c.Image, c.Description).Scan(&c.ID)
}

func (r *Repository) GetCabinByID(id int) (*models.Cabin, error) {
	c := &models.Cabin{}
	query := `
		SELECT id, name, max_capacity, regular_price, discount, image, description
		FROM cabins
		WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.MaxCapacity, &c.RegularPrice, &c.Discount, &c.Image, &c.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("cabin not found")
		}
		return nil, err
	}
	return c, nil
}

func (r *Repository) GetAllCabins() ([]models.Cabin, error) {
	query := `SELECT id, name, max_capacity, regular_price, discount, image, description FROM cabins`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cabins []models.Cabin
	for rows.Next() {
		var c models.Cabin
		if err := rows.Scan(&c.ID, &c.Name, &c.MaxCapacity, &c.RegularPrice, &c.Discount, &c.Image, &c.Description); err != nil {
			return nil, err
		}
		cabins = append(cabins, c)
	}
	return cabins, nil
}

// ---------------------
// Booking Repository Functions
// ---------------------

func (r *Repository) CreateBooking(b *models.Booking) error {
	query := `
		INSERT INTO bookings (created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`
	if b.CreatedAt.IsZero() {
		b.CreatedAt = time.Now()
	}
	return r.DB.QueryRow(
		query,
		b.CreatedAt,
		b.StartDate,
		b.EndDate,
		b.CabinID,
		b.GuestID,
		b.HasBreakfast,
		b.Observations,
		b.IsPaid,
		b.NumGuests,
	).Scan(&b.ID)
}

func (r *Repository) GetBookingByID(id int) (*models.Booking, error) {
	b := &models.Booking{}
	query := `
		SELECT id, created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests
		FROM bookings
		WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(
		&b.ID,
		&b.CreatedAt,
		&b.StartDate,
		&b.EndDate,
		&b.CabinID,
		&b.GuestID,
		&b.HasBreakfast,
		&b.Observations,
		&b.IsPaid,
		&b.NumGuests,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}
	return b, nil
}

func (r *Repository) GetAllBookings() ([]models.Booking, error) {
	query := `
		SELECT id, created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests
		FROM bookings`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(&b.ID, &b.CreatedAt, &b.StartDate, &b.EndDate, &b.CabinID, &b.GuestID, &b.HasBreakfast, &b.Observations, &b.IsPaid, &b.NumGuests); err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

func (r *Repository) GetBookings(after, before string) ([]models.Booking, error) {
	query := `
        SELECT id, created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests
        FROM bookings`

	// If filtering is requested, modify the query.
	if after != "" && before != "" {
		query += ` WHERE created_at BETWEEN $1 AND $2`
		rows, err := r.DB.Query(query, after, before)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var bookings []models.Booking
		for rows.Next() {
			var b models.Booking
			if err := rows.Scan(&b.ID, &b.CreatedAt, &b.StartDate, &b.EndDate, &b.CabinID, &b.GuestID, &b.HasBreakfast, &b.Observations, &b.IsPaid, &b.NumGuests); err != nil {
				return nil, err
			}
			bookings = append(bookings, b)
		}
		return bookings, nil
	}

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(&b.ID, &b.CreatedAt, &b.StartDate, &b.EndDate, &b.CabinID, &b.GuestID, &b.HasBreakfast, &b.Observations, &b.IsPaid, &b.NumGuests); err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

func (r *Repository) UpdateBooking(b *models.Booking) error {
	query := `
        UPDATE bookings
        SET start_date = $1,
            end_date = $2,
            cabin_id = $3,
            guest_id = $4,
            has_breakfast = $5,
            observations = $6,
            is_paid = $7,
            num_guests = $8
        WHERE id = $9`

	_, err := r.DB.Exec(query, b.StartDate, b.EndDate, b.CabinID, b.GuestID, b.HasBreakfast, b.Observations, b.IsPaid, b.NumGuests, b.ID)
	return err
}

func (r *Repository) DeleteBooking(id int) error {
	query := `DELETE FROM bookings WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *Repository) GetBookingsForDate(date string) ([]models.Booking, error) {
	query := `
        SELECT id, created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests
        FROM bookings
        WHERE start_date = $1 OR end_date = $1`

	rows, err := r.DB.Query(query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(&b.ID, &b.CreatedAt, &b.StartDate, &b.EndDate, &b.CabinID, &b.GuestID, &b.HasBreakfast, &b.Observations, &b.IsPaid, &b.NumGuests); err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}
