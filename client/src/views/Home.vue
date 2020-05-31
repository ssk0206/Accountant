<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <bill-table msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import BillTable from '@/components/BillTable.vue'

export default {
  name: 'Home',
  components: {
    BillTable
  }
}
</script>
