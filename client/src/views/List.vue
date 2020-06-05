<template>
  <div class="list">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <bill-table />
  </div>
</template>

<script>
// @ is an alias to /src
import BillTable from '@/components/BillTable.vue'

export default {
  name: 'List',
  components: {
    BillTable
  },
  props: {
    columns: Array,
    rows: Array,
  }
}
</script>
