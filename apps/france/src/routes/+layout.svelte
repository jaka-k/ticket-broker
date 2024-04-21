<script lang="ts">
  import '../app.css';
  import { ThemeWrapper, Button } from '@ticket-broker/core';

  let theme = 'light'; // Default theme

function handleThemeChange(event: Event) {
  const select = event.target as HTMLSelectElement;
  select.value ? (theme = select.value) : null;
}
</script>

<select on:change={handleThemeChange}>
  <option value="light">Light</option>
  <option value="dark">Dark</option>
  <option value="cupcake">Cupcake</option>
  <option value="bumblebee">Bumblebee</option>
</select>
<Button>Test Outside</Button>
<ThemeWrapper {theme}>

  <slot />
</ThemeWrapper>
