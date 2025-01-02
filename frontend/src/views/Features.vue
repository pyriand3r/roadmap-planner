<template>
  <v-navigation-drawer v-model="store.sidebarShow" location="right" width="400">
    <features-sidebar />
  </v-navigation-drawer>
  <v-main class="overflow-auto">
  <v-container class="d-flex flex-column ga-6 overflow-y-auto align-center justify-center">
    <transition-group name="list">
      <feature-card v-for="(feature) in store.sorted" :key="feature.ID" :feature="feature" @edit="edit"
        @remove="remove" />
    </transition-group>
  </v-container>
  </v-main>
</template>

<script setup lang="ts">
import { ref } from "vue";
import FeatureCard from "@/components/FeatureCard.vue"
import FeaturesSidebar from "@/sidebar/FeaturesSidebar.vue"
import { useFeaturesStore } from "@/store/feature"
import { useMilestoneStore } from "@/store/milestone"
import type { Feature } from "@/models"

const drawer = ref(true)

const store = useFeaturesStore()
store.init()

const milestoneStore = useMilestoneStore()
milestoneStore.init()

function edit(feature: Feature) {
  store.editFeature = JSON.parse(JSON.stringify(feature))
  store.showForm = true
}

function remove(feature: Feature) {
  store.remove(feature)
}

</script>

<style scoped>
.list-move,
/* apply transition to moving elements */
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
