<template>
  <div class="table-container">
    <vue-good-table
      :columns="columns"
      :rows="rows"
      styleClass="vgt-table condensed bordered"
      :search-options="{ enabled: true }"
      max-height="600px"
      :line-numbers="true"
      >
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
          label: '部屋番号',
          field: 'id',
          type: 'number',
          width: '80px'
        },
        {
          label: '名前',
          field: 'name',
          width: '120px'
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
  width: 320px;
  margin: 0 100px;
}
</style>