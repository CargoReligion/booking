<script lang="ts">
    import { onMount, createEventDispatcher } from 'svelte';
    import { api } from '$lib/api';
    import type { User, SlotData, Paginated } from '../../types';
    import { formatDate } from '$lib/utils';
    import { currentUser } from '$lib/userStore';

    // @ts-ignore
    const dispatch = createEventDispatcher();

    let coaches: User[] = [];
    let selectedCoach: User | null = null;
    let availableSlots: Paginated<SlotData> = {
        data: [],
        page: 1,
        pageSize: 10,
        totalCount: 0,
        totalPages: 0
    };
    let currentPage = 1;
    let currentStudentId: string | null = null;

    currentUser.subscribe(user => {
        currentStudentId = user?.id || null;
    });

    onMount(async () => {
        try {
            const response = await api.getAllUsers();
            coaches = response.filter(user => user.role === 'coach');
        } catch (error) {
            console.error('Error fetching coaches:', error);
        }
    });

    async function selectCoach(coach: User) {
        selectedCoach = coach;
        await fetchAvailableSlots(coach.id, 1);
    }

    async function fetchAvailableSlots(coachId: string, page: number) {
        try {
            availableSlots = await api.getAvailableSlots(coachId, page);
            currentPage = page;
        } catch (error) {
            console.error('Error fetching available slots:', error);
        }
    }

    function changePage(newPage: number) {
        if (selectedCoach && newPage >= 1 && newPage <= availableSlots.totalPages) {
            fetchAvailableSlots(selectedCoach.id, newPage);
        }
    }

    async function bookSlot(slotId: string) {
        if (!currentStudentId) {
            console.error('No student selected');
            return;
        }

        try {
            await api.bookSlot(slotId);
            alert('Slot booked successfully!');
            dispatch('bookingComplete');
        } catch (error) {
            console.error('Error booking slot:', error);
            alert('Failed to book slot. Please try again.');
        }
    }
</script>

{#if !selectedCoach}
    <h2>Select a Coach</h2>
    {#if coaches.length > 0}
        <ul>
            {#each coaches as coach}
                <li>
                    <button on:click={() => selectCoach(coach)}>{coach.name}</button>
                </li>
            {/each}
        </ul>
    {:else}
        <p>Loading coaches...</p>
    {/if}
{:else}
    <h2>Available Slots for {selectedCoach.name}</h2>
    {#if availableSlots.data.length > 0}
        <ul>
            {#each availableSlots.data as slot}
                <li>
                    {formatDate(slot.startTime)}
                    <button on:click={() => bookSlot(slot.id)}>Book</button>
                </li>
            {/each}
        </ul>
        
        <!-- Pagination UI -->
        <div class="pagination">
            <button on:click={() => changePage(currentPage - 1)} disabled={currentPage === 1}>
                &lt; Previous
            </button>
            
            {#each Array(availableSlots.totalPages) as _, i}
                <button 
                    on:click={() => changePage(i + 1)} 
                    class:active={currentPage === i + 1}
                >
                    {i + 1}
                </button>
            {/each}
            
            <button on:click={() => changePage(currentPage + 1)} disabled={currentPage === availableSlots.totalPages}>
                Next &gt;
            </button>
        </div>
    {:else}
        <p>No available slots for this coach.</p>
    {/if}
    
    <button on:click={() => selectedCoach = null}>Back to Coach Selection</button>
{/if}

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