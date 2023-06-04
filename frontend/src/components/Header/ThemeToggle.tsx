import { IconButton, useColorMode, useColorModeValue } from "@chakra-ui/react";
import React from "react";
import { MoonIcon, SunIcon } from '@chakra-ui/icons'

const ThemeToggle: React.FC = () => {
  const Icon = useColorModeValue(SunIcon, MoonIcon)
   const { toggleColorMode } = useColorMode()
  return (
    <IconButton 
      colorScheme={useColorModeValue('yellow', 'purple')}
      color={useColorModeValue('black', 'white')}
      aria-label='theme toggle button' 
      icon={<Icon />} 
      onClick={toggleColorMode}
    />
  )
}

export default ThemeToggle
