<template>
  <div class="home">
    <table>
      <div class="header-cls">
        <span>id</span>
        <span>avatar</span>
        <span>nickname</span>
        <span>age</span>
        <span>mobile</span>
        <span>email</span>
        <span>address</span>
      </div>
      <div class="tb-row-cls"
           v-for="(item, index) in userList"
           :key="index">
        <span>{{item.id}}</span>
        <span>{{item.avatar}}</span>
        <span>{{item.nickname}}</span>
        <span>{{item.age}}</span>
        <span>{{item.mobile}}</span>
        <span>{{item.email}}</span>
        <span>{{item.addsress}}</span>
      </div>
    </table>
    <div @click="testAction"> count: {{ getCount }} </div>
  </div>
</template>

<script>
import { fetchUsers } from '@/api/user'

export default {
  name: 'Home',
  data () {
    return {
      userList: [
        {
          nickname: 'alexluan',
          age: 12,
          address: 'shanghai',
          email: 'xfxxxx@qq.com',
          avatar: 'xxxx',
          mobile: '110',
          id: 'userid-0'
        },
        {
          nickname: 'alexluan',
          age: 12,
          address: 'shanghai',
          email: 'xfxxxx@qq.com',
          avatar: 'xxxx',
          mobile: '110',
          id: 'userid-1'
        }
      ]
    }
  },
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
    fetchUsers((res) => {
      const { code, data, msg } = res
      if (code === 0) {
        this.userList = data
      } else {
        alert(msg)
      }
    })
  }
}
</script>

<style scoped lang="scss">
.header-cls {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  span {
    display: inline-block;
    background-color: gray;
    color: white;
    width: 140px;
    height: 30px;
    line-height: 30px;
    padding-left: 8px;
    padding-right: 8px;
  }
}

.tb-row-cls {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  span {
    display: inline-block;
    width: 140px;
    height: 30px;
    line-height: 30px;
    padding-left: 8px;
    padding-right: 8px;
    text-align: center;
  }
}
</style>
