<template>
  <div class="flex flex-col min-h-screen bg-cool-blacky">
    <site-navigation/>
    <main>
      <search-input/>
    </main>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
// import SearchBox from './components/SearchBox.vue';
import SiteNavigation from '@/components/SiteNavigation.vue';
import SearchInput from '@/components/SearchBarTest.vue';

export default defineComponent({
  components: { SiteNavigation, SearchInput },
  name: 'App',
});
</script>
