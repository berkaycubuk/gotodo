<template>
    <div class="container">
        <h1>GoTodo Web Client</h1>
        <form @submit.prevent="addNewTodo">
            <input v-model="newTodo" placeholder="New Todo" name="newTodo"/>
            <button>Add</button>
        </form>
        <h2>List</h2>
        <ul>
            <li v-for="todo in todos" :key="todo.id">
                <span :class="{ done: todo.status }">
                {{ todo.title }}
                </span>
            </li>
        </ul>
    </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap');

* {
    font-family: 'Roboto', sans-serif;
}

body {
    background-color: #222;
}

.container {
    max-width: 800px;
    width: 100%;
    margin: 0 auto;
}

h1, h2 {
    color: #eaeaea;
}

form {
    display: flex;
    margin: 20px 0;
}

input {
    flex: 1;
    font-size: 1.6rem;
    padding: 10px 20px;
    color: #eaeaea;
    background-color: #555;
    border: none;
    outline: none;
}

button {
    font-size: 1.2rem;
    cursor: pointer;
    padding: 10px 20px;
    background-color: #eee;
    border: none;
    outline: none;
}

ul {
    list-style: none;
    padding: 0;
}

li {
    margin: 10px 0;
    padding: 12px;
    color: #eaeaea;
    background-color: #333;
}

li > span.done {
    text-decoration: line-through;
}
</style>

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
