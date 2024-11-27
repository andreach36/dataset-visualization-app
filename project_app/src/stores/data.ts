import { defineStore } from "pinia";
import { type DataPagination, type DataRecord, type UserFilters } from '@/types';
import { ref } from "vue";

export const useDataStore = defineStore('data', () => {
    
    // se crea lista de datos de la tabla
    const dataRecords = ref<DataRecord[]>([])

    const dataPagination = ref<DataPagination>({
        page: 1,
        page_size: 10,
        total_pages: 0,
        total_records: 0,
      })
    
    // función para cargar datos de la tabla
    // se recibe la página y los filtros del usuario
    async function loadData (page: number, filters: UserFilters) {
        try {
        // obtener el token de la session
        const sessionToken = sessionStorage.getItem('token')
        const baseUrl = "http://localhost:3333/data";
    
        // Verificar si hay filtros seleccionados
        const hasFilters =
            filters.min_age || filters.max_age || filters.education || filters.marital_status || filters.occupation || filters.income

        // Construir la url para obtener los datos
        let url = baseUrl
    
        if (hasFilters) {
            url = "http://localhost:3333/data/filter"; // Cambiar la base de la URL si hay filtros
        
        // Agregar el segmento dinámico para la edad si está definido
        if (filters.min_age || filters.max_age) {
            url += `?min_age=${filters.min_age || 0}&max_age=${filters.max_age || 100}`;
        }
        }
    
        // construir parámetros dinámicos
        const params = new URLSearchParams({
            page: String(page),
            page_size: String(dataPagination.value.page_size),
        })
    
        // agregar filtros opcionales
        if (filters.education && filters.education !== "All") {
            params.append('education', filters.education)
        }
        if (filters.marital_status && filters.marital_status !== "All") {
            params.append('marital_status', filters.marital_status)
        }
        if (filters.occupation && filters.occupation !== "All") {
            params.append('occupation', filters.occupation)
        }
        if (filters.income && filters.income !== "All") {
            params.append('income', filters.income)
        }
        // Solo agregar order_by y order_direction si tienen un valor seleccionado
        if (filters.order_by && filters.order_by !== "All") {
            params.append('order_by', filters.order_by.toLowerCase());
        }
        if (filters.order_direction && filters.order_direction !== "All") {
            params.append('order_direction', filters.order_direction.toUpperCase());
        }
    
        // Concatenar los parámetros adicionales a la URL
        if (hasFilters) {
            url += `&${params.toString()}`;  // Concatenamos a la URL de los filtros
        } else {
            url += `?${params.toString()}`;  // En caso de no tener filtros, empezamos con '?'
        }
       
        // realizar la solicitud
        const response = await fetch(url,{
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${sessionToken}`,
            },
        })
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`)
        }
        const result = await response.json()
        dataRecords.value = (result.data as DataRecord[])
        dataPagination.value = (result.meta as DataPagination)
    } catch (error) {
        console.error('Failed to load data:', error)
        dataRecords.value = []
        }
    }

    return {dataRecords, dataPagination, loadData}

})