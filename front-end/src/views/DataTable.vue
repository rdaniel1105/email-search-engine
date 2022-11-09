<template>
  <table
    class="w-full w-max-[50%] rounded-lg bg-white divide-y divide-gray-300"
  >
    <thead class="bg-gray-900">
      <tr class="text-white text">
        <th class="font-semibold text-sm uppercase px-6 py-4">Subject</th>
        <th class="font-semibold text-sm uppercase px-6 py-4">From</th>
        <th class="font-semibold text-sm uppercase px-6 py-4">To</th>
        <th class="font-semibold text-sm uppercase px-6 py-4">Date</th>
      </tr>
    </thead>
    <tbody
      class="divide-y divide-gray-200"
      v-for="emails in emailsReceived.hits"
      :key="emails._id"
    >
      <tr
        class="
          text-gray-500 text-xs
          font-semibold
          hover:bg-green-400
          focus-within:bg-green-400/60
        "
        id="selectedEmail"
        @click="$emit('display-email', emails._source?.Body)"
        tabindex="0"
      >
        <td class="px-4 py-1">
          <button>{{ emails._source?.Subject }}</button>
        </td>
        <td class="px-4 py-1">
          <button>{{ emails._source?.From }}</button>
        </td>
        <td class="px-4 py-1">
          <button class="truncate max-w-[120px]">{{ emails._source?.To }}</button>
        </td>
        <td class="px-4 py-1">
          <button>{{ emails._source?.Date }}</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { MatchedEmails } from "@/types/interface";

export default defineComponent({
  data() {
    return {};
  },
  emits: ["display-email"],
  props: {
    emailsReceived: {
      type: Object as PropType<MatchedEmails>,
      required: true,
    },
  },
  // methods: {
  //   changeBGColor(): void {
  //   const box = document.getElementById('selectedEmail')
    
  //   if (box != null) {
  //     box.style.backgroundColor = "rgb(102, 187, 106)"; 
  //   }
  // }
  // }
});

</script>

