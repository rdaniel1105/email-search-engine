/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./public/**/*.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        "cool-blacky": "#0A001C",
        "nav-color": "#e8c39e"
      }
    },
    fontFamily: {
      Poppins: "Poppins, sans-serif"
    },
    screens: {
      sm: "100px",
      md: "768px",
      lg: "992px"
    }
  },
  plugins: [],
}
