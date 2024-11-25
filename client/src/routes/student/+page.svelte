<script lang="ts">
    import { onMount } from 'svelte';
    import { api } from '$lib/api';
    import type { SlotData, Paginated } from '../../types';
    import { currentUser } from '$lib/userStore';
    import BookSession from './BookSession.svelte';
    import { formatDate } from '$lib/utils';
    import SlotDetails from '$lib/SlotDetails.svelte';
    import { userChangeStore } from '$lib/userChangeStore';

    let upcomingBookings: Paginated<SlotData> = {
        data: [],
        page: 1,
        pageSize: 10,
        totalCount: 0,
        totalPages: 0
    };
    let currentStudentId: string | null = null;
    let showBookingComponent = false;
    let selectedSlotId: string | null = null;

    currentUser.subscribe(user => {
        currentStudentId = user?.id || null;
    });
    userChangeStore.subscribe(user => {
        // Close slot details window when user changes
        selectedSlotId = null;
    });
    onMount(async () => {
        if (currentStudentId) {
            await refreshBookings();
        } else {
            console.log('No student selected. Please use the impersonate dropdown to select a student.');
        }
    });

    async function refreshBookings() {
        try {
            upcomingBookings = await api.getUpcomingBookingsForStudent();
        } catch (error) {
            console.error('Error fetching upcoming bookings:', error);
        }
    }

    function toggleBookingComponent() {
        showBookingComponent = !showBookingComponent;
    }

    async function handleBookingComplete() {
        toggleBookingComponent();
        await refreshBookings();
    }
    function viewSlotDetails(slotId: string) {
        selectedSlotId = slotId;
    }

    function closeSlotDetails() {
        selectedSlotId = null;
    }
</script>

<h1>Student Dashboard</h1>

{#if $currentUser}
    <p>Current Student: {$currentUser.name}</p>
{:else}
    <p>No student selected. Please use the impersonate dropdown to select a student.</p>
{/if}

{#if showBookingComponent}
    <button class="link-button" on:click={toggleBookingComponent}>Go Back</button>
    <BookSession on:bookingComplete={handleBookingComplete} />
{:else}
    <button on:click={toggleBookingComponent}>Book a Session</button>

    <h2>Upcoming Bookings:</h2>
    {#if upcomingBookings.data.length > 0}
        <ul>
            {#each upcomingBookings.data as booking}
                <li>{formatDate(booking.startTime)} - Coach: {booking.coachName}</li>
                <button on:click={() => viewSlotDetails(booking.id)}>View Details</button>
            {/each}
        </ul>
    {:else}
        <p>You have no upcoming bookings.</p>
    {/if}
{/if}
{#if selectedSlotId}
  <SlotDetails 
    slotId={selectedSlotId} 
    userRole="student" 
    on:close={closeSlotDetails}
  />
{/if}
<style>
    .link-button {
        background: none;
        border: none;
        padding: 0;
        color: blue;
        text-decoration: underline;
        cursor: pointer;
        font-size: 1em;
    }

    .link-button:hover,
    .link-button:focus {
        text-decoration: none;
    }
</style>