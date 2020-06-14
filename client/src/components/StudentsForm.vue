<template>
  <el-form :inline="true" :model="student" class="">
    <el-form-item label="部屋番号">
      <el-input v-model="student.id" placeholder="000"></el-input>
    </el-form-item>
    <el-form-item label="名前">
      <el-input v-model="student.name" placeholder="山田 太郎"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">新規登録</el-button>
    </el-form-item>
  </el-form>
</template>
<script>
export default {
  data() {
    return {
      student: {
        id: "",
        name: "",
      },
    };
  },
  methods: {
    onSubmit() {
      console.log(this.student);
      this.axios.post("/students", this.student).then((response) => {
        if (response.status == 201) {
          window.location.reload(); // ページごとリロード
        }
      });
    },
  },
};
</script>
