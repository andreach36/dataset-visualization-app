<script setup lang="ts">
import { HeadersTable } from '@/data_table/labels';
import { onMounted} from 'vue';
import { useDataStore } from '../stores/data';
import { useFilterStore } from '../stores/filter';
import { ref } from "vue";
import ModalComponent from "../components/ModalComponent.vue";

const dataStore = useDataStore()
const filterStore = useFilterStore()

const isModalOpened = ref(false)

const openModal = () => {
  isModalOpened.value = true
}
const closeModal = () => {
  isModalOpened.value = false
}

const submitHandler = async () => {
  await filterStore.updateUserFilters()
  dataStore.loadData(dataStore.dataPagination.page, filterStore.userFilters)
  closeModal()
}

// Cambiar página
function changePage(newPage: number) {
  if (newPage >= 1 && newPage <= dataStore.dataPagination.total_pages) {
    dataStore.loadData(newPage, filterStore.userFilters)
  }
}

// Montar el componente y cargar la primera página
onMounted(async () => {
  await filterStore.loadUserFilters(); // Espera a que los filtros se carguen
  dataStore.loadData(dataStore.dataPagination.page, filterStore.userFilters); // Carga los datos usando los filtros actualizados
});

// Exportar Data
function triggerExport() {
    filterStore.exportData = true;

    // Ejecutar lógica para exportar datos, por ejemplo:
   
    // Después de exportar, puedes resetearlo si es necesario:
    filterStore.exportData = false;
}

</script>

<template>
  <main class="bg-gray-100 flex justify-center items-center py-16 min-h-screen">
    <div class="w-[90%] max-w-5xl border rounded-lg shadow-lg mt-16">
      <!--Opciones Tabla-->
      <div class="flex justify-between items-center px-4 py-2 bg-gray-100 border-t text-center">
        <div class="py-2 px-2">
          <button class="bg-blue-800 px-2 py-2 border-t rounded text-sm text-white"  @click="openModal">Filter Options</button>
        </div>
        <ModalComponent :isOpen="isModalOpened" @modal-close="closeModal" @submit="submitHandler" name="first-modal">
        </ModalComponent>
        <div class="py-2 px-2">
          <button class="bg-green-600 px-2 py-2 border-t rounded text-sm text-white" @click="triggerExport">Export Data</button>
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
        <tr v-for="datarow in dataStore.dataRecords" :key="datarow.id">
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
        <button class="px-4 py-2 bg-sky-800 text-white rounded disabled:opacity-50" @click="changePage(dataStore.dataPagination.page - 1)">Before</button>
        <span class="text-sm text-gray-700">Page {{dataStore.dataPagination.page}} from {{ dataStore.dataPagination.total_pages }}</span>
        <button class="px-4 py-2 bg-sky-800 text-white rounded disabled:opacity-50" @click="changePage(dataStore.dataPagination.page + 1)">Next</button>
        <span  class="text-sm text-gray-700">Total Data: {{ dataStore.dataPagination.total_records }}</span>
      </div>
    </div>
  </main>
</template>
