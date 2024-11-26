<!-- src/routes/coach/past-sessions/+page.svelte -->
<script lang="ts">
    import { onMount } from 'svelte';
    import { api } from '$lib/api';
    import type { SessionFeedback, ApiResponse } from '../../../types';
    import { formatDate } from '$lib/utils';

    let pastSessions: SessionFeedback[] = [];

    // onMount(async () => {
    //   try {
    //     const response: ApiResponse<SessionFeedback[]> = await api.getPastSessionFeedbacks();
    //     pastSessions = response.data;
    //   } catch (error) {
    //     console.error('Error fetching past sessions:', error);
    //   }
    // });
</script>

<h1>Past Sessions</h1>

<a href="/coach" class="back-link">Back to Coach Dashboard</a>

<ul>
    {#each pastSessions as session (session.id)}
        <li>
            Date: {formatDate(session.slotId)}
            Satisfaction: {session.satisfaction}
            Notes: {session.notes}
        </li>
    {/each}
</ul>

<style>
    .back-link {
        display: inline-block;
        margin-bottom: 20px;
        padding: 10px 15px;
        background-color: #f0f0f0;
        color: #333;
        text-decoration: none;
        border-radius: 5px;
        transition: background-color 0.3s ease;
    }

    .back-link:hover {
        background-color: #e0e0e0;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        margin-bottom: 20px;
        padding: 15px;
        border: 1px solid #ddd;
        border-radius: 5px;
    }
</style>