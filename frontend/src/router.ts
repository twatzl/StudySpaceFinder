import Vue from 'vue';
import Router from 'vue-router';
import Spaces from './views/Spaces.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'spaces',
      component: Spaces,
    },
  ],
});
