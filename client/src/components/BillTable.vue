<template>
  <div class="table-container">
    <vue-good-table
      :columns="columns"
      :rows="rows"
      styleClass="vgt-table condensed bordered"
      :select-options="{ enabled: true }"
      :search-options="{ enabled: true }"
      max-height="600px"
      :fixed-header="true"
      :line-numbers="true"
      >
      <template slot="table-row" slot-scope="props">
        <span v-if="props.column.field == 'bill'">
          <span style="font-weight: bold;">{{props.row.bill}}円</span>
        </span>
        <span v-else>
          {{props.formattedRow[props.column.field]}}
        </span>
      </template>
    </vue-good-table>
  </div>
</template>

<script>
export default {
  name: 'my-component',
  data(){
    return {
      columns: [
        {
          label: '名前',
          field: 'name',
        },
        {
          label: '部屋番号',
          field: 'roomid',
          type: 'number',
        },
        {
          label: 'メーター値(前)',
          field: 'pre_meter_val',
          type: 'number',
        },
        {
          label: 'メーター値(今)',
          field: 'new_meter_val',
          type: 'number',
        },
        {
          label: '電気料金',
          field: 'bill',
          type: 'number',
        },
      ],
      rows: [],
    };
  },
  created: function() {
    this.getAllData()
  },
  methods: {
    getAllData() {
      this.axios.get("/students")
      .then(response => {
        if (response.status != 200) {
          console.log("レスポンスエラー")
        } else {
          this.rows = response.data
        }
      })
    }
  },
};
</script>

<style>
.table-container {
  width: 800px;
  margin: auto;
}
</style>