<template>
  <div class="register-page">
    <div class="register-container">
      <div class="row-cls">
        <span>昵称：</span><input v-model="nickname">
      </div>
      <div class="row-cls">
        <span>手机：</span><input v-model="mobile">
      </div>
      <div class="row-cls">
        <span>邮箱：</span><input v-model="email">
      </div>
      <div class="row-cls">
        <span>地址：</span><input v-model="address">
      </div>
      <div class="row-cls">
        <span>年龄：</span><input v-model="age">
      </div>
      <div class="row-cls">
        <span>头像URL：</span><input v-model="avatar">
      </div>
      <div class="login-cls" @click="registerAction">
        <span>注册</span>
      </div>
    </div>
  </div>
</template>

<script>

import { register } from '@/api/user'

export default {
  name: 'Register',
  data () {
    return {
      nickname: localStorage.getItem('nickname'),
      mobile: localStorage.getItem('mobile'),
      email: localStorage.getItem('email'),
      address: localStorage.getItem('address'),
      age: localStorage.getItem('age'),
      avatar: localStorage.getItem('avatar')
    }
  },
  watch: {
    nickname (newVal, oldVal) {
      console.log(newVal, oldVal)
      localStorage.setItem('nickname', this.nickname)
    },
    mobile (newVal, oldVal) {
      console.log(newVal, oldVal)
      localStorage.setItem('mobile', this.mobile)
    },
    email (newVal, oldVal) {
      console.log(newVal, oldVal)
      localStorage.setItem('email', this.email)
    },
    address (newVal, oldVal) {
      console.log(newVal, oldVal)
      localStorage.setItem('address', this.address)
    },
    age (newVal, oldVal) {
      console.log(newVal, oldVal)
      localStorage.setItem('age', this.age)
    },
    avatar (newVal, oldVal) {
      console.log(newVal, oldVal)
      localStorage.setItem('avatar', this.avatar)
    }
  },
  methods: {
    registerAction () {
      const registData = {
        nickname: this.nickname,
        email: this.email,
        address: this.address,
        avatar: this.avatar,
        mobile: this.mobile,
        age: Number(this.age)
      }
      console.log(registData)
      register(registData)
      localStorage.clear()
    }
  }
}
</script>

<style scoped lang="scss">
.register-page {
  position: relative;
  background: lightgreen;
  height: 100%;
}

.register-container {
  position: absolute;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: space-between;
  background: lightgrey;
  margin: auto auto;
  width: 300px;
  height: 300px;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 5px;
}

.row-cls {
  margin-top: 10px;

  span {
    display: inline-block;
    width: 80px;
    text-align: right;
  }
}

.login-cls {
  margin-top: 22px;
  span {
    display: inline-block;
    width: 133px;
    height: 33px;
    line-height: 33px;
    background-color: green;
    border-radius: 8px;
  }
}
</style>
