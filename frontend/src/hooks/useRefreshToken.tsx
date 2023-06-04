import { refreshTokenAPI } from "../utils/request";
import useAuth from "./useAuth";


const useRefreshToken = () => {
  const {user, } = useAuth();

  const refresh = async () => {
    const res = await refreshTokenAPI()
    setAuth((prev) => {
      return {...prev, user: res.data.uesr, accessToken: res.data.accessToken}})

    return res.data.accessToken
  }
  return refresh
}

export default useRefreshToken
