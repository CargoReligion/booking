// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

/// <reference types="svelte" />

declare namespace svelteHTML {
    interface HTMLProps<T> {
        'on:userChanged'?: (event: CustomEvent<User | null>) => void;
    }
}

export {};
