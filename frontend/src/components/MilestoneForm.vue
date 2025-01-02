<template>
  <v-form class="mx-1 my-3">
    <v-text-field v-model="store.editMilestone.title" label="title" variant="outlined" />
    <v-textarea v-model="store.editMilestone.description" label="description" variant="outlined" />
        <v-row class="justify-space-evenly">
            <v-btn variant="outlined" @click.prevent="$emit('clear')">clear</v-btn>
            <v-btn color="success" variant="outlined" @click.prevent="$emit('save')">save</v-btn>
        </v-row>
  </v-form>
</template>

<script setup lang="ts">
import { useMilestoneStore } from '@/store/milestone';
import { defineEmits } from 'vue';

defineEmits(['clear', 'save'])

const store = useMilestoneStore()
store.init()

</script>