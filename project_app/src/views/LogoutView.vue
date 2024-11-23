<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
const authStore = useAuthStore();

function handleLogout() {
    authStore.clearSession();
}
</script>

<template>
    <div class="bg-purple-100 flex flex-col space-y-4 h-dvh py-16">
    <p class="py-16 text-center">Thanks for using My DataSet App!</p>
    <button @click="handleLogout" class="bg-pink-600 w-1/6 text-white rounded-lg p-2 mx-auto">
      Logout
    </button>
  </div>
</template>