<template>
  <header class="border-b border-ink pt-10 pb-6 mb-8">
    <div class="max-w-masthead mx-auto px-6 md:px-10">
      <div class="flex items-baseline justify-between gap-6 flex-wrap">
        <div class="eyebrow">Vol. I &middot; No. 01</div>
        <div class="eyebrow hidden md:block">Established MMXXII</div>
        <time class="eyebrow" :datetime="todayISO">{{ todayLong }}</time>
      </div>

      <div class="my-4 paper-rule origin-left animate-draw-line"></div>

      <h1 class="masthead-title text-[clamp(2.5rem,8vw,6rem)] text-ink leading-none">
        The Archive
      </h1>

      <div class="my-3 paper-rule origin-left animate-draw-line" style="animation-delay: 0.15s"></div>

      <p class="font-serif italic text-ink-soft text-lg max-w-prose">
        Full-text search across an indexed email corpus &mdash; companion to the
        <span class="font-mono not-italic text-[0.95em] text-oxblood">indexer</span>
        and <span class="font-mono not-italic text-[0.95em]">ZincSearch</span>.
      </p>
    </div>
  </header>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const now = new Date();

const todayISO = computed(() => now.toISOString().slice(0, 10));

const todayLong = computed(() =>
  now.toLocaleDateString('en-US', {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  }),
);
</script>
