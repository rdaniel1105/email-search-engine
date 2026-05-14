<template>
  <div class="border border-ink/20 bg-paper-deep/30 p-10 text-center flex flex-col items-center gap-3">
    <div class="eyebrow" :class="accentClass">{{ kicker }}</div>
    <p class="display-headline text-2xl text-ink max-w-prose">{{ title }}</p>
    <p v-if="detail" class="font-serif italic text-ink-muted max-w-prose">{{ detail }}</p>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { SearchStatus } from '@/types/interface';

interface Props {
  status: Extract<SearchStatus, 'idle' | 'loading' | 'empty' | 'error'>;
  detail?: string;
}

const props = defineProps<Props>();

const KICKERS: Record<Props['status'], string> = {
  idle: 'Awaiting query',
  loading: 'Searching',
  empty: 'No matches',
  error: 'Error',
};

const TITLES: Record<Props['status'], string> = {
  idle: 'Enter a phrase, name, or subject to begin.',
  loading: 'Combing the corpus…',
  empty: 'No dispatches match this query.',
  error: 'The search could not be completed.',
};

const kicker = computed(() => KICKERS[props.status]);
const title = computed(() => TITLES[props.status]);

const accentClass = computed(() => (props.status === 'error' ? 'text-oxblood' : 'text-ink-muted'));
</script>
