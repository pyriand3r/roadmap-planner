<template>
  <v-container class="d-flex flex-column ga-6 overflow-y-auto align-center justify-center">
    <transition-group name="list">
      <feature-card v-for="(feature) in store.sorted" :key="feature.ID" :feature="feature" @edit="edit"
                    @remove="remove"/>
    </transition-group>
  </v-container>
</template>

<script setup lang="ts">
import FeatureCard from "@/components/FeatureCard.vue";
import {useFeaturesStore} from "@/store";
import type {Feature} from "@/models";

const store = useFeaturesStore()
store.init()

function edit(feature: Feature) {
  store.editFeature = JSON.parse(JSON.stringify(feature))
  store.showForm = true
}

function remove(feature: Feature) {
  store.remove(feature)
}

</script>

<style scoped>
.list-move, /* apply transition to moving elements */
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

/* ensure leaving items are taken out of layout flow so that moving
   animations can be calculated correctly. */
.list-leave-active {
  position: absolute;
}
</style>
