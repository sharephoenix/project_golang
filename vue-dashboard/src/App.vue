<template>
  <div id="app" class="app-cls">
    <router-view />
  </div>
</template>

<script>
export default {
  mounted () {
    console.log('[App]', this.$route.path)
    let info = localStorage.getItem('loginUserinfo')
    info = JSON.parse(info)
    if (typeof info.data.accessToken === 'string' && info.data.accessToken.length > 0) {
      if (this.$route.path === '/home') {
        return
      }
      this.$router.push('/home')
    } else {
      console.log('[info]', JSON.parse(info))
      this.$router.push('/login')
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
