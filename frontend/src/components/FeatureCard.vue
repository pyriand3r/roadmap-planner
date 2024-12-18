<template>
  <v-card width="90%" variant="outlined" density="compact" rounded="xl"
    :title="feature.title + ' (' + feature.total_score + ')'" @click="$emit('edit', feature)">
    <template #title>
      {{ feature.title }}
    </template>
    <template #append>
      <div class="d-flex flex-row ga-2">
        <v-chip variant="outlined">Count: {{ feature.total_score }}</v-chip>
        <v-chip variant="outlined" :color="featureColor(feature.type)">{{ feature.type }}</v-chip>
        <v-chip variant="outlined" :color="estimateColor(feature.estimate)">{{ feature.estimate }}</v-chip>
        <v-icon icon="mdi-delete" @click.stop="$emit('remove', feature)"></v-icon>
      </div>
    </template>
    <v-card-text :class="{ extended: extended }">
      {{ feature.description }}
    </v-card-text>
    <v-card-text>
      <div class="d-flex flex-row ga-4 ma-2">
        <div class="w-33">
          Customer Score: {{ feature.customer_score }}
          <v-progress-linear rounded="lg" height="8" v-model="feature.customer_score" color="red-lighten-2" />
        </div>
        <div class="w-33">
          Sales Score: {{ feature.sales_score }}
          <v-progress-linear rounded="lg" height="8" v-model="feature.sales_score" color="light-blue-lighten-2" />
        </div>
        <div class="w-33">
          Dev Score: {{ feature.dev_score }}
          <v-progress-linear rounded="lg" height="8" v-model="feature.dev_score" color="orange-lighten-2" />
        </div>
      </div>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { defineProps, ref } from 'vue'
import type { Feature } from '@/models'
import { estimateColor, featureColor } from '@/globals'

defineProps<{
  feature: Feature
}>()

const extended = ref(false)

function toggleExtend() {
  extended.value = !extended.value
}

</script>
