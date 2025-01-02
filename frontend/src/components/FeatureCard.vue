<template>
  <v-card width="90%" variant="outlined" density="compact" rounded="xl"
    :title="feature.title + ' (' + feature.total_score + ')'" @click="$emit('edit', feature)">
    <template #title>
      {{ feature.title }}
    </template>
    <template #prepend>
        <v-chip variant="outlined" prepend-icon="mdi-sigma">{{ feature.total_score }}</v-chip>
    </template>
    <template #append>
      <div class="d-flex flex-row ga-2">
        <v-chip variant="outlined" prepend-icon="mdi-atom" :color="featureColor(feature.type)">{{ feature.type }}</v-chip>
        <v-chip variant="outlined" prepend-icon="mdi-clock" :color="estimateColor(feature.estimate)">{{ feature.estimate }}</v-chip>
        <v-chip v-if="feature.milestone_id" variant="outlined" prepend-icon="mdi-sign-direction">{{ milestoneName }}</v-chip>
        <v-icon icon="mdi-delete" @click.stop="$emit('remove', feature)"></v-icon>
      </div>
    </template>
    <v-card-text v-if="feature.description !== ''" :class="{ extended: extended }" class="pb-0">
      {{ feature.description }}
    </v-card-text>
    <v-card-text>
      <div class="d-flex flex-row ga-4 ma-2">
        <div class="w-33 d-flex flex-row ga-4 align-center">
          <v-icon icon="mdi-account"></v-icon>
          <v-progress-linear rounded="lg" height="8" v-model="feature.customer_score" color="red-lighten-2" />
        </div>
        <div class="w-33 d-flex flex-row ga-4 align-center">
          <v-icon icon="mdi-cash"></v-icon>
          <v-progress-linear rounded="lg" height="8" v-model="feature.sales_score" color="light-blue-lighten-2" />
        </div>
        <div class="w-33 d-flex flex-row ga-4 align-center">
          <v-icon icon="mdi-wrench"></v-icon>
          <v-progress-linear rounded="lg" height="8" v-model="feature.dev_score" color="orange-lighten-2" />
        </div>
      </div>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { computed, defineProps, ref } from 'vue'
import type { Feature } from '@/models'
import { estimateColor, featureColor } from '@/globals'
import { useMilestoneStore } from '@/store/milestone'

const milestones = useMilestoneStore()
milestones.init()

const props = defineProps<{
  feature: Feature
}>()

const milestoneName = computed(() => {
  const milestone = milestones.byID(props.feature.milestone_id)
  return milestone?.title
})

const extended = ref(false)

function toggleExtend() {
  extended.value = !extended.value
}

</script>
