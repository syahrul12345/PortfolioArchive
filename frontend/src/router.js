import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Dapp from './views/Dapp.vue'
import DappAll from './views/DappAll.vue'
import AllBlogs from './views/AllBlogs.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/dapp/:id',
      name: 'dapp',
      component: Dapp
    },
    {
      path:'/dapps/',
      name:'dappAll',
      component:DappAll
    },
    {
      path:'/allBlogs/',
      name:'allBlogs',
      component:AllBlogs
    }
  ]
})
