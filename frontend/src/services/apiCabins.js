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

export async function uploadImageToS3(image) {
  //todo
  return image;
}

export async function editCabin(cabin) {
  //todo
  return cabin;
}

export async function createCabin(cabin) {
  const imageName = `${Math.random()}-${cabin.image.name}`.replaceAll("/");
  // const imagePath = "s3url/`${imageName}`";
  const formattedCabin = {
    ...cabin,
    maxCapacity: Number(cabin.maxCapacity), // Convert to integer
    regularPrice: Number(cabin.regularPrice), // Convert to float
    discount: Number(cabin.discount), // Convert to float
  };
  const response = await fetch(`${BACKEND_URL}/api/cabins`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(formattedCabin),
  });
  if (!response.ok) {
    const errorText = await response.text();
    console.error("Error creating cabin:", errorText);
    throw new Error("Failed to create cabin");
  }
  const uploadError = await uploadImageToS3(imageName);
  if (uploadError) {
    //todo
    deleteCabin(cabin);
  }

  return await response.json();
}
