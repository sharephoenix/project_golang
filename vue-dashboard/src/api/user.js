import axios from 'axios'

const instance = axios.create({
  baseURL: 'http://localhost:9090',
  headers: { version: '3' }
})

// 发送验证码给对应的邮箱
export function getMobileCode (mobile) {
  instance.get(`/sendCode/${mobile}`).then(res => {
    console.log(res)
  })
}

// 注册用户
export function register (data) {
  const { nickname, email, address, avatar, mobile, age } = data
  instance.post('/register', { nickname, email, address, avatar, mobile, age }).then(res => {
    console.log(res)
  })
}

export function fetchLogin (data) {
  const { mobile, code } = data
  instance.post('/login', { mobile, code }).then(res => {
    console.log(res)
  })
}
