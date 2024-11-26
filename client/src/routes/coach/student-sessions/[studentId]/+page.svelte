<!-- src/routes/coach/student-sessions/[studentId]/+page.svelte -->
<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { api } from '$lib/api';
    import type { SessionFeedback, ApiResponse } from '../../../../types';
    import { formatDate, UTCToLocal } from '$lib/utils';

    let sessionFeedbacks: SessionFeedback[] = [];
    let studentId: string;

    $: studentId = $page.params.studentId;

    onMount(async () => {
      try {
        const response: SessionFeedback[] = await api.getSessionFeedbackForStudent(studentId);
        sessionFeedbacks = response;
      } catch (error) {
        console.error('Error fetching session feedbacks:', error);
      }
    });

    function displayLocalTime(utcTimeString: string): string {
        const localDate = UTCToLocal(utcTimeString);
        return formatDate(localDate);
    }

    function getSatisfactionStars(satisfaction: number): string {
        return '★'.repeat(satisfaction) + '☆'.repeat(5 - satisfaction);
    }
</script>

<svelte:head>
    <title>Session Feedback for Student</title>
</svelte:head>

<div class="container">
    <h1>Session Feedback for Student</h1>

    <a href="/coach/students-with-sessions" class="back-link">← Back to Students List</a>

    {#if sessionFeedbacks.length === 0}
        <p>No feedback available for this student.</p>
    {:else}
        <div class="feedback-list">
            {#each sessionFeedbacks as feedback (feedback.id)}
                <div class="feedback-card">
                    <div class="feedback-date">{displayLocalTime(feedback.createdAt)}</div>
                    <div class="feedback-satisfaction" title="Satisfaction: {feedback.satisfaction} out of 5">
                        {getSatisfactionStars(feedback.satisfaction)}
                    </div>
                    <div class="feedback-notes">
                        <strong>Notes:</strong>
                        <p>{feedback.notes}</p>
                    </div>
                </div>
            {/each}
        </div>
    {/if}
</div>

<style>
    .container {
        max-width: 800px;
        margin: 0 auto;
        padding: 20px;
    }

    h1 {
        color: #333;
        margin-bottom: 20px;
    }

    .back-link {
        display: inline-block;
        margin-bottom: 20px;
        padding: 8px 12px;
        background-color: #f0f0f0;
        color: #333;
        text-decoration: none;
        border-radius: 4px;
        transition: background-color 0.3s ease;
    }

    .back-link:hover {
        background-color: #e0e0e0;
    }

    .feedback-list {
        display: grid;
        gap: 20px;
    }

    .feedback-card {
        background-color: #fff;
        border: 1px solid #ddd;
        border-radius: 8px;
        padding: 20px;
        box-shadow: 0 2px 4px rgba(0,0,0,0.1);
    }

    .feedback-date {
        font-size: 1.2em;
        color: #555;
        margin-bottom: 10px;
    }

    .feedback-satisfaction {
        font-size: 1.5em;
        color: #ffd700;
        margin-bottom: 10px;
    }

    .feedback-notes strong {
        display: block;
        margin-bottom: 5px;
        color: #333;
    }

    .feedback-notes p {
        margin: 0;
        color: #666;
        line-height: 1.4;
    }
</style>