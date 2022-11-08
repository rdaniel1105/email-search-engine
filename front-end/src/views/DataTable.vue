<template>
  <table
    class="
      mx-auto
      max-w-4xl
      w-full
      whitespace-nowrap
      rounded-lg
      bg-white
      divide-y divide-gray-300
      overflow-hidden
    "
  >
    <thead class="bg-gray-900">
      <tr class="text-white text-left font-semibold text-sm uppercase">
        <th class="px-6 py-4">Subject</th>
        <th class="px-6 py-4">From</th>
        <th class="px-6 py-4">To</th>
        <th class="px-6 py-4">Date</th>
      </tr>
    </thead>
    <tbody
      v-for="emails in tableEmails.hits"
      :key="emails._id"
      class="divide-y divide-gray-200"
    >
      <tr class="text-gray-500 text-sm font-semibold tracking-wide">
        <td class="px-4 py-1">{{ emails._source?.Subject }}</td>
        <td class="px-4 py-1">{{ emails._source?.From }}</td>
        <td class="px-4 py-1">{{ emails._source?.To }}</td>
        <td class="px-4 py-1">{{ emails._source?.Date }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { MatchedEmails } from "@/types/interface";

export default defineComponent({
  props: {
    email: {
      type: Object as PropType<MatchedEmails>,
      required: true,
    },
  },
  computed: {
    tableEmails() {
      return this.email;
    },
  },
});
</script>