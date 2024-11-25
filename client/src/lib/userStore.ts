// src/lib/userStore.ts
import { writable } from 'svelte/store';
import type { User } from '../types';
import { browser } from '$app/environment';

// Function to safely parse JSON
function safeJSONParse(str: string | null) {
    if (!str) return null;
    try {
        return JSON.parse(str);
    } catch (e) {
        console.error('Error parsing JSON from localStorage:', e);
        return null;
    }
}

// Initialize currentUser
const storedUser = browser ? localStorage.getItem('currentUser') : null;
const initialUser: User | null = safeJSONParse(storedUser);

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

// Initialize allUsers
const storedUsers = browser ? localStorage.getItem('allUsers') : null;
const initialUsers: User[] = safeJSONParse(storedUsers) || [];

function createAllUsersStore() {
    const { subscribe, set, update } = writable<User[]>(initialUsers);

    return {
        subscribe,
        set: (value: User[]) => {
            if (browser) {
                localStorage.setItem('allUsers', JSON.stringify(value));
            }
            set(value);
        },
        update: (updater: (value: User[]) => User[]) => {
            update(currentValue => {
                const newValue = updater(currentValue);
                if (browser) {
                    localStorage.setItem('allUsers', JSON.stringify(newValue));
                }
                return newValue;
            });
        }
    };
}

export const allUsers = createAllUsersStore();