<template>
  <section class="animate-rise">
    <SearchBar
      v-model="searchTerm"
      v-model:limit="limit"
      @search="onSearch"
    />

    <div v-if="status === 'success'" class="flex items-baseline justify-between mt-6 mb-4 flex-wrap gap-2">
      <div class="eyebrow">
        <span class="text-ink">{{ totalLabel }}</span> matches for
        <span class="font-serif not-italic italic text-ink text-base ml-1">&ldquo;{{ lastQuery }}&rdquo;</span>
      </div>
      <ArchivePagination :current-page="currentPage" :total-pages="totalPages" @change="onPageChange" />
    </div>

    <div v-if="showStatusPanel" class="mt-8">
      <StatusPanel :status="statusForPanel" :detail="errorDetail" />
    </div>

    <div v-else class="mt-2 grid lg:grid-cols-[1.05fr_1fr] gap-8 lg:gap-10 items-start">
      <ol class="stagger divide-y divide-ink/15 -mx-5">
        <li
          v-for="(hit, idx) in results.hits"
          :key="hit._id || idx"
        >
          <EmailListItem
            :email="hit"
            :index="from + idx + 1"
            :selected="selectedId === (hit._id ?? String(idx))"
            @select="selectEmail(hit, idx)"
          />
        </li>
      </ol>

      <EmailReader :email="selectedEmail" />
    </div>

    <ArchivePagination
      v-if="status === 'success'"
      :current-page="currentPage"
      :total-pages="totalPages"
      @change="onPageChange"
      class="mt-10"
    />
  </section>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
import SearchBar from '@/components/SearchBar.vue';
import EmailListItem from '@/components/EmailListItem.vue';
import EmailReader from '@/components/EmailReader.vue';
import ArchivePagination from '@/components/ArchivePagination.vue';
import StatusPanel from '@/components/StatusPanel.vue';
import { emailsSearch, SearchError } from '@/services/emailsAPI';
import type { EmailHit, SearchResponse, SearchStatus } from '@/types/interface';

const searchTerm = ref('');
const lastQuery = ref('');
const limit = ref(25);
const currentPage = ref(1);
const status = ref<SearchStatus>('idle');
const errorDetail = ref<string | undefined>(undefined);
const results = reactive<{ hits: EmailHit[]; total: number }>({ hits: [], total: 0 });
const selectedId = ref<string | null>(null);
const selectedEmail = ref<EmailHit | null>(null);

let currentRequest: AbortController | null = null;

const totalPages = computed(() => {
  if (results.total === 0 || limit.value === 0) return 1;
  return Math.max(1, Math.ceil(results.total / limit.value));
});

const totalLabel = computed(() => results.total.toLocaleString('en-US'));

const from = computed(() => (currentPage.value - 1) * limit.value);

const statusForPanel = computed(() => status.value as 'idle' | 'loading' | 'empty' | 'error');

const showStatusPanel = computed(
  () => status.value === 'idle' || status.value === 'loading' || status.value === 'empty' || status.value === 'error',
);

async function performSearch(): Promise<void> {
  const term = searchTerm.value.trim();
  if (!term) return;

  lastQuery.value = term;
  status.value = 'loading';
  errorDetail.value = undefined;
  selectedId.value = null;
  selectedEmail.value = null;

  currentRequest?.abort();
  const controller = new AbortController();
  currentRequest = controller;

  try {
    const response: SearchResponse = await emailsSearch({
      term: { term },
      limit: limit.value,
      from: from.value,
      signal: controller.signal,
    });

    if (controller.signal.aborted) return;

    results.hits = response.hits ?? [];
    results.total = response.total?.value ?? 0;

    status.value = results.total === 0 ? 'empty' : 'success';
  } catch (err) {
    if ((err as DOMException)?.name === 'AbortError') return;
    status.value = 'error';
    errorDetail.value = err instanceof SearchError ? err.message : 'An unexpected error occurred.';
    results.hits = [];
    results.total = 0;
  } finally {
    if (currentRequest === controller) currentRequest = null;
  }
}

function onSearch(): void {
  currentPage.value = 1;
  void performSearch();
}

function onPageChange(page: number): void {
  if (page === currentPage.value) return;
  currentPage.value = Math.min(Math.max(1, page), totalPages.value);
  void performSearch();
  scrollToTop();
}

function selectEmail(hit: EmailHit, idx: number): void {
  selectedId.value = hit._id ?? String(idx);
  selectedEmail.value = hit;
}

function scrollToTop(): void {
  window.scrollTo({ top: 0, behavior: 'smooth' });
}
</script>
