const BACKEND_URL = "http://localhost:8080";

export async function getCabins() {
  const response = await fetch(`${BACKEND_URL}/api/cabins`);
  if (!response.ok) {
    console.error("Error fetching cabins:", response.statusText);
    throw new Error("Booking not found");
  }
  const data = await response.json();
  return data;
}
