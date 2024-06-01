import { orderStore } from '../store';

import { apiOrderCall } from '../utils/api/apiOrderCall';
import { generateId } from '../utils/generateId';
import { getWeightedRandom } from '../utils/getWeightedRandom';

export default async function performApiCalls(
  numOfOrders: number,
  origin: string
): Promise<void> {
  const url = `http://localhost:3000/buy`;

  for (let i = 0; i < numOfOrders; i++) {
    const id = generateId(origin);
    const numOfTickets = getWeightedRandom() || 2

    try {
      const response = await apiOrderCall({ url, id, origin, numOfTickets });

      const data = await response.json();
      if (response.ok) {
        orderStore.logApiCall({
          id: id,
          countryCode: origin,
          orderAmount: numOfTickets,
          status: 'pending', // Assume pending until confirmed otherwise
        });
        // Here you might want to update status based on server response if needed
      } else {
        orderStore.updateOrderStatus(id, 'rejected');
      }
    } catch (error) {
      console.error('Error:', error);
      orderStore.updateOrderStatus(id, 'canceled');
    }
  }
}
