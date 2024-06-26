/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./static/*.{html, js}', './static/**/*.{html,js}'],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
  daisyui: {
    themes: ["dark", "coffee", "cyberpunk"],
  }
}

