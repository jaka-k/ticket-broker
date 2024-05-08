import { orderStore } from "../store";
import { generateId } from "../utils";

export default async function performApiCalls(count: number, origin: string): Promise<void> {
    const url = `http://localhost:3000/buy`;

    for (let i = 0; i < count; i++) {
        const id = generateId(origin);
        
        try {
            const response = await fetch(url, {
                method: 'POST',
                body: JSON.stringify({
                    id: id,
                    countryCode: origin,
                    orderAmount: count,
                }),
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            const data = await response.json();
            if (response.ok) {
                orderStore.logApiCall({
                    id: id,
                    countryCode: origin,
                    orderAmount: count,
                    status: 'pending', // Assume pending until confirmed otherwise
                });
                // Here you might want to update status based on server response if needed
            } else {
                orderStore.updateOrderStatus(id, "rejected");
            }
        } catch (error) {
            console.error('Error:', error);
            orderStore.updateOrderStatus(id, "canceled");
        }
    }
}