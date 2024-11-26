import { defineStore } from "pinia";
import { ref } from "vue";
import type { UserFilters } from '../types/index';

export const useFilterStore = defineStore('filters', () => {

    const userFilters = ref<UserFilters>({
        education: '',
        marital_status: '',
        occupation: '',
        income: '',
        order_by: '',
        order_direction: '',
        min_age: '',
        max_age: '',
    })

    // funcion para recibir los filtros del usuario del backend
    async function loadUserFilters(){
        try {
            // obtener el token de la session
            const sessionToken = sessionStorage.getItem('token')
            const url = "http://localhost:3333/data/filters";

            const response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${sessionToken}`,
                }
            })
            if (!response.ok){
                throw new Error(`HTTP error! Status: ${response.status}`)
            }
            const result = await response.json()
            if (result && result.filters) {
                userFilters.value = { ...userFilters.value, ...(result.filters as UserFilters) };
            } else {
                console.error('Unexpected response structure:', result);
            }

        } catch (error) {
            console.error('Failed to get User filters:', error)
        }
    }

    async function updateUserFilters(){
        try{
            // obtener el token de la session
            const sessionToken = sessionStorage.getItem('token')
            const url = "http://localhost:3333/data/filters";
            const filtersToSend = userFilters.value; 

            const response = await fetch(url, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${sessionToken}`,
                },
                body: JSON.stringify(filtersToSend)
            })
            if (!response.ok){
                throw new Error(`HTTP error! Status: ${response.status}`)
            }
            const result = await response.json()
            userFilters.value = (result.filters as UserFilters)
        } catch (error) {
            console.error('Failed to get User filters:', error)
        }
    }

   

    return { userFilters, loadUserFilters, updateUserFilters};
})