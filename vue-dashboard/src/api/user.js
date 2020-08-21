import axios from 'axios'

export function getMobileCode () {
  const instance = axios.create({
    baseURL: 'http://localhost:9090'
  })
  instance.get('/getCode/18817322818').then(res => {
    console.log(res)
  })
}

export function register (data) {
  const instance = axios.create({
    headers: { version: '3' },
    baseURL: 'http://localhost:9090'
  })
  const { nickname, email, address, avatar, mobile, age } = data
  instance.post('/register', { nickname, email, address, avatar, mobile, age }).then(res => {
    console.log(res)
  })
}
