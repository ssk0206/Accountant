<template>
  <div class="spread">
    <div>
      <div>{{ yearMonth }}</div>
      <Spreadsheet />
    </div>
    <div class="buttons">
      <el-row>
        <el-button type="primary" v-on:click="update">寮生情報更新</el-button>
      </el-row>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import Spreadsheet from "@/components/Spread.vue";

export default {
  name: "Home",
  components: {
    Spreadsheet,
  },
  computed: {
    yearMonth() {
      let today = new Date();
      let year = today.getFullYear();
      let month = today.getMonth();
      if (1 <= month && month <= 3) {
        month = "04-07";
      } else if (4 <= month && month <= 6) {
        month = "07-09";
      } else if (7 <= month && month <= 9) {
        month = "10-12";
      } else {
        month = "01-03";
      }
      return year + month;
    },
    period() {
      return this.$route.query.period
        ? this.$route.query.period
        : this.yearMonth;
    },
  },
  methods: {
    update() {
      this.axios.post("bills", { period: this.period }).then((response) => {
        if (response.status == 201) {
          window.location.reload();
        }
      });
    },
  },
};
</script>

<style>
.spread {
  height: 100%;
  margin: 0 40px 0 110px;
  width: 760px;
  display: flex;
}
.buttons {
  margin: 0 60px;
}
</style>
