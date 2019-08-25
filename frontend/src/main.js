import Vue from 'vue'
import App from './App.vue'
import Dapp from './views/Dapp.vue'
import router from './router'
import './registerServiceWorker'
import vuetify from './plugins/vuetify';
import '@babel/polyfill'
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@fortawesome/fontawesome-free/css/all.css'
import VueParticles from 'vue-particles'

Vue.use(VueParticles)
Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
