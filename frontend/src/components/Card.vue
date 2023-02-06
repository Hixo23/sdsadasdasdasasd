<script setup>
const props = defineProps(['link'])

console.log(props.link)
</script>
<template>
    <div>
        {{ props.link }}
    </div>
</template>
