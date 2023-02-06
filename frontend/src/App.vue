<script setup>
import Header from './components/Header.vue';
import Hero from './components/Hero.vue';
import Card from './components/Card.vue';

import { ref } from 'vue';
import { onMounted } from 'vue'
import axios from 'axios';



let data = ref(null)
onMounted(() => {
    axios.get('http://localhost:3000/links')
        .then(function (response) {
            data.value = response.data.links
        })
        .catch(function (error) {
            // handle error
            console.log(error);
        })
})

console.log(data.value)


</script>

<template>
    <div class="w-screen h-screen bg-gray-800 font-jetbrainsMono">
        <Header />
        <Hero />
        <div v-for="link in data" :id="link.id">
            <Card :link="link" />

        </div>
    </div>
</template>

<style scoped>

</style>
