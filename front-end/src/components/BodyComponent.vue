<template>
  <div>
    <div>
        <table>
            <thead>
                <tr>
                    <th>
                        <div v-for="th in tableHeader" :key="th" class="">
                            <span>{{ th }}</span>
                        </div>
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="td in tableData" :key="td">
                    <td>{{ td.messageID }}</td>
                    <td>{{ td.subject }}</td>
                    <td>{{ td.from }}</td>
                    <td>{{ td.to }}</td>
                    <td>{{ td.date }}</td>
                </tr>
            </tbody>
        </table>
    </div>
    <div>

    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'TableData',
  props: {
    columns: Array,
    entries: Array
  },
  computed: {
    tableHeader() : Array<any> {
        return this.columns || [];
    },
    tableData() : Array<any> {
        return this.entries || []
    }
  }
});
</script>
