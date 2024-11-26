<script lang="ts">
    import { onMount } from 'svelte';
    import { api } from '$lib/api';
    import type { User } from '../../../types';

    let users: User[] = [];

    onMount(async () => {
      try {
        const response: User[] = await api.getStudentsWithSessions();
        users = response.filter(user => user.role === 'student');
      } catch (error) {
        console.error('Error fetching users with sessions:', error);
      }
    });
</script>

<h1>Students with Past Sessions</h1>

<a href="/coach" class="back-link">Back to Coach Dashboard</a>

<ul>
    {#each users as user (user.id)}
        <li>
            {user.name}
            <a href="/coach/student-sessions/{user.id}" class="view-sessions-btn">View Sessions</a>
        </li>
    {/each}
</ul>

<style>
    .back-link, .view-sessions-btn {
        display: inline-block;
        margin: 10px 0;
        padding: 8px 12px;
        background-color: #f0f0f0;
        color: #333;
        text-decoration: none;
        border-radius: 4px;
        transition: background-color 0.3s ease;
    }

    .back-link:hover, .view-sessions-btn:hover {
        background-color: #e0e0e0;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        margin-bottom: 15px;
        padding: 10px;
        border: 1px solid #ddd;
        border-radius: 4px;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
</style>