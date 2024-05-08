<script lang="ts">
  import { onMount } from 'svelte';
  import { orderStore } from '../store';
  type OrderKeys = keyof Order;
  let orders: Order[] = [];

  // Subscription to the store
  $: $orderStore, (orders = $orderStore);

  //TODO: REMOVE INITITIAL ORDER
  // Function to add a new order
  function addOrder() {
    orderStore.logApiCall({
      id: 'new-order-id', // This should be unique
      countryCode: 'US',
      orderAmount: 123,
      status: 'pending',
    });
  }

  console.log(orders);

  // Fetch initial data or setup component
  onMount(() => {
    addOrder(); // Call this to add an order when the component mounts
  });

  // Headers should ideally be static or calculated once if dynamic
  let headers: OrderKeys[] = ['id', 'countryCode', 'orderAmount', 'status'];
</script>

<div class="card-body items-center text-center">
  <h2 class="card-title">Orders</h2>
  <div class="overflow-x-auto text-secondary max-h-96">
    <table class="table table-xs table-pin-rows">
      <thead>
        <tr>
          {#each headers as header}
            <th>{header}</th>
          {/each}
        </tr>
      </thead>
      <tbody>
        {#each orders as row (row.id)}
          <tr>
            {#each headers as header}
              <td>{row[header]}</td>
            {/each}
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>
