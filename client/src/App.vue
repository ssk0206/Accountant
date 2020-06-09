<template>
  <div id="home">
    <sidebar-menu />
    <div id="body-container">
      <router-view :studentData="rows" />
    </div>
  </div>
</template>

<script>
import SidebarMenu  from '@/components/Sidebar.vue'
export default {
  components: {
    SidebarMenu
  },
  data() {
    return {
      studentData: [
        ['201', "桜木 花道", 11, 12, 13, 1],
        ['202', "流川", 11, 14, 13],
        ['211', "赤木", 15, 12, 13],
        ['201', "桜木", 11, 12, 13, 1],
        ['202', "流川", 11, 14, 13],
        ['211', "赤木", 15, 12, 13],
      ],
      rows: [],
    }
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
}
</script>

<style>
#home {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background-color: #f4fbff;
}

#body-container {
  padding: 20px 100px;
}
</style>
