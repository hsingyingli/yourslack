import { Box, Spinner } from "@chakra-ui/react";

const LoadingPage = () => {
  return (
    <Box
        display='flex'
        w='100vw' h='100vh' 
        alignItems='center' justifyContent='center' 
      >
        <Spinner/>
      </Box>  
  )
}

export default LoadingPage
