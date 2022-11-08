<template>
  <div>
    <form @click.prevent="">
      <div class="relative p-3.5 items-center px-5">
        <div
          class="
            focus-within:shadow-slate-400/80
            focus-within:bg-gray-300/75
            focus-within:rounded-none
            focus-within:rounded-tl-2xl
            focus-within:rounded-br-2xl
            flex
            pt-1.5
            pb-0
            pl-3.5
            text-base
            min-w-[200px]
            w-full
            max-w-2xl
            mx-auto
            appearance-none
            outline-none
            bg-none bg-gray-400/40
            p-3.5
            rounded-tr-2xl rounded-bl-2xl
            transition-200
            hover:rounded-none
            hover:rounded-tl-2xl
            hover:rounded-br-2xl
            hover:bg-gray-300/75
            focus:rounded-none focus:rounded-tl-2xl focus:rounded-br-2xl
            shadow-lg shadow-slate-100/50
            hover:shadow-slate-400/80
            border-solid border-violet-700 border-opacity-25 border-2
          "
        >
          <div class="flex inset-y-0 left-0 items-center -mt-0.5 pr-3 pb-1">
            <div class="block m-auto">
              <button
                id="lupa"
                @click.prevent="searchEmail(searchTerm)"
                class="relative h-5 leading-5 w-5 mt-[3px] inline-block dark"
              >
                <svg
                  aria-hidden="true"
                  class="block h-full w-full"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                  ></path>
                </svg>
              </button>
            </div>
          </div>
          <div
            class="
              flex-grow flex-wrap flex-shrink
              basis-0
              bg-transparent
              border-none
              m-0
              p-0
              text-transparent
            "
            tabindex="0"
          >
            <input
              type="search"
              v-model="searchTerm"
              id="default-search"
              class="
                bg-transparent
                border-none
                m-0
                p-0
                flex
                break-words
                placeholder:text-zinc-500 placeholder:text-base
                flex-grow flex-wrap flex-shrink
                basis-0
                h-9
                w-full
                -mt-1
                pb-1
                focus:border-none focus:outline-none
                text-cool-blacky text-lg
                focus:transparent
              "
              placeholder="Search..."
              required
            />
          </div>
          <div class="flex pl-1">
            <label
              for="search-dropdown"
              class="
                mb-2
                text-sm
                font-medium
                text-gray-900
                sr-only
                dark:text-gray-300
              "
            ></label>
            <button
              id="dropdown-button"
              data-dropdown-toggle="dropdown"
              class="flex"
              type="button"
              gh="sda"
              viewBox="0 0 20 20"
              role="button"
              data-tooltip="Show filters"
              aria-label="Filter options"
              style=""
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <path
                  d="M3 17v2h6v-2H3zM3 5v2h10V5H3zm10 16v-2h8v-2h-8v-2h-2v6h2zM7 9v2H3v2h4v2h2V9H7zm14 
                        4v-2H11v2h10zm-6-4h2V7h4V5h-4V3h-2v6z"
                ></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
    </form>
  </div>
  <div
    v-if="entries.hits"
    class="max-w-[50%] text-white overflow-x-auto w-full h-full p-4"
  >
    <Datatable :email="entries" />
  </div>
</template>
  
<script lang="ts">
import { defineComponent, reactive, ref, toRefs } from "vue";
import Datatable from "@/views/DataTable.vue";
import { emailsSearch } from "@/services/emailsAPI";
import { MatchedEmails } from "@/types/interface";

export default defineComponent({
  name: "SearchInput",
  components: {
    Datatable,
  },
  // data() {
  //     return {
  //     entries: [] as MatchedEmails,
  //     searchTerm: ""
  //     };
  // },
  // methods: {
  //     async searchEmail(searchTerm :string): Promise<void>{
  //         interface Termix {
  //             term: string
  //         }

  //         let Term: Termix = {
  //             term: searchTerm
  //         }

  //         const resp = await emailsSearch(Term);
  //         this.entries = resp;
  //     }
  // },
  setup() {
    let emails = reactive<{ entries: MatchedEmails }>({ entries: {} });
    let searchTerm = ref("");

    const searchEmail = async (search: string): Promise<void> => {
      interface Termix {
        term: string;
      }

      let Term: Termix = {
        term: search,
      };

      try {
        const email = await emailsSearch(Term);
        emails.entries = email;

        console.log("entries", emails);
      } catch (err) {
        console.log(err);
      }
    };

    return { searchEmail, ...toRefs(emails), searchTerm };
  },
});
</script>