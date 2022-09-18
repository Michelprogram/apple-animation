<script setup lang="ts"></script>

<template>
  <div class="container-header-2">
    <div>
      <p>Modèles 14" et 16"</p>
    </div>
    <div class="informations">
      <div>
        <p>Macbook Pro</p>
      </div>
      <div>
        <p>Présentation</p>
        <p>Pourquoi choisir le Mac</p>
        <p>Caractéristiques</p>
        <div class="buy">
            <p>Acheter</p>
        </div>
      </div>
    </div>
    <div class="hr"></div>
  </div>
</template>

<style scoped lang="scss">

@use "../assets/index.scss" as color;

$little-font-size:0.75em;
$gray-opacity:0.64;

@mixin swapcolors($opacity){
    color: color.$header-color;
    opacity: $opacity;

    animation-duration: 3s;
    animation-name: colorSwap;

    @keyframes colorSwap{
        0%{
            color:white;
        }
        58%{
            color: white;
        } 
        100%{
            color: color.$header-color;
        }
    }
}

.container-header-2 {
  margin: auto;
  width: 70%;
  display: flex;
  flex-direction: column;
  z-index: 1;

  &>div:nth-child(1){
    font-size: $little-font-size;
    line-height: 20px;

    @include swapcolors($gray-opacity);

  }

  .informations{
    display: flex;
    justify-content: space-between;
    margin-bottom:1%;

    &>div:nth-child(1){
        font-size: 1.3em;
        width: 30%;

        @include swapcolors(1);

    }

    &>div:nth-child(2){
        cursor:pointer;
        display: flex;
        align-items: center;
        justify-content: space-between;
        font-size: $little-font-size;
        width: 42%;

        &>p:nth-child(1){
            @include swapcolors($gray-opacity);
        }

        &>p:nth-child(n+2){
            @include swapcolors(1);

            &:hover{
                color:color.$header-blue-button;
            }
        }
    }

    .buy{
        background-color: color.$header-blue-button;
        color:color.$header-text-color-hover;
        padding:5px 15px 5px 15px;
        border-radius: 15px;

        &:hover{
            background-color:color.$header-blue-button-hover;
        }
    }

  }

  &>div:last-child{
    width: 100%;
    height: 1px;
    background-color: rgba(128, 128, 128, 0.291);
  }
}


</style>
