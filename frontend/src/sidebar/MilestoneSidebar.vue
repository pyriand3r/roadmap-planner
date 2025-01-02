
<template>
  <v-list>
    <v-list-item title="Milestones"></v-list-item>
    <v-list-item>
      <v-btn block variant="outlined" @click="openNewForm">
        <template #prepend>
          <v-icon icon="mdi-plus-thick"></v-icon>
        </template>
        Create new milestone
      </v-btn>
    </v-list-item>
    <v-list-item v-if="store.showForm">
      <milestone-form @clear="clear" @save="save"/>
    </v-list-item>
  </v-list>
</template>

<script setup lang="ts">
import { milestoneSkel } from '@/models'
import { useMilestoneStore } from '@/store/milestone'
import MilestoneForm from '@/components/MilestoneForm.vue'

const store = useMilestoneStore()
store.init()

function openNewForm() {
  store.editMilestone = milestoneSkel()
  store.showForm = true
}

function clear() {
  store.editMilestone = milestoneSkel()
  store.showForm = false
}

function save() {
  const milestone = JSON.parse(JSON.stringify(store.editMilestone))
  store.add(milestone)
  clear()
}
</script>