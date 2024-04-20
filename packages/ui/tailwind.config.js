module.exports = {
    content: [
      "./src/**/*.{html,js,svelte,ts}",
      "../consumers/**/*.{html,js,svelte,ts}"
    ],
    plugins: [require('daisyui')],
    daisyui: {
      themes: ["lemonade", "valantine", "cupcake", "autumn" ]
    },
  };