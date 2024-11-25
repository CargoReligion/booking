// src/lib/userStore.ts
import { writable } from 'svelte/store';
import type { User } from '../types';
import { browser } from '$app/environment';

let initialUser: User | null = null;

if (browser) {
  const storedUser = localStorage.getItem('currentUser');
  initialUser = storedUser ? JSON.parse(storedUser) : null;
}

export const currentUser = writable<User | null>(initialUser);

if (browser) {
  currentUser.subscribe(value => {
    if (value) {
      localStorage.setItem('currentUser', JSON.stringify(value));
    } else {
      localStorage.removeItem('currentUser');
    }
  });
}