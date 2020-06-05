import Vue from 'vue'
import VueRouter from 'vue-router'
import List from '../views/List.vue'
import Spreadsheet from '../views/Spreadsheet.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'List',
    component: List
  },
  {
    path: '/spreadsheet',
    name: 'Spreadsheet',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Spreadsheet
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
