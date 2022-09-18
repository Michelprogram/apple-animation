<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  mounted() {
    //this.$refs.video.play()
  },
});
</script>

<template>
  <div>
    <video autoplay muted width="250" ref="video">
      <source src="../../public/large.mp4" type="video/mp4" />
    </video>
  </div>
</template>

<style lang="scss" scoped>

div {
  position: absolute;
  top:45px;
  left:0;
  width: 100%;
  height: 800px;
  z-index: 0;
}

video {
  object-fit: fill;
  width: 100%;
}
</style>
