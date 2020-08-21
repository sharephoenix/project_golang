<template>
  <div class="login-page">
    <div class="login-container">
      <div class="row-cls">
        <span>手机号码：</span><input v-model="mobile" placeholder="请输入邮箱">
      </div>
      <div class="row-cls">
        <span>验证码：</span><input v-model="code" placeholder="请输入验证码">
      </div>
      <div class="login-cls">
        <span @click="loginAction">登陆</span>
        <span @click="sendCode">发送验证码</span>
      </div>
    </div>
  </div>
</template>

<script>

import { getMobileCode } from '@/api/user'

export default {
  name: 'Login',
  data () {
    return {
      mobile: null,
      code: null
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
    loginAction () {
      console.log('[loginAction]')
    },
    sendCode () {
      const isMobile = this.isMobileNumber(this.mobile)
      if (isMobile !== false) {
        alert(isMobile)
        return
      }
      getMobileCode(this.mobile)
    },
    isMobileNumber (phone) {
      if (typeof phone !== 'string') {
        return '请输入手机号码'
      }
      var flag = false
      var message = ''
      var myreg = /^(((13[0-9]{1})|(14[0-9]{1})|(17[0-9]{1})|(15[0-3]{1})|(15[4-9]{1})|(18[0-9]{1})|(199))+\d{8})$/
      if (phone === '') {
        message = '手机号码不能为空！'
      } else if (phone.length !== 11) {
        message = '请输入11位手机号码！'
      } else if (!myreg.test(phone)) {
        message = '请输入有效的手机号码！'
      } else {
        flag = true
      }
      if (message !== '') {
        return message
      }
      return flag
    }
  }
}
</script>

<style scoped lang="scss">
.login-page {
  position: relative;
  background: lightgreen;
  height: 100%;
}

.login-container {
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
    margin-left: 5px;
    margin-right: 5px;
    color: white;
    user-select: none;
  }
}
</style>
