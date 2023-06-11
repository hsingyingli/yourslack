import { useEffect, useMemo } from "react";
import useAuth from "./useAuth";
import axios, { CreateAxiosDefaults } from "axios";


const useAxiosPrivate = (config: CreateAxiosDefaults ) => {
  const {user, refreshToken} = useAuth();
  const axiosPrivate = useMemo(() => axios.create(config), [config])

  useEffect(()=> {
    const reqIntercept = axiosPrivate.interceptors.request.use(
      config => {
        if (!config.headers['Authorization']) {
          config.headers['Authorization'] = `Bearer ${user?.accessToken}` 
        }
        return config
      }, (error) => Promise.reject(error)
    )

    const resIntercept = axiosPrivate.interceptors.response.use(
      response => response,
      async (error) => {
        const prevRequest = error?.config;
        if (error?.response?.status === 401 && !prevRequest?.sent){
          prevRequest.sent = true
          const newAccessToken = await refreshToken();
          prevRequest.headers[`Authorization`] = `Bearer ${newAccessToken}`
          return axiosPrivate(prevRequest)
        }
        return Promise.reject(error)
      }
    )
    return () => {
      axiosPrivate.interceptors.request.eject(reqIntercept)
      axiosPrivate.interceptors.response.eject(resIntercept)
    }
  }, [user, refreshToken, axiosPrivate])
  return axiosPrivate
}

export default useAxiosPrivate

