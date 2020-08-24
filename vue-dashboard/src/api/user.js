import axios from 'axios'

const instance = axios.create({
  baseURL: 'http://127.0.0.1:30009',
  headers: { version: '3' }
})

// 发送验证码给对应的邮箱
export function getMobileCode (mobile, callback) {
  instance.get(`/sendCode/${mobile}`).then(res => {
    console.log(res)
    callback(res.data)
  })
}

// 获取所有用户信息
export function fetchUsers (callback) {
  instance.get('/getUsers').then(res => {
    console.log(res)
    callback(res.data)
  })
}

// 注册用户
export function register (data, callback) {
  const { nickname, email, address, avatar, mobile, age } = data
  instance.post('/register', { nickname, email, address, avatar, mobile, age }).then(res => {
    console.log(res)
    callback(res.data)
  })
}

// 登陆并获取用户信息
export function fetchLogin (data, callback) {
  const { mobile, code } = data
  instance.post('/login', { mobile, code }).then(res => {
    console.log(res)
    localStorage.setItem('loginUserinfo', res.data.data)
    localStorage.setItem('Token', res.data.data.accessToken)
    callback(res.data)
  })
}
