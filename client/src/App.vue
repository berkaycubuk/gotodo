<template>
    <h1>GoTodo Web Client</h1>
    <form @submit.prevent="addNewTodo">
        <label>New Todo</label>
        <input v-model="newTodo" name="newTodo"/>
        <button>Add New Todo</button>
    </form>
    <h2>List</h2>
    <ul>
        <li v-for="todo in todos" :key="todo.id">{{ todo.title }}</li>
    </ul>
</template>

<script>
import { ref } from 'vue';
import axios from 'axios';

export default {
    data () {
        return {
            todos: []
        };
    },
    created () {
        axios.get('http://localhost:8000/todos')
            .then(res => {
                this.todos = res.data;
            });
    },
    setup () {
        const newTodo = ref('');
        
        function addNewTodo() {
        }

        return {
            newTodo,
            addNewTodo,
        };
    }
}
</script>
