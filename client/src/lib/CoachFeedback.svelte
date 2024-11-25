<!-- src/lib/CoachFeedback.svelte -->
<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import { api } from '$lib/api';
    import type { SessionFeedback , CreateSessionFeedback} from '../types';
    import StarRating from './StarRating.svelte';

  
    export let slotId: string;
  
    let satisfaction: number = 3;
    let notes: string = '';
    let error: string | null = null;
  
    const dispatch = createEventDispatcher();
    
    async function submitFeedback() {
      if (satisfaction < 1 || satisfaction > 5) {
        error = 'Satisfaction must be between 1 and 5.';
        return;
      }
      console.log('Submitting feedback:', { slotId, satisfaction, notes }); // Debug log
      try {
        const createSessionFeedback: CreateSessionFeedback = {
            slotId: slotId,
            satisfaction: satisfaction,
            notes: notes,
        };
        await api.createSessionFeedback(createSessionFeedback);
        dispatch('feedbackSubmitted');
        dispatch('close');
      } catch (err) {
        console.error('Error submitting feedback:', err);
        error = 'Failed to submit feedback. Please try again.';
      }
    }
  
    function close() {
      dispatch('close');
    }
    function handleRatingChange(event: CustomEvent<number>) {
        satisfaction = event.detail;
    }
  </script>
  
  <div class="modal">
    <h2>Session Feedback</h2>
    
    <label>
      Student Satisfaction (1-5):
      <StarRating bind:rating={satisfaction} on:ratingChange={handleRatingChange} />
    </label>
  
    <div class="notes-container">
        <label for="feedback-notes">Notes:</label>
        <textarea 
          id="feedback-notes" 
          bind:value={notes} 
          placeholder="Enter your feedback here..."
        ></textarea>
      </div>
  
    {#if error}
      <p class="error">{error}</p>
    {/if}
  
    <div class="buttons">
      <button on:click={submitFeedback}>Save & Close</button>
      <button on:click={close}>Cancel</button>
    </div>
  </div>
  
  <style>
    .modal {
      background: white;
      padding: 20px;
      border-radius: 5px;
      box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
      max-width: 500px;
      width: 100%;
      margin: auto;
    }
  
    .rating-container, .notes-container {
      margin-bottom: 20px;
    }
  
    label {
      display: block;
      margin-bottom: 5px;
      font-weight: bold;
    }
  
    textarea {
      width: 100%;
      height: 150px;
      padding: 10px;
      border: 1px solid #ccc;
      border-radius: 4px;
      resize: vertical;
      font-family: inherit;
      font-size: 14px;
    }
  
    .error {
      color: red;
      margin-bottom: 10px;
    }
  
    .buttons {
      display: flex;
      justify-content: flex-end;
      gap: 10px;
    }
  
    button {
      padding: 8px 16px;
      border: none;
      border-radius: 4px;
      cursor: pointer;
      font-size: 14px;
    }
  
    button:first-child {
      background-color: #4CAF50;
      color: white;
    }
  
    button:last-child {
      background-color: #f44336;
      color: white;
    }
  </style>