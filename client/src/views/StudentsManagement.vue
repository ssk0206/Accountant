<template>
  <div class="students-table">
    <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
    <students-table />
  </div>
</template>

<script>
// @ is an alias to /src
import StudentsTable from '@/components/StudentsTable.vue'

export default {
  components: {
    StudentsTable
  },
  props: {
    columns: Array,
    rows: Array,
  }
}
</script>
