<template>
  <div class="spread">
    <Spreadsheet />
  </div>
</template>

<script>
// @ is an alias to /src
import Spreadsheet from '@/components/Spread.vue'

export default {
  name: 'Home',
  components: {
    Spreadsheet
  }
}
</script>

<style>
.spread {
  height: 100%;
}
</style>
