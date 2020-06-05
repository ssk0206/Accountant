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
          data: [
            ['201', "桜木 花道", 11, 12, 13, 1],
            ['202', "流川", 11, 14, 13],
            ['211', "赤木", 15, 12, 13],
            ['201', "桜木", 11, 12, 13, 1],
            ['202', "流川", 11, 14, 13],
            ['211', "赤木", 15, 12, 13],
          ],
          colWidths: [80, 100,120,120,120,200],
          colHeaders: ['部屋番号', '名前', 'メーター値(旧)', 'メーター値(新)', '料金', '備考'],
          rowHeaders: [],
        }
      };
    },
    components: {
      HotTable
    },
  }
</script>

<style scoped>
@import '~handsontable/dist/handsontable.full.css';

.spread-container {
  margin: 0 0 0 140px;
}
</style>