import { getToday } from "../utils/helpers";

const BACKEND_URL = "http://localhost:8080";

export async function getBooking(id) {
  const response = await fetch(`${BACKEND_URL}/api/bookings/${id}`);
  if (!response.ok) {
    console.error("Error fetching booking:", response.statusText);
    throw new Error("Booking not found");
  }
  const data = await response.json();
  return data;
}

// Get all bookings created between a given date and today.
export async function getBookingsAfterDate(date) {
  const response = await fetch(
    `${BACKEND_URL}/api/bookings?after=${encodeURIComponent(
      date
    )}&before=${encodeURIComponent(getToday({ end: true }))}`
  );
  if (!response.ok) {
    console.error("Error fetching bookings:", response.statusText);
    throw new Error("Bookings could not be loaded");
  }
  const data = await response.json();
  return data;
}

// Get all stays that start between a given date and today.
export async function getStaysAfterDate(date) {
  // If stays are part of the bookings table, you might have a separate endpoint or use the same one with different filtering.
  const response = await fetch(
    `${BACKEND_URL}/api/bookings?staysAfter=${encodeURIComponent(
      date
    )}&staysBefore=${encodeURIComponent(getToday())}`
  );
  if (!response.ok) {
    console.error("Error fetching stays:", response.statusText);
    throw new Error("Bookings could not be loaded");
  }
  const data = await response.json();
  return data;
}

// Get today's activity (e.g., check-ins or check-outs).
export async function getStaysTodayActivity() {
  const today = getToday();
  const response = await fetch(
    `${BACKEND_URL}/api/stays/today-activity?date=${encodeURIComponent(today)}`
  );
  if (!response.ok) {
    console.error("Error fetching today's activity:", response.statusText);
    throw new Error("Bookings could not be loaded");
  }
  const data = await response.json();
  return data;
}

export async function updateBooking(id, obj) {
  const response = await fetch(`${BACKEND_URL}/api/bookings/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(obj),
  });
  if (!response.ok) {
    console.error("Error updating booking:", response.statusText);
    throw new Error("Booking could not be updated");
  }
  const data = await response.json();
  return data;
}

export async function deleteBooking(id) {
  const response = await fetch(`${BACKEND_URL}/api/bookings/${id}`, {
    method: "DELETE",
  });
  if (!response.ok) {
    console.error("Error deleting booking:", response.statusText);
    throw new Error("Booking could not be deleted");
  }
  const data = await response.json();
  return data;
}
