<template>
  <div class="spread">
    <Spreadsheet :studentData="formattedStudentData" />
  </div>
</template>

<script>
// @ is an alias to /src
import Spreadsheet from '@/components/Spread.vue'

export default {
  name: 'Home',
  components: {
    Spreadsheet
  },
  props: {
    studentData: {
      type: Array,
    }
  },
  computed: {
    formattedStudentData () {
      let arr = this.studentData.map(element => {
        return Object.values(element).filter((val, i) => {
          return (i != 0 && i != 6 && i != 7);
        })
      });
      return arr
    }
  },
}
</script>

<style>
.spread {
  height: 100%;
  margin: 0 auto;
  width: 760px;
}
</style>
