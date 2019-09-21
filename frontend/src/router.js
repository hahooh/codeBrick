import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import { subtitles } from './constants/routes'

Vue.use(Router)

const routes = subtitles.map(subtitle => {
  return {
    path: subtitle.url,
    name: subtitle.title,
    component: () => import('./views/List.vue')
  }
});
routes.push({
  path: '/',
  name: 'home',
  component: Home
});
routes.push({
  path: '*',
  name: 'PageNotFount',
  component: () => import('./views/PageNotFound.vue')
})

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})
