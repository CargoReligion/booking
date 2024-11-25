<script lang="ts">
    import { api } from '$lib/api';
    import type { CreateSlotData } from '../../../types';
  
    let startTime = '';
    let currentCoachId: number = 1; // This should be set to the actual logged-in coach's ID
  
    async function handleSubmit() {
      const slotData: CreateSlotData = {
        startTime: startTime,
      };
  
      try {
        await api.createSlot(slotData);
        alert('Slot created successfully!');
        startTime = '';
      } catch (error) {
        console.error('Error creating slot:', error);
        alert('Failed to create slot. Please try again.');
      }
    }
  </script>
  
  <h1>Create New Slot</h1>
  <form on:submit|preventDefault={handleSubmit}>
    <label>
      Start Time:
      <input type="datetime-local" bind:value={startTime} required>
    </label>
    <button type="submit">Create Slot</button>
  </form>