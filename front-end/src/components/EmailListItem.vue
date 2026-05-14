<template>
  <article
    tabindex="0"
    role="button"
    :aria-pressed="selected"
    @click="emit('select')"
    @keydown.enter.prevent="emit('select')"
    @keydown.space.prevent="emit('select')"
    class="group cursor-pointer pl-5 pr-4 py-5 border-l-2 transition-colors"
    :class="selected ? 'border-oxblood bg-paper-deep/70' : 'border-transparent hover:border-ink hover:bg-paper-deep/40'"
  >
    <div class="flex items-baseline justify-between gap-4 mb-2">
      <span class="eyebrow text-oxblood">№ {{ index }}</span>
      <time v-if="email.date" class="font-mono text-[0.7rem] text-ink-muted">{{ shortDate }}</time>
    </div>

    <h3 class="display-headline text-xl md:text-2xl text-ink mb-2 group-hover:underline underline-offset-[6px] decoration-ink/30">
      {{ email.subject || 'Untitled dispatch' }}
    </h3>

    <div class="font-serif text-[0.95rem] text-ink-soft mb-1 italic">
      <span class="smallcaps text-ink-muted not-italic mr-1">from</span>{{ email.from || 'unknown sender' }}
    </div>
    <div v-if="email.to" class="font-serif text-[0.9rem] text-ink-muted mb-3 italic truncate">
      <span class="smallcaps not-italic mr-1">to</span>{{ email.to }}
    </div>

    <p v-if="preview" class="font-serif text-ink-soft text-[0.95rem] leading-snug line-clamp-2">
      {{ preview }}
    </p>
  </article>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { EmailHit } from '@/types/interface';

interface Props {
  email: EmailHit;
  index: number;
  selected: boolean;
}

const props = defineProps<Props>();

const emit = defineEmits<{
  (e: 'select'): void;
}>();

const shortDate = computed(() => {
  const raw = props.email.date;
  if (!raw) return '';
  const parsed = new Date(raw);
  if (Number.isNaN(parsed.getTime())) return raw;
  return parsed.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' });
});

const preview = computed(() => {
  const body = props.email.body ?? '';
  return body.replace(/\s+/g, ' ').trim().slice(0, 220);
});
</script>
