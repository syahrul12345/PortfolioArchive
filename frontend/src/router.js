import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Dapp from './views/Dapp.vue'
import DappAll from './views/DappAll.vue'
import AllBlogs from './views/AllBlogs.vue'
import Resume from './views/Resume.vue'
import Contact from './views/Contact.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      meta: { title: 'Portfolio' },
      beforeEnter: (to,from,next) => {
        document.title = to.meta.title
        next()
      }
    },
    {
      path: '/dapp/:id',
      name: 'dapp',
      component: Dapp,
      meta: { title: 'Dapp' },
      beforeEnter: (to,from,next) => {
        document.title = to.meta.title
        next()
      }
    },
    {
      path:'/dapps/',
      name:'dappAll',
      component:DappAll,
      meta: { title: 'All Dapps' },
      beforeEnter: (to,from,next) => {
        document.title = to.meta.title
        next()
      }
    },
    {
      path:'/allBlogs/',
      name:'allBlogs',
      component:AllBlogs,
      meta: { title: 'All Blogs' },
      beforeEnter: (to,from,next) => {
        document.title = to.meta.title
        next()
      }
    },
    { 
      path:'/resume/',
      name:'resume',
      component:Resume,
      meta: { title: 'Resume' },
      beforeEnter: (to,from,next) => {
        document.title = to.meta.title
        next()
      }
    },
    { 
      path:'/contact/',
      name:'contact',
      component:Contact,
      meta: { title: 'Contact Me' },
      beforeEnter: (to,from,next) => {
        document.title = to.meta.title
        next()
      }
    }
  ],

})
