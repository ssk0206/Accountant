<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <Admin msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import Admin from '@/components/Admin.vue'

export default {
  name: 'Home',
  components: {
    Admin
  }
}
</script>
