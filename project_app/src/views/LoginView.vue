<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { reactive } from 'vue';
import { useRouter } from 'vue-router';

  const authStore = useAuthStore()
  const router = useRouter()

  const loginInputs = reactive({
    username: "",
    password: ""
  })

  async function signIn(){
    const response = await fetch('http://localhost:3333/auth/login',{
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(loginInputs)
    })
    const data = await response.json()

    if (data.session) {
      authStore.setSession(data.session)
      router.push("/data")
    } else {
      console.error("No token received in the response.")
    }
    
  }
</script>

<template>
    <div class="bg-purple-100 flex flex-col space-y-4 h-dvh py-16">
    <p class="py-16 text-center">This is a web application for a dataset of User Income. </p>
    <form @submit.prevent="signIn" class="space-y-6 text-center">
      <h3 class="font-bold text-xl text-pink-600">Login Page</h3>
      <label class="block">Username</label>
      <input v-model="loginInputs.username" type="text" class="w-1/6 rounded-lg border-solid p-2">
      <label class="block">Password</label>
      <input v-model="loginInputs.password" type="password" class="w-1/6 rounded-lg border-solid p-2">
      <br>
      <button class="bg-pink-600 w-1/6 text-white rounded-lg p-2">Login</button>
    </form>
      <br>
      <p class="text-center">Don't have an account?   <RouterLink to="/register" class="font-bold text-teal-400">Register</RouterLink></p>
  </div>
</template>