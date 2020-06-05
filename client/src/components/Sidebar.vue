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
</style>