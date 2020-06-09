<template>
  <div class="spread-container">
    <hot-table :settings="hotSettings"></hot-table>
  </div>
</template>

<script>
import { HotTable } from '@handsontable/vue';

export default {
  data: function() {
    return {
      hotSettings: {
        data: this.studentData,
        colWidths: [80, 100,120,120,120,200],
        colHeaders: ['部屋番号', '名前', 'メーター値(旧)', 'メーター値(新)', '料金', '備考'],
        rowHeaders: [],
      }
    };
  },
  components: {
    HotTable
  },
  props: {
    studentData: {
      type: Array
    }
  }
}
</script>

<style scoped>
@import '~handsontable/dist/handsontable.full.css';
</style>