/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./public/**/*.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        // Warm-paper editorial palette
        paper: {
          DEFAULT: '#F5EFE3',
          deep: '#EDE5D2',
          edge: '#E2D9C2',
        },
        ink: {
          DEFAULT: '#1B1B17',
          soft: '#3A372E',
          muted: '#76705F',
          faint: '#A8A28E',
        },
        oxblood: {
          DEFAULT: '#7A1B1B',
          dark: '#5C1414',
          light: '#9A2C2C',
        },
        ochre: '#A88532',
      },
      fontFamily: {
        display: ['Fraunces', 'Georgia', 'serif'],
        serif: ['"EB Garamond"', 'Georgia', 'serif'],
        mono: ['"IBM Plex Mono"', 'ui-monospace', 'monospace'],
      },
      letterSpacing: {
        widest: '0.18em',
        masthead: '0.04em',
      },
      maxWidth: {
        prose: '68ch',
        masthead: '1280px',
      },
      keyframes: {
        rise: {
          '0%': { opacity: '0', transform: 'translateY(8px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        drawLine: {
          '0%': { transform: 'scaleX(0)' },
          '100%': { transform: 'scaleX(1)' },
        },
      },
      animation: {
        rise: 'rise 0.5s cubic-bezier(0.2, 0.7, 0.2, 1) both',
        'draw-line': 'drawLine 0.6s cubic-bezier(0.2, 0.7, 0.2, 1) both',
      },
    },
    screens: {
      sm: '480px',
      md: '768px',
      lg: '1024px',
      xl: '1280px',
    },
  },
  plugins: [],
};
