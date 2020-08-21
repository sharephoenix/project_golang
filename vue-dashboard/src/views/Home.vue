<template>
  <div class="home">
    <img alt="Vue logo"
         src="../assets/logo.png">
         this is my home
  </div>
</template>

<script>
export default {
  name: 'Home',
  mounted () {
    console.log(this.$route)
  }
}
</script>
