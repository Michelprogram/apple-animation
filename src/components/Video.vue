<script setup lang="ts"></script>

<template>
  <div class="container-video">
    <video autoplay muted width="250" ref="video">
      <source src="../../public/large.mp4" type="video/mp4" />
    </video>
  </div>
  <div class="fake-container"></div>
</template>

<style lang="scss" scoped>
.container-video {
  position: absolute;
  top: 45px;
  left: 0;
  width: 100%;
  height: 800px;
  z-index: -1;

  video{
    object-fit: fill;
    width: 100%;
  }
}

.fake-container{
    height: 800px;
}
</style>
