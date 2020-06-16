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
          {
            type: "numeric",
          },
          {
            type: "numeric",
          },
          {
            readOnly: true,
            type: "numeric",
          },
          {
            type: "numeric",
          },
          {},
        ],
        afterChange: function(changes_arr, source) {
          if (changes_arr) {
            changes_arr.forEach((changes) => {
              if (
                (source == "edit" ||
                  source == "CopyPaste.paste" ||
                  source == "Autofill.fill") &&
                (changes[1] == 2 || changes[1] == 3 || changes[1] == 5)
              ) {
                let row = changes[0];
                let col = changes[1];
                data[row][col] = Number(data[row][col]);
                data[row][4] =
                  (data[row][3] - data[row][2]) * 20 + Number(data[row][5]);
                this.render();
              }
            });
          }
        },
      };
    },
    formattedStudentData() {
      var data = [];
      this.studentData.forEach((element) => {
        var arr = [];
        arr.push(element["room_id"]);
        arr.push(element["name"]);
        arr.push(element["pre_meter_value"]);
        arr.push(element["new_meter_value"]);
        arr.push(element["bill"]);
        arr.push(element["additional_fee"]);
        arr.push(element["remark"]);
        arr.push(element["id"]);

        data.push(arr);
      });
      return data;
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
      var request = [];
      this.hotSettings.data.forEach((element) => {
        var map = {};
        map["room_id"] = element[0];
        map["name"] = element[1];
        map["pre_meter_val"] = element[2];
        map["new_meter_val"] = element[3];
        map["bill"] = element[4];
        map["additional_fee"] = element[5];
        map["remark"] = element[6];
        map["id"] = element[7];
        request.push(map);
      });

      this.axios.put("bills", JSON.stringify(request)).then((response) => {
        if (response.status != 204) {
          console.log(response);
        }
      });
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
