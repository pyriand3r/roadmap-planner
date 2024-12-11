<template>
  <v-app>
    <v-layout class="rounded rounded-md h-100 w-100 position-fixed">
      <v-app-bar title="Roadmap Planner">
        <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      </v-app-bar>
      <v-navigation-drawer v-model="drawer" location="right" width="400">
        <features-sidebar/>
      </v-navigation-drawer>
      <v-main class="overflow-auto">
        <features/>
      </v-main>
    </v-layout>
  </v-app>
</template>

<script setup lang="ts">
import Features from '@/views/Features.vue'
import FeaturesSidebar from "@/sidebar/FeaturesSidebar.vue";
import {ref} from "vue";
const drawer = ref(true)
</script>
