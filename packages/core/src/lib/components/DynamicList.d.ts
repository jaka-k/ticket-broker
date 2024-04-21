/// <reference types="svelte" />
import { SvelteComponentTyped } from 'svelte';

interface DynamicListProps {
    items: Array<any>;
}

export default class DynamicList extends SvelteComponentTyped<DynamicListProps> {}
