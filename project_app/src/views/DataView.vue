<script setup lang="ts">
import { type DataPagination, type DataRecord } from '@/types';
import { HeadersTable, filterOptionsEducation, filterOptionsOccupation, filterOptionsMaritalStatus, filterOptionsIncome, orderOptions } from '@/data_table/labels';
import { onMounted, ref, watch } from 'vue';

// se crea lista de datos de la tabla
const dataRecords = ref<DataRecord[]>([])

const dataPagination = ref<DataPagination>({
  page: 1,
  page_size: 10,
  total_pages: 0,
  total_records: 0,
})

// rango de edad y valores seleccionados
const filters = ref({
  education : '',
  marital_status : '',
  occupation : '',
  income: '',
  order_by: '',
  order_direction: '',
  min_age: 0,
  max_age: 100,
  export: false,

})


// función para cargar datos de la tabla
async function loadData (page: number) {
  try {

    let baseUrl = "http://localhost:3333/data";

     // Verificar si hay filtros seleccionados
     const hasFilters =
      filters.value.min_age || filters.value.max_age || filters.value.education ||
      filters.value.marital_status || filters.value.occupation || filters.value.income;

    let url = baseUrl

    if (hasFilters) {
      url = "http://localhost:3333/data/filter"; // Cambiar la base de la URL si hay filtros
    
     // Agregar el segmento dinámico para la edad si está definido
     if (filters.value.min_age || filters.value.max_age) {
      url += `?min_age=${filters.value.min_age || 0}&max_age=${filters.value.max_age || 100}`;
    }
  }

    // construir parámetros dinámicos
    const params = new URLSearchParams({
      page: String(page),
      page_size: String(dataPagination.value.page_size),
    })

    // agregar filtros opcionales
    if (filters.value.education && filters.value.education !== "All") {
      params.append('education', filters.value.education)
    }
    if (filters.value.marital_status && filters.value.marital_status !== "All") {
      params.append('marital_status', filters.value.marital_status)
    }
    if (filters.value.occupation && filters.value.occupation !== "All") {
      params.append('occupation', filters.value.occupation)
    }
    if (filters.value.income && filters.value.income !== "All") {
      params.append('income', filters.value.income)
    }
    // Solo agregar order_by y order_direction si tienen un valor seleccionado
    if (filters.value.order_by && filters.value.order_by !== "All") {
      params.append('order_by', filters.value.order_by.toLowerCase());
    }
    if (filters.value.order_direction && filters.value.order_direction !== "All") {
      params.append('order_direction', filters.value.order_direction.toUpperCase());
    }

    if (filters.value.export == true) {
      params.append('export', filters.value.export.toString())
    }

    // Concatenar los parámetros adicionales a la URL
    if (hasFilters) {
      url += `&${params.toString()}`;  // Concatenamos a la URL de los filtros
    } else {
      url += `?${params.toString()}`;  // En caso de no tener filtros, empezamos con '?'
    }
    console.log("Request URL:", url); // Verificar la URL construida

    // realizar la solicitud
    const response = await fetch(url)
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

// Cambiar página
function changePage(newPage: number) {
  if (newPage >= 1 && newPage <= dataPagination.value.total_pages) {
    loadData(newPage)
  }
}


// Observar cambios en filtros
watch(filters, () => {
  loadData(1); // Recargar la tabla cuando se cambia el filtro
}, {deep: true});

// Montar el componente y cargar la primera página
onMounted(() => {
  loadData(dataPagination.value.page);
});

</script>

<template>
  <main class="bg-purple-100 flex justify-center items-center py-16 min-h-screen">
    <div class="w-[90%] max-w-5xl border rounded-lg shadow-lg mt-16">
      <!--Filtros-->
      <div class="flex justify-between items-center px-4 py-2 bg-gray-100 border-t text-center text-xs">
        <div>
          <label>Filter Education</label>
          <select v-model="filters.education" class="border rounded px-2 py-1">
            <option v-for="option in filterOptionsEducation" :key="option" :value="option">{{ option }}</option>
          </select>
      </div>
      <div>
          <label>Filter Marital Status</label>
          <select v-model="filters.marital_status" class="border rounded px-2 py-1">
            <option v-for="option in filterOptionsMaritalStatus" :key="option" :value="option">{{ option }}</option>
          </select>
      </div>
      <div>
          <label>Filter Occupation</label>
          <select v-model="filters.occupation" class="border rounded px-2 py-1">
            <option v-for="option in filterOptionsOccupation" :key="option" :value="option">{{ option }}</option>
          </select>
      </div>
      <div>
          <label>Filter Income</label>
          <select v-model="filters.income" class="border rounded px-2 py-1">
            <option v-for="option in filterOptionsIncome" :key="option" :value="option">{{ option }}</option>
          </select>
      </div>
      <div>
          <label>Order By</label>
          <select v-model="filters.order_by" class="border rounded px-2 py-1">
            <option v-for="header in HeadersTable" :key="header" :value="header">{{ header }}</option>
          </select>
      </div>
      <div>
          <label>Order Direction</label>
          <select v-model="filters.order_direction" class="border rounded px-2 py-1">
            <option v-for="option in orderOptions" :key="option" :value="option">{{ option }}</option>
          </select>
      </div>
      <div>
          <label>Age Range</label>
          <div class="flex gap-2">
            <input type="number" v-model="filters.min_age" class="border rounded px-2 py-1 w-20" placeholder="Min" />
            <input type="number" v-model="filters.max_age" class="border rounded px-2 py-1 w-20" placeholder="Max" />
          </div>
        </div>
        <div class="py-2 px-2">
          <button class="bg-green-600 px-2 py-2 border-t rounded text-sm text-white" @click="filters.export=true">Export Data</button>
        </div>
      </div>
      
      <table class="table-auto w-full border-collapse">
      <thead>
        <tr class="bg-gray-200">
          <!-- Generar los encabezados dinámicamente -->
          <th v-for="header in HeadersTable" :key="header" class="border border-gray-300 px-4 py-2 text-xs font-medium text-gray-700">
            {{ (header) }}
          </th>
        </tr>
      </thead>
      <tbody>
        <!-- Generar filas dinámicamente -->
        <tr v-for="datarow in dataRecords" :key="datarow.id">
          <td class="border border-gray-300 px-4 py-2 text-center text-xs">
            {{ (datarow.Age) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Work_Class) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Education) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Marital_Status) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Occupation) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Relationship) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Race) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Sex) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Capital_Gain) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Capital_Loss) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Hours_Per_Week) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Native_Country) }}
          </td>
          <td class="border border-gray-300 px-2 py-2 text-center text-xs">
            {{ (datarow.Income) }}
          </td>
        </tr>
      </tbody>
     </table>
     <!-- Mostrar la paginación -->
      <div class="flex justify-between items-center px-4 py-2 bg-gray-100 border-t">
        <button class="px-4 py-2 bg-sky-800 text-white rounded disabled:opacity-50" @click="changePage(dataPagination.page - 1)">Before</button>
        <span class="text-sm text-gray-700">Page {{dataPagination.page}} from {{ dataPagination.total_pages }}</span>
        <button class="px-4 py-2 bg-sky-800 text-white rounded disabled:opacity-50" @click="changePage(dataPagination.page + 1)">Next</button>
        <span  class="text-sm text-gray-700">Total Data: {{ dataPagination.total_records }}</span>
      </div>
    </div>
  </main>
</template>
