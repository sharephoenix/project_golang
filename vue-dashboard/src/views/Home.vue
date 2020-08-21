<template>
  <div class="home">
    <img alt="Vue logo"
         src="../assets/logo.png">
         this is my home {{ getCount }}

         <div @click="testAction"> aaaa </div>
  </div>
</template>

<script>
export default {
  name: 'Home',
  computed: {
    getCount () {
      return this.$store.state.count
    }
  },
  methods: {
    testAction () {
      console.log('[testAction]')
      // this.$store.commit('increment', {
      //   amount: 100
      // })
      this.$store.dispatch('increment', { amount: 101 })
    }
  },
  mounted () {
    console.log(this.$route)
  }
}
</script>
