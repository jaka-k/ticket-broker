<script lang="ts">
  type OrderCall = {
    id: number
    countryCode: 'fr' | 'de' | 'gb' | 'es';
    timestamp: Date;
    orderAmmount: number;
    status: "pending" | "canceled" | "finished"
  };
  export let data: OrderCall[] = []; // Expects an array of objects

  // Optional: Define headers if your data objects have consistent keys
  // This could also be passed as a prop if different instances of the component need different headers
  let headers = Object.keys(data[0] || {});
</script>

<div class="overflow-x-auto">
  <table class="table table-xs">
    <!-- Head -->
    <thead>
      <tr>
        {#each headers as header}
          <th>{header}</th>
        {/each}
      </tr>
    </thead>
    <!-- Body -->
    <tbody>
      {#each data as row (row.id)}
        <tr>
          {#each headers as header}
            <td>{row["timestamp"]}</td>
          {/each}
        </tr>
      {/each}
    </tbody>
  </table>
</div>
