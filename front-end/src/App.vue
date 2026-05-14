<template>
  <div class="min-h-screen flex flex-col grain">
    <SiteMasthead />
    <main class="flex-1 max-w-masthead w-full mx-auto px-6 md:px-10 pb-24">
      <SearchBody />
    </main>
    <SiteFooter />
  </div>
</template>

<script setup lang="ts">
import SiteMasthead from '@/components/SiteMasthead.vue';
import SiteFooter from '@/components/SiteFooter.vue';
import SearchBody from '@/components/SearchBody.vue';
</script>
