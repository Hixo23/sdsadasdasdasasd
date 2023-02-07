<script setup>
import {deleteLink} from '../utility/deleteLink'
const props = defineProps(['link'])

    const { ID, url, name } = props.link

</script>
<template>
    <div class="bg-[#0F0B0B] p-4 h-32 text-white flex flex-col justify-center items-center rounded-lg">
        <h3>Name: {{ name }}</h3>
        <p>URL: {{ url }}</p>
        <a>Link: http://localhost:3000/links/r/{{ name }}</a>
        <div class="text-white text-2xl flex justify-center gap-4">
            <button @click="test"><i class="fa-solid fa-pen-to-square"></i></button>
            <button @click="deleteLink(ID)"><i class="fa-solid fa-trash"></i></button>
        </div>
    </div>
</template>
