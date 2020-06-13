<template>
  <el-container>
    <el-aside width="400px"><students-table /></el-aside>
    <el-main><students-form /></el-main>
  </el-container>
</template>

<script>
// @ is an alias to /src
import StudentsTable from "@/components/StudentsTable.vue";
import StudentsForm from "@/components/StudentsForm.vue";

export default {
  components: {
    StudentsTable,
    StudentsForm,
  },
  props: {
    columns: Array,
    rows: Array,
  },
};
</script>
