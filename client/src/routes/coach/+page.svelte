<script lang="ts">
    import { onMount } from 'svelte';
    import { api } from '$lib/api';
    import type { SlotData, CreateSlotData, Paginated } from '../../types';
    import { currentUser } from '$lib/userStore';

    let upcomingSlots: Paginated<SlotData> = {
        data: [],
        page: 1,
        pageSize: 10,
        totalCount: 0,
        totalPages: 0
    };
    let currentPage = 1;
    let selectedDate: string = new Date().toISOString().split('T')[0];
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

    async function refreshSlots(page: number = 1) {
        try {
            console.log('Refreshing slots for coach:', currentCoachId);
            upcomingSlots = await api.getUpcomingSlots(page, 10);
            currentPage = page;
        } catch (error) {
            console.error('Error fetching upcoming slots:', error);
        }
    }

    function changePage(newPage: number) {
        if (newPage >= 1 && newPage <= upcomingSlots.totalPages) {
            refreshSlots(newPage);
        }
    }

    function localToUTC(dateString: string, timeString: string): string {
        const localDate = new Date(`${dateString}T${timeString}`);
        return localDate.toISOString();
    }

    async function createSlot() {
        if (!currentCoachId) {
            alert('Please select a coach to impersonate first.');
            return;
        }

        const utcStartTime = localToUTC(selectedDate, selectedTime);
        const slotData: CreateSlotData = {
            startTime: utcStartTime
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
<button on:click={createSlot}>Create Slot</button>

<h2>Upcoming Slots:</h2>
{#if upcomingSlots.data.length > 0}
    <ul>
        {#each upcomingSlots.data as slot}
            <li>{formatDate(slot.startTime)} - {slot.booked ? 'Booked' : 'Available'}</li>
        {/each}
    </ul>
    
    <!-- Pagination UI -->
    <div class="pagination">
        <button on:click={() => changePage(currentPage - 1)} disabled={currentPage === 1}>
            &lt; Previous
        </button>
        
        {#each Array(upcomingSlots.totalPages) as _, i}
            <button 
                on:click={() => changePage(i + 1)} 
                class:active={currentPage === i + 1}
            >
                {i + 1}
            </button>
        {/each}
        
        <button on:click={() => changePage(currentPage + 1)} disabled={currentPage === upcomingSlots.totalPages}>
            Next &gt;
        </button>
    </div>
{:else}
    <p>You have no upcoming slots.</p>
{/if}

<nav>
    <a href="/coach/past-sessions">View Past Sessions</a>
</nav>

<style>
    .pagination {
        display: flex;
        justify-content: center;
        margin-top: 1rem;
    }

    .pagination button {
        margin: 0 0.25rem;
        padding: 0.5rem 1rem;
        border: 1px solid #ccc;
        background-color: #fff;
        cursor: pointer;
    }

    .pagination button.active {
        background-color: #007bff;
        color: #fff;
    }

    .pagination button:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }
</style>