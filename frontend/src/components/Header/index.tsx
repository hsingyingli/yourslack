import { Box, Flex, Spacer, Text, useColorModeValue } from "@chakra-ui/react";
import ThemeToggle from "./ThemeToggle";
import React from "react";


const Header: React.FC = () => {
  return (
    <Flex 
      as="header"  
      maxW={"100vw"}
      w={"100%"}
      paddingX={3} paddingY={2} 
      pos='fixed' top={0} left={0} 
      align="center"
      bg={useColorModeValue('#ffffff40', '#20202380')}
      css={{ backdropFilter: 'blur(10px)' }}

    >
      <Text>YourSlack</Text>
      <Spacer/>
      <ThemeToggle/>
    </Flex>
  )
}

export default Header
