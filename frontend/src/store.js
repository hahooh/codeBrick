import Vue from 'vue'
import Vuex from 'vuex'
import listStore from './store/list'
import titleStore from './store/title'

Vue.use(Vuex)


export default new Vuex.Store({
  modules: {
    list: {
      ...listStore,
      namespaced: true
    },
    title: {
      ...titleStore,
      namespaced: true
    }
  }
})
