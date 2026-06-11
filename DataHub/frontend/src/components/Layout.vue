<template>
  <div class="layout-wrapper">
    <Sidebar />
    <div class="layout-content">
      <router-view />
    </div>
  </div>
</template>

<script setup>
import Sidebar from './Sidebar.vue'
import { RouterView } from 'vue-router'
</script>

<style scoped>
.layout-wrapper {
  display: flex;
  min-height: 100vh;
  background-color: #0f172a;
}

.layout-content {
  flex: 1;
  margin-left: 260px;
  padding: 24px;
  overflow-y: auto;
}
</style>
