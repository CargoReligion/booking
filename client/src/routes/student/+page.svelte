<script lang="ts">
    import { onMount } from 'svelte';
    import { api } from '$lib/api';
    import type { SlotData, ApiResponse } from '../../types';
    import type { AxiosResponse } from 'axios';
  
    let availableSlots: SlotData[] = [];
    let upcomingBookings: SlotData[] = [];
    let currentStudentId: string = 'test'; // This should be set to the actual logged-in student's ID
  
    onMount(async () => {
      await refreshData();
    });
  
    async function refreshData() {
      try {
        const [availableResponse, upcomingResponse]: [
          AxiosResponse<ApiResponse<SlotData[]>>,
          AxiosResponse<ApiResponse<SlotData[]>>
        ] = await Promise.all([
          api.getAvailableSlots(),
          api.getUpcomingBookingsForStudent()
        ]);
        availableSlots = availableResponse.data.data;
        upcomingBookings = upcomingResponse.data.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }
  
    async function bookSlot(id: string) {
      try {
        await api.bookSlot(id);
        alert('Slot booked successfully!');
        await refreshData();
      } catch (error) {
        console.error('Error booking slot:', error);
        alert('Failed to book slot. Please try again.');
      }
    }
  </script>
  
  <h1>Student Dashboard</h1>
  
  <h2>Available Slots:</h2>
  <ul>
    {#each availableSlots as slot (slot.id)}
      <li>
        {slot.startTime} - Coach: {slot.coachId}
        <button on:click={() => bookSlot(slot.id)}>Book</button>
      </li>
    {/each}
  </ul>
  
  <h2>Upcoming Bookings:</h2>
  <ul>
    {#each upcomingBookings as booking (booking.id)}
      <li>{booking.startTime} - Coach: {booking.coachId}</li>
    {/each}
  </ul>