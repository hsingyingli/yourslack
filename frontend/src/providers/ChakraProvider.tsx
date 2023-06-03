import React from 'react';
import { ChakraProvider as Chakra } from '@chakra-ui/react'
import { ColorModeScript } from '@chakra-ui/react'
import theme from '../utils/theme';

interface Props {
  children: React.ReactNode
}

const ChakraProvider: React.FC<Props> = ({children}) => {
  return (
    <Chakra theme={theme}>
      <ColorModeScript initialColorMode={theme.config.initialColorMode} />
      {children}
    </Chakra>
  )
}

export default ChakraProvider
