const BACKEND_URL = "http://localhost:8080";

export async function getSettings() {
  const response = await fetch(`${BACKEND_URL}/api/settings`);
  if (!response.ok) {
    const errorText = await response.text();
    console.error("Error fetching settings:", errorText);
    throw new Error("Failed to fetch settings");
  }
  const data = await response.json();
  return data;
}

export async function updateSettings(settings) {
  const response = await fetch(`${BACKEND_URL}/api/settings`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(settings),
  });
  if (!response.ok) {
    const errorText = await response.text();
    console.error("Error updating settings:", errorText);
    throw new Error("Failed to update settings");
  }
  return await response.json();
}
