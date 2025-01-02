<template>
  <v-card width="90%" variant="outlined" density="compact" rounded="xl"
    :title="milestone.title" @click="$emit('edit', milestone)">
    <template #title>
      {{ milestone.title }}
    </template>
    <v-card-text>
      {{ milestone.description }}
    </v-card-text>
    <v-card-text v-if="featureList.length > 0" class="d-flex flex-row">
      <details class="flex-fill" @click.stop>
        <summary class="mb-2">Features: {{ featureList.length }}</summary>
        <div class="d-flex flex-column ga-4">
          <feature-card v-for="(feature, idx) in featureList" :key="idx" :feature="feature" class="w-100" />
        </div>
      </details>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { computed, defineProps, ref } from 'vue'
import type { Milestone } from '@/models'
import FeatureCard from './FeatureCard.vue'
import { useFeaturesStore } from '@/store/feature'

const features = useFeaturesStore()
features.init()

const props = defineProps<{
  milestone: Milestone
}>()

const extended = ref(false)
const featureList = computed(() => {
  return features.byMilestone(props.milestone.ID)
})

function toggleExtend() {
  extended.value = !extended.value
}
</script>
