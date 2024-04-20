<script>
    import '../app.css';
  export let theme = 'lemonade';
</script>

<div class={`theme-${theme}`}>
  <slot />
</div>
