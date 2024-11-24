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
async function exportData() {
  try {
    const baseUrl = "http://localhost:3333/data/filter"; // Asegúrate de que esta URL esté correcta
    const params = new URLSearchParams();

    // Añadir filtros si es necesario
    if (filterStore.userFilters.min_age) {
      params.append('min_age', filterStore.userFilters.min_age);
    }
    if (filterStore.userFilters.max_age) {
      params.append('max_age', filterStore.userFilters.max_age);
    }
    if (filterStore.userFilters.education && filterStore.userFilters.education !== "All") {
      params.append('education', filterStore.userFilters.education);
    }
    if (filterStore.userFilters.marital_status && filterStore.userFilters.marital_status !== "All") {
      params.append('marital_status', filterStore.userFilters.marital_status);
    }
    if (filterStore.userFilters.occupation && filterStore.userFilters.occupation !== "All") {
      params.append('occupation', filterStore.userFilters.occupation);
    }
    if (filterStore.userFilters.income && filterStore.userFilters.income !== "All") {
      params.append('income', filterStore.userFilters.income);
    }
    // Añadir el parámetro de exportación
    params.append('export', 'true');

    const response = await fetch(`${baseUrl}?${params.toString()}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${sessionStorage.getItem('token')}`,
      },
    });

    if (response.ok) {
      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = 'filteredData.csv';  // Cambia el nombre del archivo si lo deseas
      document.body.appendChild(a);
      a.click();
      a.remove();
    } else {
      console.error("Failed to export data");
    }
    } catch (error) {
        console.error("Error exporting data:", error);
    }
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
          <button class="bg-green-600 px-2 py-2 border-t rounded text-sm text-white" @click="exportData">Export Data</button>
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
