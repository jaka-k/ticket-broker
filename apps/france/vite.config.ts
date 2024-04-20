import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
	plugins: [sveltekit()],
	/* resolve: {
		alias: {
			'@ticket-broker/ui': path.resolve(__dirname, './packages/ui/src'),
			'@ticket-broker/ui/styles': path.resolve(__dirname, './packages/ui/src/lib/app.css')
		  }
	} */
  });