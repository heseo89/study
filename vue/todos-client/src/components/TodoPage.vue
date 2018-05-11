<template>
<div class="container">
  <h2>Todo List</h2>
  <div class="input-group" style="margin-bottom:10px;">
  <input type="text" class="form-control" placeholder="할일을 입력하세요" v-model="name" v-on:keyup.enter="createTodo(name)">
  <span class="input-group-btn">
    <button class="btn btn-default" type="button" v-on:click="createTodo(name)">추가</button>
  </span>
  </div>
  <ul class="list-group">
  <li class="list-group-item" v-for="(todo, index) in todo_list" :key="index">
  {{todo.name}}
    <div class="btn-group pull-right" style="font-size: 12px; line-height: 1;">
      <button type="button" class="btn-group btn-default" v-on:click="deleteTodo(index)">
      삭제
      </button>
    </div>
  </li>
  </ul>
</div>
</template>

<script>
export default {
  name: 'TodoPage',
  data () {
    return {
      name: null,
      todo_list: []
    }
  },
  methods: {
    deleteTodo (index) {
      this.todo_list.splice(index, 1)
    },
    createTodo (name) {
      if (name != null) {
        this.todo_list.push({name: name})
        this.name = null
      }
    }
  }
}
</script>
