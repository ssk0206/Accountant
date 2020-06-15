<template>
  <div class="spread-container">
    <hot-table :settings="hotSettings"></hot-table>
    <el-button type="success" v-on:click="submit">
      保存
    </el-button>
  </div>
</template>

<script>
import { HotTable } from "@handsontable/vue";

export default {
  components: {
    HotTable,
  },
  computed: {
    hotSettings() {
      var data = this.formattedStudentData;
      return {
        data: data,
        colWidths: [80, 100, 120, 120, 120, 120, 200],
        colHeaders: [
          "部屋番号",
          "名前",
          "メーター値(旧)",
          "メーター値(新)",
          "料金",
          "追加料金",
          "備考",
        ],
        rowHeaders: [],
        columns: [
          {
            readOnly: true,
          },
          {
            readOnly: true,
          },
          {},
          {},
          {
            readOnly: true,
          },
          {},
          {},
        ],
        afterChange: function(changes, source) {
          if (
            (source == "edit" ||
              source == "CopyPaste.paste" ||
              source == "Autofill.fill") &&
            (changes[0][1] == 2 || changes[0][1] == 3)
          ) {
            var row = changes[0];
            data[row[0]][row[1]] = Number(data[row[0]][row[1]]);
            data[row[0]][4] = (data[row[0]][3] - data[row[0]][2]) * 20;
            this.render();
          }
        },
      };
    },
    formattedStudentData() {
      this.studentData.forEach((element) =>
        Object.keys(element).forEach(function(key) {
          if (key == "bill") {
            element[key] =
              (element["new-meter-value"] - element["pre-meter-value"]) * 20;
          }
        })
      );

      let arr = this.studentData.map((element) => {
        return Object.values(element).filter((val, i) => {
          return i != 7;
        });
      });
      return arr;
    },
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
  data() {
    return {
      studentData: [],
    };
  },
  methods: {
    getAllData() {
      this.axios
        .get("bills", { params: { period: this.period } })
        .then((response) => {
          if (response.status != 200) {
            console.log("レスポンスエラー");
          } else {
            this.studentData = response.data;
          }
        });
    },
    update() {
      this.axios.post("bills", { period: this.period }).then((response) => {
        if (response.status == 201) {
          window.location.reload();
        }
      });
    },
    submit() {
      console.log(this.hotSettings.data);
    },
  },
  created: function() {
    this.getAllData();
  },
};
</script>

<style scoped>
@import "~handsontable/dist/handsontable.full.css";
</style>
