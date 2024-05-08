<script lang="ts">
  import { performApiCalls } from "..";

  // Prop to accept an event handler
  export let countryCode = 'gb';
  export let label = 'Tickets';


  export let handleClick = (orderAmount: number) => performApiCalls(orderAmount, countryCode); 
</script>

<div class="w-40 mx-auto flex items-center justify-between gap-4">

  <button class="btn btn-secondary" on:click={() => handleClick(10)}>
    10 {label}
  </button>
  <button class="btn btn-accent" on:click={() => handleClick(100)}>
    100 {label}
  </button>
</div>
