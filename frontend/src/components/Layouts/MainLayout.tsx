import { Box } from "@chakra-ui/react";
import React from "react";
import Header from "../Header";

interface Props {
  children: React.ReactNode
}

const MainLayout: React.FC<Props> = ({children}) => {
  return (
    <Box >
      <Header/>
      <Box 
        as='main' mt={"56px"}
        width={"100vw"} minHeight={"calc(100vh - 56px)"}
      >
      {children} 
      </Box>
    </Box>
  )
}

export default MainLayout
