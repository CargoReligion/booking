<!-- src/lib/SlotDetails.svelte -->
<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import { api } from '$lib/api';
    import type { SlotDetails } from '../types';
    import { formatDate } from '$lib/utils';
  
    export let slotId: string;
    export let userRole: 'coach' | 'student';
  
    let details: SlotDetails | null = null;
    let loading = true;
    let error: string | null = null;
  
    const dispatch = createEventDispatcher();
  
    async function fetchDetails() {
      try {
        loading = true;
        error = null;
        details = await api.getSlotDetails(slotId);
      } catch (err) {
        error = 'Failed to load slot details.';
        console.error(err);
      } finally {
        loading = false;
      }
    }
  
    fetchDetails();
  
    function close() {
      dispatch('close');
    }
  </script>
  
  {#if loading}
    <p>Loading details...</p>
  {:else if error}
    <p>{error}</p>
  {:else if details}
    <div>
      <h2>Slot Details</h2>
      <p>Date: {formatDate(details.startTime)}</p>
      <p>
        {#if userRole === 'coach'}
          Student: {details.studentName}
          <br>
          Phone: {details.studentPhoneNumber}
        {:else}
          Coach: {details.coachName}
          <br>
          Phone: {details.coachPhoneNumber}
        {/if}
      </p>
      <button on:click={close}>Close</button>
    </div>
  {/if}
  
  <style>
    div {
      padding: 1em;
      border: 1px solid #ccc;
      border-radius: 4px;
    }
  </style>