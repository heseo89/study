import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        todos: [],
    },
    mutations: {
        add_todo(state, name){
          if (name != null) {
            state.todos.push({
                name: name
            })
          }
        },
        del_todo(state, index){
          state.todos.splice(index, 1)
        },
    },
    actions: {
        addTodo({commit}){
            commit('ADD_TODO')
        },
        removeTodo({commit}, todo){
            commit('REMOVE_TODO', todo)
        }
    }
})
