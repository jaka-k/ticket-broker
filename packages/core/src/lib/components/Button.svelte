<script>
  // Prop to accept an event handler
  export let onClick = () => {}; 
  export let label = 'Click me';
</script>

<button class="btn btn-primary p-28" on:click={onClick}>
  {label}
</button>
