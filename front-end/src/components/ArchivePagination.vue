<template>
  <nav v-if="totalPages > 1" class="flex items-center justify-between gap-4 mt-8 pt-4 border-t border-ink/20" aria-label="Pagination">
    <div class="eyebrow">
      Page <span class="text-ink">{{ currentPage }}</span> of <span class="text-ink">{{ totalPages }}</span>
    </div>

    <div class="flex items-center divide-x divide-ink/30 border border-ink/40 font-mono text-xs uppercase tracking-widest">
      <button
        type="button"
        class="px-3 py-2 hover:bg-ink hover:text-paper transition-colors disabled:opacity-30 disabled:hover:bg-transparent disabled:hover:text-ink disabled:cursor-not-allowed"
        :disabled="currentPage === 1"
        @click="emit('change', 1)"
        aria-label="First page"
      >
        &#xab;
      </button>
      <button
        type="button"
        class="px-3 py-2 hover:bg-ink hover:text-paper transition-colors disabled:opacity-30 disabled:hover:bg-transparent disabled:hover:text-ink disabled:cursor-not-allowed"
        :disabled="currentPage === 1"
        @click="emit('change', currentPage - 1)"
      >
        Prev
      </button>
      <button
        type="button"
        class="px-3 py-2 hover:bg-ink hover:text-paper transition-colors disabled:opacity-30 disabled:hover:bg-transparent disabled:hover:text-ink disabled:cursor-not-allowed"
        :disabled="currentPage === totalPages"
        @click="emit('change', currentPage + 1)"
      >
        Next
      </button>
      <button
        type="button"
        class="px-3 py-2 hover:bg-ink hover:text-paper transition-colors disabled:opacity-30 disabled:hover:bg-transparent disabled:hover:text-ink disabled:cursor-not-allowed"
        :disabled="currentPage === totalPages"
        @click="emit('change', totalPages)"
        aria-label="Last page"
      >
        &#xbb;
      </button>
    </div>
  </nav>
</template>

<script setup lang="ts">
interface Props {
  currentPage: number;
  totalPages: number;
}

defineProps<Props>();

const emit = defineEmits<{
  (e: 'change', page: number): void;
}>();
</script>
