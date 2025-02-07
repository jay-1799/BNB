package main

import (
	"backend/internal/db"
	"backend/internal/handlers"
	"backend/internal/repository"
	"fmt"
	"log"
	"net/http"

	// "fyne.io/fyne/v2/storage/repository"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	db := db.InitDB()
	defer db.Close()

	// Create repository instance.
	repo := &repository.Repository{DB: db}

	// Initialize handlers with dependency injection.
	guestHandler := &handlers.GuestHandler{Repo: repo}
	cabinHandler := &handlers.CabinHandler{Repo: repo}
	bookingHandler := &handlers.BookingHandler{Repo: repo}

	r := mux.NewRouter()

	r.HandleFunc("/api/guests", guestHandler.GetAllGuests).Methods("GET")
	r.HandleFunc("/api/guests/{id:[0-9]+}", guestHandler.GetGuestByID).Methods("GET")

	r.HandleFunc("/api/cabins", cabinHandler.GetAllCabins).Methods("GET")
	r.HandleFunc("/api/cabins/{id:[0-9]+}", cabinHandler.GetCabinByID).Methods("GET")

	r.HandleFunc("/api/bookings", bookingHandler.GetAllBookings).Methods("GET")
	r.HandleFunc("/api/bookings/{id:[0-9]+}", bookingHandler.GetBookingByID).Methods("GET")
	r.HandleFunc("/api/bookings/{id:[0-9]+}", bookingHandler.UpdateBooking).Methods("PUT")
	r.HandleFunc("/api/bookings/{id:[0-9]+}", bookingHandler.DeleteBooking).Methods("DELETE")

	r.HandleFunc("/api/stays/today-activity", bookingHandler.GetTodaysActivity).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5174"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	fmt.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
