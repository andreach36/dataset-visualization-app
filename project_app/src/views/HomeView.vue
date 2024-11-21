<script setup lang="ts">
import { type DataPagination, type DataRecord } from '@/types';
import { onMounted, ref } from 'vue';



// los headers de la tabla
const HeadersTable = [
  "Age",
  "Education",
  "Marital Status",
  "Occupation",
  "Income",
]


// se crea lista de datos de la tabla
const dataRecords = ref<DataRecord[]>([])

const dataPagination = ref<DataPagination>({
  page: 1,
  page_size: 10,
  total_pages: 0,
  total_records: 0,
})

// función para cargar datos de la tabla
async function loadData (page: number) {
  try {
    const response = await fetch(`http://localhost:3333/data?page=${page}&limit=${dataPagination.value.page_size}`)
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

// Montar el componente y cargar la primera página
onMounted(() => {
  loadData(dataPagination.value.page);
});

</script>

<template>
  <main class="bg-gray-100 flex justify-center items-center py-16 min-h-screen">
    <div class="w-[90%] max-w-5xl border rounded-lg shadow-lg mt-16">
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
          <td class="border border-gray-300 px-4 py-2 text-center text-xs">
            {{ (datarow.Education) }}
          </td>
          <td class="border border-gray-300 px-4 py-2 text-center text-xs">
            {{ (datarow.Marital_Status) }}
          </td>
          <td class="border border-gray-300 px-4 py-2 text-center text-xs">
            {{ (datarow.Occupation) }}
          </td>
          <td class="border border-gray-300 px-4 py-2 text-center text-xs">
            {{ (datarow.Income) }}
          </td>
        </tr>
      </tbody>
     </table>
     <!-- Mostrar la paginación -->
      <div class="flex justify-between items-center px-4 py-2 bg-gray-100 border-t">
        <button class="px-4 py-2 bg-gray-800 text-white rounded disabled:opacity-50" @click="changePage(dataPagination.page - 1)">Anterior</button>
        <span class="text-sm text-gray-700">Página actual {{dataPagination.page}} de {{ dataPagination.total_pages }}</span>
        <button class="px-4 py-2 bg-gray-800 text-white rounded disabled:opacity-50" @click="changePage(dataPagination.page + 1)">Siguiente</button>
        <span  class="text-sm text-gray-700">Total Datos: {{ dataPagination.total_records }}</span>
      </div>
    </div>
  </main>
</template>
