export const apiOrderCall = async ({
  url,
  id,
  origin,
  numOfTickets,
}: {
  url: string;
  id: string;
  origin: string;
  numOfTickets: number;
}) =>
  fetch(url, {
    method: 'POST',
    body: JSON.stringify({
      id,
      countryCode: origin,
      numOfTickets,
    }),
    headers: {
      'Content-Type': 'application/json',
    },
  });
