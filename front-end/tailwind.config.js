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
    container: {
      padding: "2rem",
      center: true
    },
    screens: {
      sm: "640px",
      md: "768px"
    }
  },
  plugins: [],
}
