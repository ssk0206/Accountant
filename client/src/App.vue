<template>
  <div id="home">
    <sidebar-menu />
    <div id="body-container">
      <router-view />
    </div>
  </div>
</template>

<script>
import SidebarMenu from "@/components/Sidebar.vue";
export default {
  components: {
    SidebarMenu,
  },
};
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
