import { extendTheme, type ThemeConfig } from '@chakra-ui/react'
import { mode, StyleFunctionProps } from '@chakra-ui/theme-tools'

const config: ThemeConfig = {
  initialColorMode: 'light',
  useSystemColorMode: false,
}

const styles = {
  global: (props: StyleFunctionProps) => ({
    body: {
      bg: mode('#ffffff40', '#20202380')(props),
    },
  }),
}

const theme = extendTheme({ config, styles})

export default theme
