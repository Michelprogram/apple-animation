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

.container-header-2 {
  margin: auto;
  width: 70%;
  display: flex;
  flex-direction: column;

  &>div:nth-child(1){
    font-size: $little-font-size;
    opacity: $gray-opacity;
    color:color.$header-color;
    line-height: 20px;
  }

  .informations{
    display: flex;
    justify-content: space-between;
    margin-bottom:1%;

    &>div:nth-child(1){
        font-size: 1.3em;
        color:color.$header-color;
        width: 30%;
    }

    &>div:nth-child(2){
        cursor:pointer;
        display: flex;
        align-items: center;
        justify-content: space-between;
        font-size: $little-font-size;
        width: 42%;

        &>p:nth-child(1){
            opacity: $gray-opacity;
        }

        &>p:nth-child(n+2){
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
