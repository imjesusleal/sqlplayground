/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./static/*.html", "./static/js/*.js"],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
  daisyui: {
    themes: ["dark", "cupcake", "cyberpunk"],
  }
}

