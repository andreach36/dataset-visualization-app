<script setup lang="ts">
import { ref } from 'vue';

// Define los encabezados a partir de la interfaz UserRecordTable
const headers = ref([
  "Age",
  "Workclass",
  "FinalWeight",
  "Education",
  "Education_Num",
  "Marital_Status",
  "Occupation",
  "Relationship",
  "Race",
  "Sex",
  "Capital_Gain",
  "Capital_Loss",
  "Hours_Per_Week",
  "Native_Country",
  "Income_Range",
]);

</script>

<template>
  <main class="bg-gray-100 flex justify-center min-h-screen py-16">
    <div class="overflow-x-auto w-[90%] max-w-4xl border border-gray-300 rounded-md shadow-md bg-white mt-16">
      <table class="table-auto w-full border-collapse">
      <thead>
        <tr class="bg-gray-200">
          <!-- Generar los encabezados dinámicamente -->
          <th v-for="header in headers" :key="header" class="border border-gray-300 px-4 py-2 text-xs font-medium text-gray-700">
            {{ (header) }}
          </th>
        </tr>
      </thead>
      <tbody>
        <!-- Aquí puedes agregar dinámicamente las filas si tienes datos -->
        <tr>
          <td v-for="header in headers" :key="header" class="border border-gray-300 px-4 py-2">
            <!-- Ejemplo de valores predeterminados -->
            -
          </td>
        </tr>
      </tbody>
     </table>
    </div>
  </main>
</template>
