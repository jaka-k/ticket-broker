<script>
  // Prop to accept an event handler
  export let onClick = () => {}; 
  export let label = 'Click me';
</script>

<button class="btn btn-primary" on:click={onClick}>
  {label}
</button>
