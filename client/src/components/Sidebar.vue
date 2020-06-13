<template>
  <sidebar-menu :menu="menu" :width="widthstyle" :collapsed="true" ></sidebar-menu>
</template>

<script>
import { SidebarMenu } from 'vue-sidebar-menu'
export default {
  name: 'Sidebar',
  components: {
    SidebarMenu
  },
  data() {
    return {
      menu: [
        {
           href: '/',
           title: '寮生管理',
           icon: 'fa fa-address-book'
        },
        {
          href: '/spreadsheet',
          title: 'スプレッドシート',
          icon: 'el-icon-s-grid',
        },
      ],
      widthstyle: "200px",
    }
  },
}
</script>

<style>
  @import url('../styles/sidebar.scss');

#home > div.v-sidebar-menu > div > div.vsm--list > div:nth-child(1) > a > i {
  background-color: rgb(219, 149, 43);
}

#home > div.v-sidebar-menu > div > div.vsm--list > div:nth-child(2) > a > i {
  background-color: rgb(54, 180, 128);
}
</style>