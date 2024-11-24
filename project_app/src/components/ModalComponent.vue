<script setup lang="ts">
import { defineProps, defineEmits, ref } from "vue";
import {onClickOutside} from '@vueuse/core'
import { useFilterStore } from "@/stores/filter";
import { HeadersTable, filterOptionsEducation, filterOptionsOccupation, filterOptionsMaritalStatus, filterOptionsIncome, orderOptions } from '@/data_table/labels';

const filterStore = useFilterStore()

const props = defineProps({
  isOpen: Boolean,
})

const emit = defineEmits(["submit"])

const target = ref(null)
onClickOutside(target, ()=>emit('submit'))
</script>

<template>
    <div v-if="isOpen" class="fixed z-50 top-0 left-0 w-full h-full bg-black/50">
    <div class="modal-wrapper">
      <div class="w-80 m-[150px_auto] py-5 px-8 bg-white rounded-md shadow-md flex flex-col space-y-4 items-center" ref="target">
        <div>
          <label>Filter Education</label>
          <select v-model="filterStore.userFilters.education" class="border rounded px-2 py-1">
            <option v-for="option in filterOptionsEducation" :key="option" :value="option">{{ option }}</option>
          </select>
        </div>
        <div>
          <label>Filter Marital Status</label>
          <select v-model="filterStore.userFilters.marital_status" class="border rounded px-2 py-1">
            <option v-for="option in filterOptionsMaritalStatus" :key="option" :value="option">{{ option }}</option>
          </select>
        </div>
        <div>
          <label>Filter Occupation</label>
          <select v-model="filterStore.userFilters.occupation" class="border rounded px-2 py-1">
            <option v-for="option in filterOptionsOccupation" :key="option" :value="option">{{ option }}</option>
          </select>
        </div>
        <div>
          <label class="px-2">Filter Income</label>
          <select v-model="filterStore.userFilters.income" class="border rounded px-2 py-1">
            <option v-for="option in filterOptionsIncome" :key="option" :value="option">{{ option }}</option>
          </select>
        </div>
        <div>
          <label class="px-2">Order By</label>
          <select v-model="filterStore.userFilters.order_by" class="border rounded px-2 py-1">
            <option v-for="header in HeadersTable" :key="header" :value="header">{{ header }}</option>
          </select>
        </div>
        <div>
          <label class="px-2">Order Direction</label>
          <select v-model="filterStore.userFilters.order_direction" class="border rounded px-2 py-1">
            <option v-for="option in orderOptions" :key="option" :value="option">{{ option }}</option>
          </select>
        </div>
        <div>
          <label>Age Range</label>
          <div class="flex gap-2 py-2">
            <input type="number" v-model="filterStore.userFilters.min_age" class="border rounded px-2 py-1 w-20" placeholder="Min" />
            <input type="number" v-model="filterStore.userFilters.max_age" class="border rounded px-2 py-1 w-20" placeholder="Max" />
          </div>
        </div>
        <div class="modal-footer">
          <slot name="footer">
            <div>
              <button class="bg-blue-800 px-2 py-2 border-t rounded text-white" @click.stop="emit('submit')">Submit</button>
            </div>
          </slot>
        </div>
      </div>
    </div>
  </div>
</template>
