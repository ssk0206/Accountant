<template>
  <el-table class="table-container" :data="rows" height="700" width="100%">
    <el-table-column prop="id" label="部屋番号" width="100"> </el-table-column>
    <el-table-column prop="name" label="名前" width="120"> </el-table-column>
    <el-table-column width="120" align="right">
      <template slot-scope="scope">
        <el-button
          size="mini"
          type="danger"
          @click="handleDelete(scope.$index, scope.row)"
          >削除</el-button
        >
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "my-component",
  data() {
    return {
      rows: [],
    };
  },
  created: function() {
    this.getAllData();
  },
  methods: {
    getAllData() {
      this.axios.get("/students").then((response) => {
        if (response.status != 200) {
          console.log("レスポンスエラー");
        } else {
          this.rows = response.data;
        }
      });
    },
    handleDelete(index, student) {
      this.axios.delete("/students", student.id).then((response) => {
        console.log(response.status);
      });
    },
  },
};
</script>

<style>
.table-container {
  margin: 20px 0 0 0;
}
</style>
