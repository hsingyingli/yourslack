import { extendTheme, type ThemeConfig } from '@chakra-ui/react'
import { mode, StyleFunctionProps } from '@chakra-ui/theme-tools'

const config: ThemeConfig = {
  initialColorMode: 'dark',
  useSystemColorMode: false,
}

const styles = {
  global: (props: StyleFunctionProps) => ({
    body: {
      bg: mode('white', 'gray.500')(props),
    },
  }),
}

const theme = extendTheme({ config, styles})

export default theme
