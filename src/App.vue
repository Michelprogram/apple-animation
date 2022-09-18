<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import Header from "./components/Header.vue";
import Modeles from "./components/Modeles.vue"
</script>

<template>
  <Header class="header"/>
  <Modeles class="modeles"/>
</template>

<style scoped lang="scss">
.header{
  margin-bottom:1%;
}
</style>
