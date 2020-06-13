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
          title: '電気料金表',
          icon: 'fas fa-list',
        },
        {
          href: '/spreadsheet',
          title: 'スプレッドシート',
          icon: 'el-icon-s-grid',
        },
        {
           href: '/',
           title: '寮生管理',
           icon: 'fa fa-address-book'
        }
      ],
      widthstyle: "200px",
    }
  },
}
</script>

<style>
  @import url('../styles/sidebar.scss');

#home > div.v-sidebar-menu > div > div.vsm--list > div:nth-child(1) > a > i {
  background-color: rgb(73, 123, 211);
}

#home > div.v-sidebar-menu > div > div.vsm--list > div:nth-child(2) > a > i {
  background-color: rgb(54, 180, 128);
}

#home > div.v-sidebar-menu > div > div.vsm--list > div:nth-child(3) > a > i {
  background-color: rgb(219, 149, 43);
}
</style>