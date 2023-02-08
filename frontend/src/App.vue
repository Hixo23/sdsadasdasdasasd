<script setup>
import Header from './components/Header.vue';
import Hero from './components/Hero.vue';
import Card from './components/Card.vue';
import { data } from './utility/fetchData'




</script>

<template>
    <div class="w-screen min-h-screen bg-[#0B0C10] font-jetbrainsMono">
        <Header />
        <Hero />
        <div  class="flex flex-wrap items-center justify-center gap-2 mt-4">
            <Card v-for="link in data" :id="link.id" :link="link" />
        </div>
    </div>
</template>

<style scoped>

</style>
