<script lang="ts">
    import { onMount} from 'svelte';
    import { api, setUserID } from '$lib/api';
    import type { User } from '../types';
    import { goto } from '$app/navigation';
    import { currentUser } from '$lib/userStore';
    import { userChangeStore } from '$lib/userChangeStore';
    import { browser } from '$app/environment';
    
    let users: User[] = [];
    let selectedUser: User | null = null;
  
    onMount(async () => {
      try {
        const response: User[] = await api.getAllUsers();
        users = response;
        
        if (browser && $currentUser) {
          // Find the persisted user in the fetched users array
          selectedUser = users.find(user => user.id === $currentUser.id) || null;
          if (selectedUser) {
            setUserID(selectedUser.id);
          }
        }
      } catch (error) {
        console.error('Error fetching users:', error);
      }
    });
  
    function handleUserSelect() {
      if (selectedUser) {
        setUserID(selectedUser.id);
        currentUser.set(selectedUser);
        userChangeStore.set(selectedUser);
        if (selectedUser.role === 'coach') {
          goto('/coach');
        } else {
          goto('/student');
        }
      } else {
        currentUser.set(null);
        userChangeStore.set(null);
        setUserID('');
      }
    }

</script>
<div style="display: flex; justify-content: flex-end; padding: 1rem;">
  <select bind:value={selectedUser} on:change={handleUserSelect}>
    <option value={null}>Impersonate</option>
    {#each users as user (user.id)}
      <option value={user}>{user.name}: {user.role}</option>
    {/each}
  </select>
</div>

<slot/>