<!-- src/lib/StarRating.svelte -->
<script lang="ts">
    import { createEventDispatcher } from 'svelte';
  
    export let rating: number = 0;
    export let maxStars: number = 5;
  
    let hoveredRating: number = 0;
  
    const dispatch = createEventDispatcher();
  
    function handleClick(star: number) {
      rating = star;
      dispatch('ratingChange', rating);
    }
  
    function handleMouseEnter(star: number) {
      hoveredRating = star;
    }
  
    function handleMouseLeave() {
      hoveredRating = 0;
    }
  </script>
  
  <div class="star-rating" on:mouseleave={handleMouseLeave}>
    {#each Array(maxStars) as _, i}
      <span
        class="star"
        class:filled={i < rating}
        class:hovered={i < hoveredRating}
        on:click={() => handleClick(i + 1)}
        on:mouseenter={() => handleMouseEnter(i + 1)}
        on:keypress={() => handleClick(i + 1)}
        tabindex="0"
        role="button"
      >
        ★
      </span>
    {/each}
  </div>
  
  <style>
    .star-rating {
      font-size: 24px;
      color: #ddd;
      display: inline-block;
    }
  
    .star {
      cursor: pointer;
      transition: color 0.2s ease;
    }
  
    .star.filled {
      color: #ffd700;
    }
  
    .star.hovered {
      color: #ffed85;
    }
  </style>