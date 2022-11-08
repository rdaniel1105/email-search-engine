<!-- <template>
    <table-data :columns="columns" :entries="entries.hits"/>
</template>

<script lang="ts">
import TableData from '@/components/BodyComponent.vue'
import SearchInput from '../components/SearchBarTest.vue';
import { emailsSearch } from '@/services/emailsAPI';
import { EmailsTypes, MatchedEmails } from '@/types/interface';


export default {
    name: 'DataTable',
    data() {
        return {
            columns: [
            { name: 'messageID', text: 'Message-ID'},
            { name: 'subject', text: 'Subject'},
            { name: 'from', text: 'From'},
            { name: 'to', text: 'To'},
            { name: 'date', text: 'Date'},
            { name: 'body', text: 'Body'},
        ],
        entries: {} as MatchedEmails,
        };
    },
    components: {
        TableData,
        SearchInput
    },
    methods: {
        async searchEmail(searchTerm :string): Promise<void>{
            console.log(searchTerm);
            
            const resp = await emailsSearch(searchTerm);
            console.log(resp)
        }
    }
}
</script> -->