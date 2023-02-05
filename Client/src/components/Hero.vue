<script></script>

<template>
    <div class="w-full h-auto overflow-hiddenflex flex-col items-center space-y-12">
        <h2 class="text-white text-3xl text-center">Create Short Links!</h2>
        <form action="">
            <div class="bg-gray-600 w-[28rem] py-12 gap-4 mx-auto flex flex-col justify-center rounded-lg">

                <input class="py-2 px-4 rounded-xl w-96 mx-auto" placeholder="Paste a link" type="text">
                <input class="py-2 px-4 rounded-xl w-96 mx-auto" placeholder="Type a name" type="text">
                <button class="bg-blue-500 px-8 py-2 rounded-xl w-96 mx-auto hover:bg-transparent hover:border hover:border-blue-500 hover:text-white transition-all duration-150">Shorten</button>
            </div>
        </form>
    </div>
</template>