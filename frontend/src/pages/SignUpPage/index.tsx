import { Box, Button, FormControl, FormLabel, Input, FormErrorMessage, Heading } from "@chakra-ui/react";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import useAuth from "../../hooks/useAuth";

const SignUpPage: React.FC = () => {
  const {signUpUser} = useAuth()
  const navigate = useNavigate()
  const [username, setUsername] = useState<string>('')
  const [email, setEmail] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [confirmPassword, setConfirmPassword] = useState<string>('')
  const isError = confirmPassword !== '' && password !== '' && confirmPassword !== password

  const handleSignUp = async (e: React.FormEvent) => {
    e.preventDefault()
    if (isError) return

    const isSuccess = await signUpUser(username, email, password)
    console.log(isSuccess)
    if (isSuccess) {
      navigate('/signin')
    }
  }

  return (
    <Box display='flex' w='100vw' h='calc(100vh - 56px)' 
      alignItems={'center'} justifyContent='center'>
      <Box borderWidth={1} px={12} py={4} rounded="md" w={'md'} m={2}>
        <Heading size='lg' mb={5}>Sign Up</Heading>
        <form onSubmit={handleSignUp}>
          <FormControl isRequired>
            <FormLabel>Username</FormLabel>
            <Input type='text' value={username} onChange={(e) => setUsername(e.target.value)} />
          </FormControl>
          <FormControl isRequired>
            <FormLabel>Email</FormLabel>
            <Input type='email' value={email} onChange={(e) => setEmail(e.target.value)} />
          </FormControl>
          <FormControl isRequired>
            <FormLabel>Password</FormLabel>
            <Input type='password' minLength={8} value={password} onChange={(e) => setPassword(e.target.value)} />
          </FormControl>
          <FormControl isRequired isInvalid={isError}>
            <FormLabel>Confirm Password</FormLabel>
            <Input type='password' minLength={8} value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} />
            {
              isError && <FormErrorMessage>Confirm Password must be same as Password.</FormErrorMessage>
            }
          </FormControl>
          <Button type='submit' mt={2} disabled={isError}>Signin</Button>
        </form>
      </Box>
    </Box>
  )
}

export default SignUpPage
