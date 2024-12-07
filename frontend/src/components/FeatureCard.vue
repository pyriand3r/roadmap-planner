<template>
    <v-card width="90%" variant="outlined" density="compact" :border="featureTypeBorder(feature.type)" rounded="xl"
        :title="feature.title + ' (' + feature.total_score + ')'" @click="$emit('edit', feature)">
        <template #append>
            <v-icon icon="mdi-delete" @click="$emit('remove', feature)"></v-icon>
        </template>
        <v-card-text>
            <div class="d-flex flex-row ga-4 ma-2">
                <div class="w-50">
                    Customer Score: {{ feature.customer_score }}
                    <v-progress-linear rounded="lg" height="8" v-model="feature.customer_score" color="red-lighten-2" />

                </div>
                <div class="w-50">
                    Sales Score: {{ feature.sales_score }}
                    <v-progress-linear rounded="lg" height="8" v-model="feature.sales_score"
                        color="light-blue-lighten-2" />

                </div>
            </div>
            <div class="d-flex flex-row ga-4 ma-2">
                <div class="w-50">
                    Internal Score: {{ feature.internal_score }}
                    <v-progress-linear rounded="lg" height="8" v-model="feature.internal_score"
                        color="orange-lighten-2" />
                </div>
                <div class="w-50">
                    Impact Score: {{ feature.impact_score }}
                    <v-progress-linear rounded="lg" height="8" v-model="feature.impact_score" color="teal-lighten-2" />
                </div>
            </div>
        </v-card-text>
    </v-card>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
import type { Feature } from '../models'

defineProps<{
    feature: Feature
}>()

function featureTypeBorder(type: string): string {
    let style = "lg opacity-50 "
    switch (type) {
        case "Bug":
            style += " error"
            break
        case "Feature":
            style += " warning"
            break
        case "Improvement":
            style += " info"
    }
    return style
}

</script>