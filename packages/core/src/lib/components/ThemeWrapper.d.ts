/// <reference types="svelte" />
import { SvelteComponentTyped } from 'svelte';

interface ThemeWrapperProps {
    theme: string;
}

export default class ThemeWrapper extends SvelteComponentTyped<ThemeWrapperProps> {}
