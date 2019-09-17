import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'

Vue.use(Router)

export const subtitles = [
  {
    title: 'Inventory',
    url: '/inventory',
    method: 'fetchInventories',
    headers: [
      { text: 'NO', value: 'Id' },
      { text: 'Vin#', value: 'VehicleIdentificationNumber' },
      { text: 'Model', value: 'ModelName' },
      { text: 'Make', value: 'Producer' },
      { text: 'Year', value: 'Year' },
      { text: 'MSRP', value: 'MSRP' },
      { text: 'Status', value: 'Status' },
      { text: 'Booked', value: 'Booked' },
      { text: 'Listed', value: 'Listed' },
    ]
  },
  {
    title: 'Commission',
    url: '/commision',
    method: 'fetchCommistions',
  },
  {
    title: 'Manage Market',
    url: '/market',
    method: 'fetchMarkets'
  },
  {
    title: 'Manage Customer',
    url: '/user',
    method: 'fetchUsers',
  },
  {
    title: 'Report Setting',
    url: '/setting',
    method: 'fetchSettings',
  },
  {
    title: 'Sign Out',
    url: '/auth',
    method: 'logOut'
  }
];

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
