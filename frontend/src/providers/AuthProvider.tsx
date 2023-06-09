import React, { useState, createContext, useEffect } from "react";
import { useToast } from '@chakra-ui/react'
import {User}  from "../utils/types"
import {
  loginUserAPI, 
  logoutUserAPI, 
  refreshTokenAPI,
  signUpUserAPI,
} from "../utils/request"
import LoadingPage from "../components/LoadingPage";

interface AuthContextInterface {
  user: User | null
  loginUser: (email: string, password: string) => Promise<boolean>
  logoutUser: () => Promise<void>
  refreshToken: () => Promise<String>
  signUpUser: (username: string, email: string, password: string) => Promise<boolean>
}

const initContext = {
  user: null,
  loginUser: async () => false,
  logoutUser: async () => {},
  refreshToken: async () => '',
  signUpUser: async () => false,
}

const AuthContext = createContext<AuthContextInterface>(initContext)

interface AuthProviderInterface {
  children: React.ReactNode
}

const AuthProvider: React.FC<AuthProviderInterface> = ({children}) => {
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const toast = useToast() 

  const signUpUser = async (username: string, email: string, password: string) => {
    try {
      await signUpUserAPI(username, email, password)
      toast({
        title: `Create Account`,
        status: 'success',
        duration: 1000,
        isClosable: true, 
        position: 'top',
      })
      return true
    } catch (error) {
      console.log(error)
      toast({
        title: `Create Account Fail!!`,
        status: 'error',
        duration: 1000,
        isClosable: true, 
        position: 'top',
      })
      return false
    }
  }

  const loginUser =  async (email: string, password: string) => {
    try {
      const res = await loginUserAPI(email, password)
      setUser({
        id: res.data.id,
        accessToken: res.data.access_token,
        username: res.data.username,
        email: res.data.email,
        expiredAt: res.data.expired_at
      })
      toast({
        title: `Hello, ${res.data.username}`,
        status: 'success',
        duration: 1000,
        isClosable: true, 
        position: 'top',
      })
      return true
    } catch (error) {
      toast({
        title: `Fail to login`,
        status: 'error',
        duration: 1000,
        isClosable: true, 
        position: 'top',
      })
      return false
    }
  }

  const refreshToken = async () => {
    try {
      const res = await refreshTokenAPI()
      setUser({
        id: res.data.id,
        accessToken: res.data.access_token,
        username: res.data.username,
        email: res.data.email,
        expiredAt: res.data.expired_at
      })
      return res.data
    } catch (error) {
      setUser(null)
      return ''
    }
  }

  const logoutUser = async () => {
    try {
      await logoutUserAPI()
      toast({
        title: `Byebye, ${user?.username}`,
        status: 'success',
        duration: 1000,
        isClosable: true, 
      })
      setUser(null)
    } catch (error) {
      toast({
        title: `Fail to logout`,
        status: 'error',
        duration: 1000,
        isClosable: true, 
      })
    }
  }

  useEffect(() => {
    setIsLoading(true) 
    const init = async () => {
      await refreshToken()
      setIsLoading(false)
    }
    init()
  }, [])

  return isLoading 
    ? 
      <LoadingPage/>
    : (
      <AuthContext.Provider value={{user, loginUser, logoutUser, refreshToken, signUpUser }}>
        {children}
      </AuthContext.Provider>
    )
}

export default AuthProvider
export {
  AuthContext
}
