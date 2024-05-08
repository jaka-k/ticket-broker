/// <reference types="svelte" />
import { SvelteComponentTyped } from 'svelte';

interface ButtonsProps {
    countryCode?: string;
    label?: string;
    
}

export default class Buttons extends SvelteComponentTyped<ButtonProps> {}
