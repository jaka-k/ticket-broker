/// <reference types="svelte" />
import { SvelteComponentTyped } from 'svelte';

interface ButtonProps {
    text?: string;
    onClick?: () => void;
}

export default class Button extends SvelteComponentTyped<ButtonProps> {}
