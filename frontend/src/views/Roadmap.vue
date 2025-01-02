<template>
  <v-navigation-drawer v-model="store.sidebarShow" location="right" width="400">
    <milestone-sidebar />
  </v-navigation-drawer>
  <v-main class="overflow-auto">
  <v-container class="d-flex flex-column ga-6 overflow-y-auto align-center justify-center">
    <transition-group name="list">
      <milestone-card v-for="(milestone) in store.milestones" :key="milestone.ID" :milestone="milestone" @edit="edit"
        @remove="remove" />
    </transition-group>
  </v-container>
  </v-main>
</template>

<script setup lang="ts">
import MilestoneCard from "@/components/MilestoneCard.vue"
import type { Milestone } from "@/models"
import MilestoneSidebar from "@/sidebar/MilestoneSidebar.vue"
import { useMilestoneStore } from "@/store/milestone"

const store = useMilestoneStore()
store.init()

function edit(milestone: Milestone) {
  store.editMilestone = JSON.parse(JSON.stringify(milestone))
  store.showForm = true
}
function remove(milestone: Milestone) {
  store.remove(milestone)
}
</script>