<template>
  <aside
    class="border border-ink/30 bg-paper-deep/40 p-6 md:p-8 lg:sticky lg:top-6 max-h-[calc(100vh-3rem)] overflow-y-auto"
    aria-live="polite"
  >
    <template v-if="email">
      <div class="eyebrow text-oxblood mb-3">Reading Pane</div>

      <h2 class="display-headline text-2xl md:text-3xl text-ink mb-4">
        {{ email._source?.Subject || 'Untitled dispatch' }}
      </h2>

      <dl class="grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-sm mb-6 border-t border-b border-ink/15 py-3">
        <dt class="eyebrow self-center">From</dt>
        <dd class="font-serif text-ink-soft break-all">{{ email._source?.From }}</dd>
        <dt class="eyebrow self-center">To</dt>
        <dd class="font-serif text-ink-soft break-all">{{ email._source?.To }}</dd>
        <dt v-if="email._source?.Date" class="eyebrow self-center">Date</dt>
        <dd v-if="email._source?.Date" class="font-mono text-xs text-ink-muted">{{ email._source.Date }}</dd>
        <dt v-if="messageId" class="eyebrow self-center">ID</dt>
        <dd v-if="messageId" class="font-mono text-[0.7rem] text-ink-faint break-all">{{ messageId }}</dd>
      </dl>

      <div class="font-serif text-ink leading-relaxed whitespace-pre-wrap text-[1rem]">
        <span v-if="bodyText" class="first-letter:font-display first-letter:text-5xl first-letter:font-bold first-letter:text-oxblood first-letter:float-left first-letter:mr-2 first-letter:leading-[0.85]">{{ bodyText }}</span>
        <span v-else class="italic text-ink-faint">No body content.</span>
      </div>
    </template>

    <template v-else>
      <div class="flex flex-col items-start gap-3 py-4">
        <div class="eyebrow text-ink-muted">Reading Pane</div>
        <p class="font-serif italic text-ink-muted max-w-prose">
          Select a dispatch from the column on the left to read it here.
        </p>
      </div>
    </template>
  </aside>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { EmailHit } from '@/types/interface';

interface Props {
  email: EmailHit | null;
}

const props = defineProps<Props>();

const messageId = computed(() => props.email?._source?.['Message-ID']);
const bodyText = computed(() => props.email?._source?.Body?.trim() ?? '');
</script>
