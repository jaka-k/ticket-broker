<script>
    // Import DaisyUI classes for additional styling
    import 'daisyui/dist/full.css';
  
    // Prop to accept an event handler
    export let onClick = () => {};  // Default to an empty function if no handler is provided
    export let label = 'Click me';  // Default button text
  </script>
  
  <button class="btn btn-primary" on:click={onClick}>
    {label}
  </button>