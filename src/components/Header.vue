<script setup lang="ts"></script>

<template>
  <div class="container-header">
    <ul>
      <li>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512">
          <path
            fill="#f5f5f798"
            d="M318.7 268.7c-.2-36.7 16.4-64.4 50-84.8-18.8-26.9-47.2-41.7-84.7-44.6-35.5-2.8-74.3 20.7-88.5 20.7-15 0-49.4-19.7-76.4-19.7C63.3 141.2 4 184.8 4 273.5q0 39.3 14.4 81.2c12.8 36.7 59 126.7 107.2 125.2 25.2-.6 43-17.9 75.8-17.9 31.8 0 48.3 17.9 76.4 17.9 48.6-.7 90.4-82.5 102.6-119.3-65.2-30.7-61.7-90-61.7-91.9zm-56.6-164.2c27.3-32.4 24.8-61.9 24-72.5-24.1 1.4-52 16.4-67.9 34.9-17.5 19.8-27.8 44.3-25.6 71.9 26.1 2 49.9-11.4 69.5-34.3z"
          />
        </svg>
      </li>
      <li>Store</li>
      <li>Mac</li>
      <li>iPad</li>
      <li>iPhone</li>
      <li>Watch</li>
      <li>AirPods</li>
      <li>TV & Maison</li>
      <li>Exclusivités Apple</li>
      <li>Accessoires</li>
      <li>Assistance</li>
      <li>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
          <path
            fill="#f5f5f798"
            d="M416 208c0 45.9-14.9 88.3-40 122.7L502.6 457.4c12.5 12.5 12.5 32.8 0 45.3s-32.8 12.5-45.3 0L330.7 376c-34.4 25.2-76.8 40-122.7 40C93.1 416 0 322.9 0 208S93.1 0 208 0S416 93.1 416 208zM208 352c79.5 0 144-64.5 144-144s-64.5-144-144-144S64 128.5 64 208s64.5 144 144 144z"
          />
        </svg>
      </li>
      <li>
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
          <path
            fill="#f5f5f798"
            d="M160 112c0-35.3 28.7-64 64-64s64 28.7 64 64v48H160V112zm-48 48H48c-26.5 0-48 21.5-48 48V416c0 53 43 96 96 96H352c53 0 96-43 96-96V208c0-26.5-21.5-48-48-48H336V112C336 50.1 285.9 0 224 0S112 50.1 112 112v48zm24 96c-13.3 0-24-10.7-24-24s10.7-24 24-24s24 10.7 24 24s-10.7 24-24 24zm200-24c0 13.3-10.7 24-24 24s-24-10.7-24-24s10.7-24 24-24s24 10.7 24 24z"
          />
        </svg>
      </li>
    </ul>
  </div>
</template>

<style scoped lang="scss">
@use "../assets/index.scss" as color;

.container-header {
  height: 45px;
  background-color: color.$header-color;
  color: color.$header-text-color;
  font-weight: 400;
  font-size: 0.7em;
  display: flex;
  align-items: center;
  justify-content: center;

  ul {
    list-style-type: none;
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 70%;
    height: 100%;

    li {
      cursor: pointer;
      transition: color 0.2s ease-in;
      &:hover {
        color: color.$header-text-color-hover;
      }
    }
  }

  svg {
    width: 13px;
    &:hover {
      path {
        transition: fill 0.2s ease-in;
        fill: color.$header-text-color-hover;
      }
    }
  }
}
</style>
