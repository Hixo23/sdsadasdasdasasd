<script>

    export default {
        data() {
            return {
                json: {
                    name: "",
                    url: ""
                },
                errors: {
                    name: "",
                    url: ""
                },
                valid: false
            }
        },
        methods: {
            createLink() {
                this.valid = true
                if(this.json.name.trim().length < 4) {
                    this.valid = false;
                    this.errors.name = "Name must be at least 4 character long"
                } else {
                    this.valid = true
                    this.errors.name = ""
                }

                if(this.json.url.trim().length < 4) {
                    this.valid = false;
                    this.errors.url = "URL must be at least 4 character long"
                } else {
                    this.valid = true
                    this.errors.url = ""
                }

                if(this.valid) {
                    console.log(this.json)
                }
            }
        }
    }
</script>

<template>
    <div class="w-full h-auto overflow-hiddenflex flex-col items-center space-y-12">
        <h2 class="text-white text-3xl text-center">Create Short Links!</h2>
        <form action="">
            <div class="bg-gray-600 w-[28rem] py-12 gap-4 mx-auto flex flex-col justify-center rounded-lg">
                <div class="mx-auto">
                    <input class="py-2 px-4 rounded-xl w-96" v-model="json.url" placeholder="Paste a link" type="text">
                    <p class="text-red-600">{{ errors.url }}</p>
                </div>
                <div class="mx-auto">
                    <input class="py-2 px-4 rounded-xl w-96" v-model="json.name" placeholder="Type a name" type="text">
                    <p class="text-red-600">{{ errors.name }}</p>
                </div>
                <button type="button" @click="createLink()" class="bg-blue-500 px-8 py-2 rounded-xl w-96 mx-auto hover:bg-transparent hover:border hover:border-blue-500 hover:text-white transition-all duration-150">Shorten</button>
            </div>
        </form>
    </div>
</template>
