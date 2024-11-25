<script lang="ts">
    import { onMount } from 'svelte';
    import { api } from '$lib/api';
    import type { SlotData, User, Paginated } from '../../types';
    import { currentUser, allUsers } from '$lib/userStore';
    import { formatDate } from '$lib/utils';
    import BookSession from './BookSession.svelte';

    let upcomingBookings: Paginated<SlotData> = {
        data: [],
        page: 1,
        pageSize: 10,
        totalCount: 0,
        totalPages: 0
    };
    let currentStudentId: string | null = null;
    let showBookingComponent = false;
    let coaches: User[] = [];

    currentUser.subscribe(user => {
        currentStudentId = user?.id || null;
    });

    $: {
        console.log('Reactive statement running, allUsers:', $allUsers);
        coaches = $allUsers.filter(user => user.role === 'coach');
        console.log('Filtered coaches:', coaches);
    }
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

    $: {
        if ($allUsers.length > 0) {
            coaches = $allUsers.filter(user => user.role === 'coach');
            console.log('Coaches updated after allUsers change:', coaches);
        }
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
    <BookSession />
{:else}
    <button on:click={toggleBookingComponent}>Book a Session</button>

    <!-- <h2>Upcoming Bookings:</h2>
    {#if upcomingBookings.data.length > 0}
        <ul>
            {#each upcomingBookings.data as booking}
                <li>{formatDate(booking.startTime)} - Coach: {booking.coachId}</li>
            {/each}
        </ul>
    {:else}
        <p>You have no upcoming bookings.</p>
    {/if} -->
{/if}