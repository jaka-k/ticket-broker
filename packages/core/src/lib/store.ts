import { writable } from 'svelte/store';


const orders = writable<Order[]>([]);
const { subscribe, set, update } = orders


function createOrder({ id, countryCode, orderAmount, status }: Order) {
  update((currentOrders: Order[]) => {
    const newOrder: Order = { id, countryCode, orderAmount, status };
    return [...currentOrders, newOrder];
  });
}

export function updateOrderStatus(orderID: string, newStatus: Status) {
  update((allOrders) => {
    const index: number = allOrders.findIndex((order) => order.id === orderID);
    const order = allOrders[index];

    if (!order) {
      console.warn('Order not found for updating:', orderID);
      return allOrders;
    }

    if (index !== -1) {
      const updatedOrder: Order = {
        id: order.id,
        countryCode: order.countryCode,
        orderAmount: order.orderAmount,
        status: newStatus,
      };
      return allOrders.map((order, idx) =>
        idx === index ? updatedOrder : order
      );
    }
    return allOrders;
  });
}

function createOrderStore() {
  return {
    subscribe,
    logApiCall: (order: Order) => createOrder(order),
    updateOrderStatus,
    reset: () => set([]),
  };
}

export const orderStore = createOrderStore();
