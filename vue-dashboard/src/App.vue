<template>
  <div id="app" class="app-cls">
    <router-view />
  </div>
</template>

<style lang="scss">
@import './styles/base.scss';

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background: yellow;
}

.app-cls {
  border: 1px solid red;
  height: 100%;
}
</style>
