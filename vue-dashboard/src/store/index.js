import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    count: 1
  },
  mutations: {
    increment (state, payload) {
      // 变更状态
      state.count++
    }
  },
  actions: {
    increment (context, payload) {
      context.commit('increment', payload)
    }
  },
  modules: {
  }
})
// store.commit('increment', 10)
// store.commit('increment', {
//   amount: 10
// })

// store.dispatch('increment') // action
