"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.getMobileCode = getMobileCode;
exports.fetchUsers = fetchUsers;
exports.register = register;
exports.fetchLogin = fetchLogin;

var _axios = _interopRequireDefault(require("axios"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

var instance = _axios["default"].create({
  baseURL: 'http://127.0.0.1:30009',
  headers: {
    version: '3'
  }
}); // 发送验证码给对应的邮箱


function getMobileCode(mobile, callback) {
  instance.get("/sendCode/".concat(mobile)).then(function (res) {
    console.log(res);
    callback(res.data);
  });
} // 获取所有用户信息


function fetchUsers(callback) {
  instance.get('/getUsers').then(function (res) {
    console.log(res);
    callback(res.data);
  });
} // 注册用户


function register(data, callback) {
  var nickname = data.nickname,
      email = data.email,
      address = data.address,
      avatar = data.avatar,
      mobile = data.mobile,
      age = data.age;
  instance.post('/register', {
    nickname: nickname,
    email: email,
    address: address,
    avatar: avatar,
    mobile: mobile,
    age: age
  }).then(function (res) {
    console.log(res);
    callback(res.data);
  });
} // 登陆并获取用户信息


function fetchLogin(data, callback) {
  var mobile = data.mobile,
      code = data.code;
  instance.post('/login', {
    mobile: mobile,
    code: code
  }).then(function (res) {
    console.log(res);
    localStorage.setItem('loginUserinfo', res.data.data);
    localStorage.setItem('Token', res.data.data.accessToken);
    callback(res.data);
  });
}