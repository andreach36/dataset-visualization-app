<script setup lang="ts">
import { onBeforeMount } from 'vue';
import { useAuthStore } from './stores/auth';
import NavBar from './components/NavBar.vue';

const authStore = useAuthStore()

onBeforeMount(authStore.init)

</script>


<template>
  <NavBar></NavBar>
  <RouterView></RouterView>
</template>