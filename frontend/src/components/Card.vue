<script setup>
import axios from 'axios';
const props = defineProps(['link'])

    const { ID, url, name } = props.link

    const deleteLink = () => {
        axios.delete(`http://localhost:3000/links/${ID}`).then(res => {
            console.log(res)
        })
    }
</script>
<template>
    <div class="bg-[#0F0B0B] w-[28rem] h-32 text-white flex flex-col justify-center items-center rounded-lg">
        <h3>Name: {{ name }}</h3>
        <p>URL: {{ url }}</p>
        <a>Link: http://localhost:3000/links/r/{{ name }}</a>
        <div class="text-white text-2xl flex justify-center gap-4">
            <button><i class="fa-solid fa-pen-to-square"></i></button>
            <button @click="deleteLink()"><i class="fa-solid fa-trash"></i></button>
        </div>
    </div>
</template>
