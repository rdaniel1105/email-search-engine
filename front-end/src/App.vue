<template>
  <div class="flex flex-col min-h-screen bg-cool-blacky">
    <site-navigation/>
    <search-bar-test/>
    <main>
      
    </main>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
// import SearchBox from './components/SearchBox.vue';
import SiteNavigation from '@/components/SiteNavigation.vue';
import SearchBarTest from '@/components/SearchBarTest.vue';

export default defineComponent({
  components: { SiteNavigation, SearchBarTest },
  name: 'App'
});
</script>
