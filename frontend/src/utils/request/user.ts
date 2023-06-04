import axios from "axios";

const axiosBase = axios.create({
  baseURL: "https://localhost",
})

const refreshTokenAPI = async () => {
  return axiosBase.get('/refresh', {
        withCredentials: true
  })
}

export {
  refreshTokenAPI
}
