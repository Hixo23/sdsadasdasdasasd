<script setup>
import { fetchData } from '../utility/fetchData'
import { ref } from 'vue';
import axios from 'axios';
let json = ref({
    name: "",
    url: ""
})
let errors = ref({
    name: "",
    url: ""
})

let valid
const createLink = () => {
    valid = true
    if (json.value.name.trim().length < 1) {
        errors.value.name = "Name must be at least 1 character long"
        valid = false;
    } else {
        valid = true
        errors.value.name = ""
    }

    if (json.value.url.trim().length < 1) {
        errors.value.url = "URL must be at least 1 character long"
        valid = false;
    } else {
        valid = true
        errors.value.url = ""
    }

    if (valid) {
        axios.post('http://localhost:3000/links', {
            name: json.value.name,
            url: json.value.url
        })
        setTimeout(() => {
            fetchData()
        }, 1000)

    } else {
        console.log(errors)
    }
}
</script>


<template>
    <div v-motion :initial="{
        opacity: 0,
        y: 50,
    }" :enter="{
    opacity: 1,
    y: 0,
}" class="w-full h-auto overflow-hiddenflex flex-col items-center space-y-12">
        <h2 class="text-white text-3xl text-center ">Create Short Links!</h2>
        <form action="">
            <div class="bg-[#0B0C10] w-[28rem] py-12 gap-4 mx-auto flex flex-col justify-center rounded-lg">
                <div class="mx-auto">
                    <input class="py-2 px-4 rounded-xl w-96 text-black" v-model="json.name" placeholder="Type a name"
                        type="text">
                    <p class="text-red-600 mt-2">{{ errors.name }}</p>
                </div>
                <div class="mx-auto">
                    <input class="py-2 px-4 rounded-xl w-96 text-black" v-model="json.url" placeholder="Paste a link"
                        type="text">
                    <p class="text-red-600 mt-2">{{ errors.url }}</p>
                </div>
                <button type="button" @click="createLink()"
                    class="bg-slate-800	text-[#C5C6C7] px-8 py-2 rounded-xl w-96 mx-auto hover:bg-transparent hover:border hover:border-blue-500 hover:text-white transition-all duration-150">Shorten</button>
            </div>
        </form>
    </div>
</template>
