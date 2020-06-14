<template>
  <el-container>
    <el-aside width="340px" style="margin:0 0 0 110px">
      <students-table />
    </el-aside>
    <el-main>
      <students-form />
    </el-main>
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
