import { Box, Button, FormControl, FormLabel, Heading, Input } from "@chakra-ui/react";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import useAuth from "../../hooks/useAuth";

const SignInPage: React.FC = () => {
  const {loginUser} = useAuth()
  const navigate = useNavigate()
  const [email, setEmail] = useState<string>('')
  const [password, setPassword] = useState<string>('')

  const handleSignIn = async (e: React.FormEvent) => {
    e.preventDefault()
    const isSuccess = await loginUser(email, password)
    if (isSuccess) {
      navigate('/')
    }
  }

  return (
    <Box display='flex' w='100vw' h='calc(100vh - 56px)' 
      alignItems={'center'} justifyContent='center'>
      <Box borderWidth={1} px={12} py={4} rounded="md" w={'md'} m={2}>
        <Heading size='lg' mb={5}>Sign In</Heading>
        <form onSubmit={handleSignIn}>
          <FormControl isRequired>
            <FormLabel>Email</FormLabel>
            <Input type='email' value={email} onChange={(e) => setEmail(e.target.value)} />
          </FormControl>
          <FormControl isRequired>
            <FormLabel>Password</FormLabel>
            <Input type='password' minLength={8} value={password} onChange={(e) => setPassword(e.target.value)} />
          </FormControl>
          <Button type='submit' mt={2}>Signin</Button>
        </form>
      </Box>
    </Box>
  )
}

export default SignInPage
