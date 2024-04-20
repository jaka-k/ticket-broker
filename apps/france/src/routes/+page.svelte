<script lang="ts">
  import { ThemeWrapper } from '@ticket-broker/ui';

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
<ThemeWrapper {theme}>
  <h1 class="text-red-800">Content with dynamically changing theme</h1>
  <p>Depending on user selection, the theme changes.</p>
</ThemeWrapper>
