<template>
  <div id="app" class="app-cls">
    <router-view />
  </div>
</template>

<script>
export default {
  mounted () {
    let info = localStorage.getItem('loginUserinfo')
    info = JSON.parse(info)
    if (typeof info.data.accessToken === 'string' && info.data.accessToken.length > 0) {
      this.$nextTick(() => {
        this.$router.push('/home')
      })
    } else {
      console.log('[info]', JSON.parse(info))
      this.$nextTick(() => {
        this.$router.push('/login')
      })
    }
  }
}
</script>

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
