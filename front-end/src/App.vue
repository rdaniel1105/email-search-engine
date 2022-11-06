<template>
  <div class="flex flex-col min-h-screen bg-cool-blacky">
    <site-navigation/>
    <search-box/>
    <main>
      
    </main>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SearchBox from './components/SearchBox.vue';
import SiteNavigation from './components/SiteNavigation.vue';

export default defineComponent({
  components: { SiteNavigation, SearchBox },
  name: 'App'
});
</script>

<style>

</style>
