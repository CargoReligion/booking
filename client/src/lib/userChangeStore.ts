// src/lib/userChangeStore.ts
import { writable } from 'svelte/store';
import type { User } from '../types'; 

export const userChangeStore = writable<User | null>(null);