<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import Header from "./components/Header.vue";
import Modeles from "./components/Modeles.vue";
import Video from "./components/Video.vue";
import Text from "./components/Text.vue";
import Animation from "./components/Animation.vue";
</script>

<template>
  <div class="before-animation">
    <Header class="container" />
    <Modeles class="container" />
    <Video />
    <Text class="container" />
  </div>

  <Animation class="container" />
</template>

<style scoped lang="scss">

.before-animation{
  height: 1500px;
}


.container {
  margin-bottom: 1%;

  &:nth-child(2) {
    position: relative;
  }
}
</style>
