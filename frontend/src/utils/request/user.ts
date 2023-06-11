import axios from "axios";

const axiosBase = axios.create({
  baseURL: "http://localhost:9010",
})

const refreshTokenAPI = () => {
  return axiosBase.post('/v1/user/renew_access', {}, {
    withCredentials: true
  })
}

const signUpUserAPI = (username: string, email: string, password: string) => {
  return axiosBase.post('/v1/user', {
    username, email, password,
  }, {
    headers: { 'Content-Type': 'application/json' },
    withCredentials: true
    })
}

const loginUserAPI = (email: string, password: string) => {
  return axiosBase.post('/v1/user/login', {
    email,
    password
  }, {
    headers: { 'Content-Type': 'application/json' },
    withCredentials: true
    })
}
const logoutUserAPI = async () => {
  return axiosBase.post('/v1/user/logout', {
    withCredentials: true
  })
}

export {
  refreshTokenAPI,
  signUpUserAPI,
  loginUserAPI,
  logoutUserAPI
}
