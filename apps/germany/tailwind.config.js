// In each app's tailwind.config.js
const { resolve } = require('path');
const sharedConfig = require(resolve(__dirname, '../../tailwind.config.js'));

module.exports = {
  ...sharedConfig,
  content: [
    './src/**/*.{html,js,svelte,ts}',
    resolve(__dirname, '../../packages/core/src/**/*.{html,js,svelte,ts}'),
  ],

  daisyui: {
    themes: ['aqua'],
  },
};
