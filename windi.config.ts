import { defineConfig } from 'windicss/helpers';
import colors from 'windicss/colors';
import typography from 'windicss/plugin/typography';

export default defineConfig({
  darkMode: 'class',
  // https://windicss.org/posts/v30.html#attributify-mode
  attributify: true,

  plugins: [typography()],
  theme: {
    extend: {
      colors: {
        neonGreen: '#9bf73c',
        neonGreenShade: '#f2ffe6',
        textDark: '#303030',
      },
      fontFamily: {
        sans: ['Noto Sans SC', 'sans-serif'],
        bahnschrift: ['Bahnschrift', 'Noto Sans SC', 'sans-serif'],
        jlinxin: ['JLinXin', 'Noto Sans SC', 'sans-serif'],
      },
      typography: {
        DEFAULT: {
          css: {
            maxWidth: '65ch',
            color: 'inherit',
            a: {
              color: 'inherit',
              opacity: 0.75,
              fontWeight: '500',
              textDecoration: 'underline',
              '&:hover': {
                opacity: 1,
                color: colors.teal[600],
              },
            },
            b: { color: 'inherit' },
            strong: { color: 'inherit' },
            em: { color: 'inherit' },
            h1: { color: 'inherit' },
            h2: { color: 'inherit' },
            h3: { color: 'inherit' },
            h4: { color: 'inherit' },
            code: { color: 'inherit' },
          },
        },
      },
    },
  },
});
