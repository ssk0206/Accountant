<template>
  <div id="home">
    <sidebar-menu />
    <!-- <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div> -->
    <div id="body-container">
      <router-view/>
    </div>
  </div>
</template>

<script>
import SidebarMenu  from '@/components/Sidebar.vue'
export default {
  components: {
    SidebarMenu
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

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

#body-container {
  padding: 20px 100px;
}
</style>
