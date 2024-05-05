import { generateId } from "../utils";

export default async function performApiCalls(
  count: number,
  origin: string
): Promise<void> {
  console.log(count, origin);


  const url = `http://localhost:3000/buy`;

  for (let i = 0; i < count; i++) {
   
    fetch(url, {
        method: 'POST',
        body: JSON.stringify({
          id:  generateId(origin),
          countryCode: origin,
          orderAmount: count,
        }),
      })
      .then(async (response) => response.json())
      .then((data) => console.log(data))
      .catch((error) => console.error('Error:', error));
  }
}
