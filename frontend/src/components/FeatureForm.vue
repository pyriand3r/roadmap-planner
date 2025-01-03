<template>
    <v-form class="mx-1 my-3">
        <v-text-field v-model="store.editFeature.title" label="title" variant="outlined"></v-text-field>
        <v-textarea v-model="store.editFeature.description" label="description" variant="outlined"></v-textarea>
        <v-select v-model="store.editFeature.type" variant="outlined" label="Type" prepend-icon="mdi-atom"
            :items="feature_type_list"></v-select>
        <v-select v-model="store.editFeature.milestone_id" variant="outlined" label="Milestone" prepend-icon="mdi-sign-direction"
            :items="milestoneStore.milestones" item-value="ID" clearable></v-select>
        <v-icon icon="mdi-clock"></v-icon><v-label>Estimate</v-label>
        <div class="d-flex flex-row mb-3 mt-1 w-100">
            <v-btn-toggle v-model="store.editFeature.estimate" :color="colorEstimate" group class="w-100 d-flex">
                <v-btn v-for="value in feature_estimate_list" :idx="value" :value="value"
                    class="flex-fill">
                    {{ value }}
                </v-btn>
            </v-btn-toggle>
        </div>
        <v-label>Customer score</v-label>
        <v-slider v-model="store.editFeature.customer_score" color="red-lighten-2" :min="0" :max="100" :step="5"
            show-ticks thumb-label prepend-icon="mdi-account">
            <template v-slot:append>
                <v-text-field v-model="store.editFeature.customer_score" density="compact" style="width: 80px"
                    type="number" variant="outlined" hide-details></v-text-field>
            </template>
        </v-slider>
        <v-label>Sales score</v-label>
        <v-slider v-model="store.editFeature.sales_score" color="light-blue-lighten-2" :min="0" :max="100" :step="5"
            show-ticks thumb-label prepend-icon="mdi-cash">
            <template v-slot:append>
                <v-text-field v-model="store.editFeature.sales_score" density="compact" style="width: 80px"
                    type="number" variant="outlined" hide-details></v-text-field>
            </template>
        </v-slider>
        <v-label>Dev score</v-label>
        <v-slider v-model="store.editFeature.dev_score" color="orange-lighten-2" :min="0" :max="100" :step="5"
            show-ticks thumb-label prepend-icon="mdi-wrench">
            <template v-slot:append>
                <v-text-field v-model="store.editFeature.dev_score" density="compact" style="width: 80px"
                    type="number" variant="outlined" hide-details></v-text-field>
            </template>
        </v-slider>
        <v-row class="justify-space-evenly">
            <v-btn variant="outlined" @click.prevent="$emit('clear')">clear</v-btn>
            <v-btn color="success" variant="outlined" @click.prevent="$emit('save')">save</v-btn>
        </v-row>
    </v-form>
</template>

<script setup lang="ts">
import { feature_type_list, feature_estimate_list } from '@/models'
import { useFeaturesStore } from '@/store/feature'
import { useMilestoneStore } from '@/store/milestone'
import { estimateColor } from '@/globals'
import {computed, defineEmits} from "vue"

defineEmits(['clear', 'save'])

const store = useFeaturesStore()
store.init()

const milestoneStore = useMilestoneStore()
milestoneStore.init()

const colorEstimate = computed((): string => {
  return estimateColor(store.editFeature.estimate)
})

</script>
