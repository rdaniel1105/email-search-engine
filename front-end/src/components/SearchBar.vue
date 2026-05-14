<template>
  <form
    @submit.prevent="emit('search')"
    class="border-2 border-ink bg-paper-deep/60 backdrop-blur-[1px] focus-within:bg-paper-deep transition-colors"
  >
    <div class="flex flex-col md:flex-row items-stretch">
      <label class="sr-only" for="search-input">Search term</label>
      <div class="flex-1 flex items-center gap-3 px-5 py-4 border-b md:border-b-0 md:border-r border-ink/40">
        <svg
          class="w-5 h-5 shrink-0 text-ink-soft"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="1.75"
          aria-hidden="true"
        >
          <circle cx="11" cy="11" r="7" />
          <path d="m20 20-3.5-3.5" stroke-linecap="round" />
        </svg>
        <input
          id="search-input"
          :value="modelValue"
          @input="emit('update:modelValue', ($event.target as HTMLInputElement).value)"
          type="search"
          autocomplete="off"
          spellcheck="false"
          :placeholder="placeholder"
          class="w-full bg-transparent border-0 outline-none font-serif text-xl placeholder:text-ink-faint placeholder:italic text-ink"
        />
      </div>

      <div class="flex items-stretch divide-x divide-ink/40 md:divide-x">
        <div class="flex items-center gap-2 px-4 py-3">
          <span class="eyebrow">per page</span>
          <select
            :value="limit"
            @change="emit('update:limit', Number(($event.target as HTMLSelectElement).value))"
            class="bg-transparent font-mono text-sm font-medium text-ink outline-none cursor-pointer"
            aria-label="Results per page"
          >
            <option v-for="opt in limitOptions" :key="opt" :value="opt">{{ opt }}</option>
          </select>
        </div>

        <button
          type="submit"
          class="px-6 py-3 bg-ink text-paper font-mono uppercase tracking-widest text-xs hover:bg-oxblood transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
          :disabled="!modelValue.trim()"
        >
          Search
        </button>
      </div>
    </div>
  </form>
</template>

<script setup lang="ts">
interface Props {
  modelValue: string;
  limit: number;
  placeholder?: string;
}

withDefaults(defineProps<Props>(), {
  placeholder: 'Search the corpus by phrase, name, or subject…',
});

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
  (e: 'update:limit', value: number): void;
  (e: 'search'): void;
}>();

const limitOptions = [10, 25, 50, 100];
</script>
