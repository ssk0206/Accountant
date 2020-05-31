import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from 'element-ui'
import locale from 'element-ui/lib/locale/lang/ja'
import VueGoodTablePlugin from 'vue-good-table'
// import axios from 'axios'
import VueAxios from 'vue-axios'

import 'vue-good-table/dist/vue-good-table.css'

import './styles.scss'

const axiosBase = require('axios');
const axios = axiosBase.create({
  baseURL: 'http://localhost:8080', // バックエンドB のURL:port を指定する
  headers: {
    'Content-Type': 'application/json',
  },
  responseType: 'json'
});

Vue.use(ElementUI, { locale })
Vue.use(VueAxios, axios)
Vue.use(VueGoodTablePlugin)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
