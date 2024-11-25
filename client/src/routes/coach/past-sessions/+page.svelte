<script>
    import { onMount } from 'svelte';
    import { api } from '$lib/api';
  
    let pastSessions = [];
  
    onMount(async () => {
      const response = await api.getPastSessionFeedbacks();
      pastSessions = response.data;
    });
  </script>
  
  <h1>Past Sessions</h1>
  
  <ul>
    {#each pastSessions as session}
      <li>
        Date: {session.date}
        Student: {session.student_name}
        Satisfaction: {session.satisfaction}
        Notes: {session.notes}
      </li>
    {/each}
  </ul>