<script setup lang="ts">
</script>

<template>
  <nav class="bg-teal-400 fixed w-full z-20 top-0 start-0">
    <div class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
      <h1 class="font-bold text-white text-3xl">Dataset Visualization App</h1>
      <ul class="flex flex-row space-x-10">
        <li>
            <RouterLink to="/logout" class="text-white text-lg">Logout</RouterLink>
        </li>
      </ul>
    </div>
  </nav>
</template>