<template>
  <div class="flex flex-col min-h-screen bg-cool-blacky">
    <site-navigation/>
    <main class="h-full">
      <search-body/>
    </main>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
// import SearchBox from './components/SearchBox.vue';
import SiteNavigation from '@/components/SiteNavigation.vue';
import SearchBody from '@/components/SearchBody.vue';

export default defineComponent({
  components: { SiteNavigation, SearchBody },
  name: 'App',
});
</script>
