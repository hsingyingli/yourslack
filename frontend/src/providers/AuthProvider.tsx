import React, { useState, createContext } from "react";
import {User}  from "../utils/types"


interface AuthContextInterface {
  user: User | null
  loginUser: () => void
}

const initContext = {
  user: null,
  loginUser: () => {}
}

const AuthContext = createContext<AuthContextInterface>(initContext)

interface AuthProviderInterface {
  children: React.ReactNode
}

const AuthProvider: React.FC<AuthProviderInterface> = ({children}) => {
  const [user, setUser] = useState<User | null>(null)
  
  const loginUser = () => {}


  return (
    <AuthContext.Provider value={{user, loginUser}}>
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider
export {
  AuthContext
}
