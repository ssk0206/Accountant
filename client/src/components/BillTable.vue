<template>
  <div>
    <vue-good-table
      :columns="columns"
      :rows="rows"/>
  </div>
</template>

<script>
export default {
  name: 'my-component',
  data(){
    return {
      columns: [
        {
          label: 'Name',
          field: 'name',
        },
        {
          label: '部屋番号',
          field: 'roomid',
          type: 'number',
        },
        {
          label: '前のメーター値',
          field: 'pre_meter_val',
          type: 'number',
        },
        {
          label: '今のメーター値',
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