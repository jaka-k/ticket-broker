import { writable } from "svelte/store";

type Order = {
  id: string;
  countryCode: string;
  orderAmount: number;
  status: 'pending' | 'success' | 'rejected';
};

function createOrderStore() {
    const { subscribe, set, update } = writable([]);
  
    return {
      subscribe,
      logApiCall: (order: Order) => update((orders) => [...orders, order]),
      reset: () => set([])
    };
  }
  
  export const orderStore = createOrderStore();