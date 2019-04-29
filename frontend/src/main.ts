import Vue from 'vue';
import VueMaterial from 'vue-material';
import Multiselect from'vue-multiselect';
import App from './App.vue';
import router from './router';
import store from './store';
import './registerServiceWorker';
import 'vue-material/dist/vue-material.min.css';
import 'vue-material/dist/theme/default.css';

Vue.use(VueMaterial);
Vue.component('multiselect', Multiselect)
Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
