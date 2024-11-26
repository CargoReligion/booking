<!-- src/routes/coach/+page.svelte -->
<script lang="ts">
    import { onMount } from 'svelte';
    import { fade } from 'svelte/transition';
    import { api } from '$lib/api';
    import type { SlotData, CreateSlotData, ApiResponse, Paginated } from '../../types';
    import { currentUser } from '$lib/userStore';
    import { formatDate, localToUTC } from '$lib/utils';
    import SlotDetails from '$lib/SlotDetails.svelte';
    import CoachFeedback from '$lib/CoachFeedback.svelte';

    let upcomingSlots: Paginated<SlotData> = {
        data: [],
        page: 1,
        pageSize: 10,
        totalCount: 0,
        totalPages: 0
    };
    let selectedDate: string = new Date().toISOString().split('T')[0];
    let selectedTime: string = '09:00';
    let currentCoachId: string | null = null;
    let selectedSlotId: string | null = null;
    let feedbackSlotId: string | null = null;
    let currentPage = 1;

    currentUser.subscribe(user => {
        currentCoachId = user?.id || null;
    });

    onMount(async () => {
        if (currentCoachId) {
            await refreshSlots(1);
        }
    });

    async function refreshSlots(page: number) {
        try {
            const response: Paginated<SlotData> = await api.getUpcomingSlots(page);
            upcomingSlots = response;
            currentPage = page;
        } catch (error) {
            console.error('Error fetching upcoming slots:', error);
        }
    }

    async function createSlot() {
        if (!currentCoachId) {
            alert('Please select a coach to impersonate first.');
            return;
        }

        const localDateTime = new Date(`${selectedDate}T${selectedTime}`);
        const localHours = localDateTime.getHours();
        
        if (localHours < 9 || localHours >= 17) {
            alert('Please select a time between 9 AM and 5 PM.');
            return;
        }

        const utcDateTime = localToUTC(localDateTime);

        const slotData: CreateSlotData = {
            startTime: localDateTime.toISOString()
        };

        try {
            await api.createSlot(slotData);
            alert('Slot created successfully!');
            await refreshSlots(1);
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

    function viewSlotDetails(slotId: string) {
        selectedSlotId = slotId;
    }

    function openFeedbackModal(slotId: string) {
        feedbackSlotId = slotId;
    }

    function closeSlotDetails() {
        selectedSlotId = null;
    }

    function closeFeedbackModal() {
        feedbackSlotId = null;
    }

    function handleFeedbackSubmitted() {
        refreshSlots(currentPage);
        closeFeedbackModal();
    }

    function changePage(newPage: number) {
        if (newPage >= 1 && newPage <= upcomingSlots.totalPages) {
            refreshSlots(newPage);
        }
    }
</script>

<svelte:head>
    <title>Coach Dashboard</title>
</svelte:head>

<div class="dashboard">
    <header>
        <h1>Coach Dashboard</h1>
        {#if $currentUser}
            <p>Welcome, <span class="coach-name">{$currentUser.name}</span>!</p>
        {:else}
            <p class="warning">No coach selected. Please use the impersonate dropdown to select a coach.</p>
        {/if}
    </header>

    <div class="dashboard-grid">
        <section class="card create-slot">
            <h2>Create New Slot</h2>
            <form on:submit|preventDefault={createSlot}>
                <div class="form-group">
                    <label for="date">Date:</label>
                    <input type="date" id="date" bind:value={selectedDate} min={new Date().toISOString().split('T')[0]}>
                </div>
                <div class="form-group">
                    <label for="time">Time:</label>
                    <select id="time" bind:value={selectedTime}>
                        {#each timeSlots as time}
                            <option value={time}>{time}</option>
                        {/each}
                    </select>
                </div>
                <button type="submit" class="btn-primary">Create Slot</button>
            </form>
        </section>

        <section class="card upcoming-slots">
            <h2>Upcoming Slots</h2>
            {#if upcomingSlots.data.length > 0}
                <ul class="slot-list">
                    {#each upcomingSlots.data as slot (slot.id)}
                        <li class="slot-item" transition:fade="{{ duration: 300 }}">
                            <div class="slot-info">
                                <span class="slot-date">{formatDate(new Date(slot.startTime))}</span>
                                <span class="slot-status {slot.booked ? 'booked' : 'available'}">
                                    {slot.booked ? 'Booked' : 'Available'}
                                </span>
                            </div>
                            <div class="slot-actions">
                                <button class="btn-secondary" on:click={() => viewSlotDetails(slot.id)}>View Details</button>
                                {#if slot.booked}
                                    <button class="btn-secondary" on:click={() => openFeedbackModal(slot.id)}>Enter Feedback</button>
                                {/if}
                            </div>
                        </li>
                    {/each}
                </ul>
                <div class="pagination">
                    <button on:click={() => changePage(currentPage - 1)} disabled={currentPage === 1} class="btn-secondary">Previous</button>
                    <span>Page {currentPage} of {upcomingSlots.totalPages}</span>
                    <button on:click={() => changePage(currentPage + 1)} disabled={currentPage === upcomingSlots.totalPages} class="btn-secondary">Next</button>
                </div>
            {:else}
                <p class="no-slots">No upcoming slots.</p>
            {/if}
        </section>

        <section class="card quick-actions">
            <h2>Quick Actions</h2>
            <div class="action-buttons">
                <a href="/coach/students-with-sessions" class="btn-secondary">View Past Sessions</a>
            </div>
        </section>
    </div>
</div>

{#if selectedSlotId}
    <SlotDetails 
        slotId={selectedSlotId} 
        userRole="coach" 
        on:close={closeSlotDetails}
    />
{/if}

{#if feedbackSlotId}
    <CoachFeedback 
        slotId={feedbackSlotId} 
        on:close={closeFeedbackModal}
        on:feedbackSubmitted={handleFeedbackSubmitted}
    />
{/if}
<style>
    .dashboard {
        max-width: 1200px;
        margin: 0 auto;
        padding: 2rem;
        font-family: Arial, sans-serif;
    }

    header {
        margin-bottom: 2rem;
        text-align: center;
    }

    h1 {
        color: #2c3e50;
        font-size: 2.5rem;
        margin-bottom: 0.5rem;
    }

    .coach-name {
        color: #3498db;
        font-weight: bold;
    }

    .warning {
        color: #e74c3c;
    }

    .dashboard-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
        gap: 2rem;
    }

    .card {
        background-color: #fff;
        border-radius: 8px;
        padding: 1.5rem;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        transition: all 0.3s ease;
    }

    .card:hover {
        transform: translateY(-5px);
        box-shadow: 0 6px 8px rgba(0, 0, 0, 0.15);
    }

    h2 {
        color: #2c3e50;
        font-size: 1.5rem;
        margin-bottom: 1rem;
        border-bottom: 2px solid #3498db;
        padding-bottom: 0.5rem;
    }

    .form-group {
        margin-bottom: 1rem;
    }

    label {
        display: block;
        margin-bottom: 0.5rem;
        color: #34495e;
    }

    input[type="date"],
    select {
        width: 100%;
        padding: 0.5rem;
        border: 1px solid #bdc3c7;
        border-radius: 4px;
        font-size: 1rem;
    }

    .btn-primary,
    .btn-secondary {
        display: inline-block;
        padding: 0.5rem 1rem;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 1rem;
        transition: background-color 0.3s ease;
    }

    .btn-primary {
        background-color: #3498db;
        color: white;
    }

    .btn-primary:hover {
        background-color: #2980b9;
    }

    .btn-secondary {
        background-color: #2ecc71;
        color: white;
    }

    .btn-secondary:hover {
        background-color: #27ae60;
    }

    .slot-list {
        list-style-type: none;
        padding: 0;
    }

    .slot-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 1rem;
        border-bottom: 1px solid #ecf0f1;
    }

    .slot-item:last-child {
        border-bottom: none;
    }

    .slot-info {
        display: flex;
        flex-direction: column;
    }

    .slot-date {
        font-weight: bold;
        color: #2c3e50;
    }

    .slot-status {
        padding: 0.25rem 0.5rem;
        border-radius: 12px;
        font-size: 0.9em;
        margin-top: 0.5rem;
    }

    .slot-status.available {
        background-color: #2ecc71;
        color: white;
    }

    .slot-status.booked {
        background-color: #f39c12;
        color: white;
    }

    .slot-actions {
        display: flex;
        gap: 0.5rem;
    }

    .no-slots {
        color: #7f8c8d;
        font-style: italic;
    }
    .pagination {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-top: 1rem;
    }

    .pagination button {
        padding: 0.5rem 1rem;
        background-color: #3498db;
        color: white;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    .pagination button:hover {
        background-color: #2980b9;
    }

    .pagination button:disabled {
        background-color: #bdc3c7;
        cursor: not-allowed;
    }

    @media (max-width: 768px) {
        .dashboard-grid {
            grid-template-columns: 1fr;
        }

        .slot-item {
            flex-direction: column;
            align-items: flex-start;
        }

        .slot-actions {
            margin-top: 1rem;
        }
    }
</style>