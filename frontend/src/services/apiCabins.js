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

export async function deleteCabin(id) {
  const response = await fetch(`${BACKEND_URL}/api/cabins/${id}`, {
    method: "DELETE",
  });
  if (!response.ok) {
    console.error("Error deleting cabin:", response.statusText);
    throw new Error("Cabin could not be deleted");
  }
  const data = await response.json();
  return data;
}

export async function createCabin(cabin) {
  const response = await fetch(`${BACKEND_URL}/api/cabins`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(cabin),
  });
  if (!response.ok) {
    const errorText = await response.text();
    console.error("Error creating cabin:", errorText);
    throw new Error("Failed to create cabin");
  }
  return await response.json();
}
