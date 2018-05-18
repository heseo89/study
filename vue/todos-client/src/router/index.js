import Vue from 'vue'
import Router from 'vue-router'
import TodoPage from '@/components/TodoPage'

Vue.use(Router)

Vue.component('todo-input', {
    template: '<div class="input-group" style="margin-bottom:10px;"><input type="text" class="form-control" placeholder="할일을 입력하세요" v-model="name" v-on:keyup.enter="createTodo(name)"><span class="input-group-btn"><button class="btn btn-default" type="button" v-on:click="createTodo(name)">{{ title }}</button></span></div>',
    data: function () {
      return {
        title: '추가',
        name: null,
      }
    },
    methods: {
        createTodo (name) {
          this.$store.commit('add_todo', name)
          this.name = null
        }
    }
})
Vue.component('todo-list', {
    template: '<ul class="list-group"><li class="list-group-item" v-for="(todo, index) in this.$store.state.todos" :key="index">{{todo.name}}<div class="btn-group pull-right" style="font-size: 12px; line-height: 1;"><button type="button" class="btn-group btn-default" v-on:click="deleteTodo(index)">{{ title }}</button></div></li></ul>',
    data: function () {
        return {
          title: '삭제'
        }
    },
    methods: {
        deleteTodo (index) {
          this.$store.commit('del_todo', index)
        }
    }
})

export default new Router({
  routes: [
    {
      path: '/',
      name: 'TodoPage',
      component: TodoPage
    }
  ]
})
