<script lang="ts">
    import { onMount } from 'svelte';
    import { fade } from 'svelte/transition';
    import { api } from '$lib/api';
    import type { SlotData, Paginated, ApiResponse } from '../../types';
    import { currentUser } from '$lib/userStore';
    import BookSession from './BookSession.svelte';
    import { formatDate, UTCToLocal, UTCToEST } from '$lib/utils';

    let upcomingBookings: Paginated<SlotData> = {
        data: [],
        page: 1,
        pageSize: 10,
        totalCount: 0,
        totalPages: 0
    };
    let currentStudentId: string | null = null;
    let showBookingComponent = false;
    let currentPage = 1;

    currentUser.subscribe(user => {
        currentStudentId = user?.id || null;
    });

    onMount(async () => {
        if (currentStudentId) {
            await refreshBookings(1);
        } else {
            console.log('No student selected. Please use the impersonate dropdown to select a student.');
        }
    });

    async function refreshBookings(page: number) {
        try {
            upcomingBookings = await api.getUpcomingBookingsForStudent(page);
            currentPage = page;
        } catch (error) {
            console.error('Error fetching upcoming bookings:', error);
        }
    }

    function toggleBookingComponent() {
        showBookingComponent = !showBookingComponent;
    }

    async function handleBookingComplete() {
        toggleBookingComponent();
        await refreshBookings(1);
    }

    function changePage(newPage: number) {
        if (newPage >= 1 && newPage <= upcomingBookings.totalPages) {
            refreshBookings(newPage);
        }
    }

    function displayLocalTime(utcTimeString: string): string {
        const estDate = UTCToEST(utcTimeString);
        return formatDate(estDate, 'America/New_York');
    }
</script>

<svelte:head>
    <title>Student Dashboard</title>
</svelte:head>

<div class="dashboard">
    <header>
        <h1>Student Dashboard</h1>
        {#if $currentUser}
            <p>Welcome, <span class="student-name">{$currentUser.name}</span>!</p>
        {:else}
            <p class="warning">No student selected. Please use the impersonate dropdown to select a student.</p>
        {/if}
    </header>

    <div class="dashboard-grid">
        <section class="card upcoming-bookings">
            <h2>Upcoming Bookings</h2>
            {#if upcomingBookings.data.length > 0}
                <ul class="booking-list">
                    {#each upcomingBookings.data as booking (booking.id)}
                        <li class="booking-item" transition:fade="{{ duration: 300 }}">
                            <span class="booking-date">{displayLocalTime(booking.startTime)}</span>
                            <span class="booking-coach">Coach: {booking.coachName}</span>
                        </li>
                    {/each}
                </ul>
                <div class="pagination">
                    <button on:click={() => changePage(currentPage - 1)} disabled={currentPage === 1} class="btn-secondary">Previous</button>
                    <span>Page {currentPage} of {upcomingBookings.totalPages}</span>
                    <button on:click={() => changePage(currentPage + 1)} disabled={currentPage === upcomingBookings.totalPages} class="btn-secondary">Next</button>
                </div>
            {:else}
                <p class="no-bookings">You have no upcoming bookings.</p>
            {/if}
        </section>

        <section class="card quick-actions">
            <h2>Quick Actions</h2>
            <div class="action-buttons">
                <button on:click={toggleBookingComponent} class="btn-primary">
                    {showBookingComponent ? 'Close Booking' : 'Book a Session'}
                </button>
            </div>
        </section>
    </div>

    {#if showBookingComponent}
        <div class="booking-component" transition:fade="{{ duration: 300 }}">
            <BookSession on:bookingComplete={handleBookingComplete} />
        </div>
    {/if}
</div>
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

    .student-name {
        color: #e74c3c;
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
        border-bottom: 2px solid #e74c3c;
        padding-bottom: 0.5rem;
    }

    .booking-list {
        list-style-type: none;
        padding: 0;
    }

    .booking-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 1rem;
        border-bottom: 1px solid #ecf0f1;
    }

    .booking-item:last-child {
        border-bottom: none;
    }

    .booking-date {
        font-weight: bold;
        color: #2c3e50;
    }

    .booking-coach {
        color: #7f8c8d;
    }

    .no-bookings {
        color: #7f8c8d;
        font-style: italic;
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
        background-color: #e74c3c;
        color: white;
    }

    .btn-primary:hover {
        background-color: #c0392b;
    }

    .btn-secondary {
        background-color: #3498db;
        color: white;
    }

    .btn-secondary:hover {
        background-color: #2980b9;
    }

    .pagination {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-top: 1rem;
    }

    .booking-component {
        margin-top: 2rem;
        padding: 1rem;
        background-color: #f9f9f9;
        border-radius: 8px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }

    @media (max-width: 768px) {
        .dashboard-grid {
            grid-template-columns: 1fr;
        }

        .booking-item {
            flex-direction: column;
            align-items: flex-start;
        }

        .booking-coach {
            margin-top: 0.5rem;
        }
    }
</style>