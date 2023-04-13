<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  data() {
    return {
      images: new Array<HTMLImageElement>(),
      canvas: CanvasRenderingContext2D,
      offset: Number(0),
      position: Number(0),
      frameCount: Number(86),
    };
  },
  mounted() {
    this.preloadImage;
    //@ts-ignore
    this.canvas = this.$refs.canvas.getContext("2d");
    this.updateImage;
    window.addEventListener("scroll", this.scrollEvents);
  },
  computed: {
    preloadImage() {
      for (let index = 0; index < 87; index++) {
        const img = new Image();
        img.src = `https://www.apple.com/105/media/us/macbook-pro-14-and-16/2021/a1c5d17e-d8e4-4fa8-b70a-bc61bd266412/anim/hero-specs//large/large_${index
          .toString()
          .padStart(4, "0")}.jpg`;
        this.images.push(img);
      }
    },
    updateImage() {
      //@ts-ignore
      this.canvas.drawImage(this.images[this.offset], 0, 0);
    },
    appearText() {
      if (this.offset == 69) {
        return "container-text-info-mac show";
      }
      return "container-text-info-mac hide";
    },
  },
  methods: {
    scrollEvents(_: Event) {
      //Arrivé à 1500 l'animation se déclenche

      if (window.scrollY >= 1500) {
        const html = document.documentElement;
        const scrollTop = html.scrollTop - 1500;
        const maxScrollTop = html.scrollHeight - window.innerHeight;
        const scrollFraction = scrollTop / maxScrollTop;

        this.offset = Math.min(
          this.frameCount - 1,
          Math.ceil(scrollFraction * this.frameCount)
        );
        this.updateImage;
      }
    },
  },
});
</script>

<template>
  <div class="container-animation">
    <canvas ref="canvas" width="1336" height="786" />

    <div :class="appearText">
      <p>Jusqu'à</p>
      <div>
        <span>13x</span>
        <p>plus rapide pour les performances graphiques</p>
      </div>
    </div>

    <div :class="appearText">
      <p>Jusqu'à</p>
      <div>
        <span>21</span>
        <p>
          heures <br />
          d’autonomie
        </p>
      </div>
    </div>

    <div :class="appearText">
      <p>Jusqu'à</p>
      <div>
        <span>11x</span>
        <p>plus rapide pour l’apprentissage automatique</p>
      </div>
    </div>

    <div :class="appearText">
      <p>Jusqu'à</p>
      <div>
        <span>3,7x</span>
        <p>plus rapide pour les performances de CPU</p>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
@use "../assets/variables" as color;

.container-animation {
  width: 80%;
  height: 800vh;
  justify-content: center;
  margin: auto;
  position: relative;

  canvas {
    position: sticky;
    top: 0px;
    width: 100%;
  }

  .container-text-info-mac {
    position: fixed;
    color: rgba(0, 0, 0, 0.5);
    width: 22%;
    overflow: hidden;

    &:nth-child(2) {
      bottom: 48%;
      left: 16%;
    }

    &:nth-child(3) {
      bottom: 10%;
      left: 30%;
    }

    &:nth-child(4) {
      bottom: 19%;
      right: 27%;
    }
    &:nth-child(5) {
      top: 21%;
      right: 19%;
    }

    p:first-child {
      position: relative;
      font-size: 1.02em;
      margin-top: 15px;
      transition: opacity 2s, transform 3s cubic-bezier(0, 0, 0.2, 1);

      &::before {
        position: absolute;
        left: 0%;
        top: -15px;
        background-color: color.$title-color;
        content: "";
        opacity: 1;
        transition: all 2s;
      }
    }

    div {
      display: flex;

      span:first-child {
        font-size: 6.25em;
        letter-spacing: 1px;
        font-weight: 600;
        line-height: 1.04;
        color: color.$title-color;

        transition: opacity 2s, transform 2.5s cubic-bezier(0, 0, 0.2, 1);
      }
      p {
        align-self: center;
        margin-left: 5%;
        font-size: 1.1em;
        transition: opacity 2.15s 0.15s,
          transform 3.15s cubic-bezier(0, 0, 0.2, 1) 0.15s;
      }
    }

    &.hide {
      p:first-child {
        transform: translateY(50px);
        opacity: 0;
        &::before {
          width: 0px;
          height: 0px;
        }
      }
      div {
        * {
          opacity: 0;
        }

        span:first-child {
          transform: translateY(100px);
        }
        p {
          transform: translateY(130px);
        }
      }
    }

    &.show {
      p:first-child {
        transform: translateY(0px);
        &::before {
          width: 40px;
          height: 3px;
        }
      }

      div {
        * {
          transform: translateY(0px);
        }
      }
    }
  }
}
</style>
