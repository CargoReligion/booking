<script lang="ts">
    import { onMount } from 'svelte';
    import { api } from '$lib/api';
    import type { SlotData, CreateSlotData, ApiResponse } from '../../types';
    import { currentUser } from '$lib/userStore';

    let upcomingSlots: SlotData[] = [];
    let selectedDate: string = new Date().toISOString().split('T')[0]; // Today's date
    let selectedTime: string = '09:00';
    
    let currentCoachId: string | null = null;

    currentUser.subscribe(user => {
        currentCoachId = user?.id || null;
    });

    onMount(async () => {
        if (currentCoachId) {
            await refreshSlots();
        } else {
            console.log('No coach selected. Please use the impersonate dropdown to select a coach.');
        }
    });


    async function refreshSlots() {
        try {
            const response: SlotData[] = await api.getUpcomingSlots();
            upcomingSlots = response;
        } catch (error) {
            console.error('Error fetching upcoming slots:', error);
        }
    }

    function formatDate(dateString: string): string {
        const date = new Date(dateString);
        return date.toLocaleString(undefined, {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: '2-digit',
            minute: '2-digit',
            timeZoneName: 'short'
        });
    }

    async function createSlot() {
        if (!currentCoachId) {
            alert('Please select a coach to impersonate first.');
            return;
        }

        const startTime = `${selectedDate}T${selectedTime}:00Z`;
        const slotData: CreateSlotData = {
            startTime: startTime,
        };

        try {
            await api.createSlot(slotData);
            alert('Slot created successfully!');
            await refreshSlots();
            // Reset time selection
            selectedTime = '09:00';
        } catch (error) {
            console.error('Error creating slot:', error);
            alert('Failed to create slot. Please try again.');
        }
    }

    const timeSlots = Array.from({ length: 33 }, (_, i) => {
        const hours = Math.floor(i / 4) + 9;
        const minutes = (i % 4) * 15;
        return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}`;
    });
</script>

<h1>Coach Dashboard</h1>

{#if $currentUser}
    <p>Current Coach: {$currentUser.name}</p>
{:else}
    <p>No coach selected. Please use the impersonate dropdown to select a coach.</p>
{/if}

<h2>Create New Slot</h2>
{#if $currentUser}
    <p>Current Coach: {$currentUser.name}</p>
{:else}
    <p>No coach selected. Please use the impersonate dropdown to select a coach.</p>
{/if}
<div>
    <label>
        Date:
        <input type="date" bind:value={selectedDate} min={new Date().toISOString().split('T')[0]}>
    </label>
</div>
<div>
    <label>
        Time:
        <select bind:value={selectedTime}>
            {#each timeSlots as time}
                <option value={time}>{time}</option>
            {/each}
        </select>
    </label>
</div>
<button on:click={createSlot} disabled={!currentCoachId}>Create Slot</button>

<h2>Upcoming Slots:</h2>
{#if upcomingSlots.length > 0}
    <ul>
        {#each upcomingSlots as slot}
            <li>{formatDate(slot.startTime)} - {slot.booked ? 'Booked' : 'Available'}</li>
        {/each}
    </ul>
{:else}
    <p>You have no upcoming slots.</p>
{/if}
<nav>
    <a href="/coach/past-sessions">View Past Sessions</a>
</nav>