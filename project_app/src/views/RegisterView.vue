<script setup lang="ts">
import { reactive } from 'vue';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

  const authStore = useAuthStore()
  const router = useRouter()
  const registerInputs = reactive({
    username: "",
    password: ""
  })

  async function registerUser(){
    // validar que los campos no estén vacíos
    if (!registerInputs.username || !registerInputs.password) {
      alert("Username and password are required")
      return
    }
    // realizar la solicitud al backend
    const response = await fetch('http://localhost:3333/auth/register',{
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: registerInputs.username,
        password: registerInputs.password
      })
    })
    // manejo de la respuesta
    if (response.ok){
      const data = await response.json()
      authStore.setSession(data.session)
      router.push("/data")
      
    } else {
      const error = await response.json()
      console.error(error)
    }

  }
</script>

<template>
    <div class="bg-purple-100 flex flex-col space-y-4 h-dvh py-16">
    <p class="py-16 text-center">This is a web application for a dataset of User Income.</p>
    <form @submit.prevent="registerUser" class="space-y-6 text-center">
      <h3 class="font-bold text-xl text-pink-600">Register Page</h3>
      <label class="block">Username</label>
      <input v-model="registerInputs.username" type="text" class="w-1/6 rounded-lg border-solid p-2">
      <label class="block">Password</label>
      <input v-model="registerInputs.password" type="password" class="w-1/6 rounded-lg border-solid p-2">
      <br>
      <button class="bg-pink-600 w-1/6 text-white rounded-lg p-2">Register</button>
    </form>
  </div>
</template>
