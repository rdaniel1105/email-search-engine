<template>
  <div>
    <form>
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
            max-w-[100%]
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
            <button
              id="dropdownLimitButton"
              data-dropdown-toggle="dropdownLimit"
              data-dropdown-placement="rigth"
              class="flex"
              viewBox="0 0 20 20"
              @click.prevent="toggleFilter()"
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
          <div
                id="dropdownLimit"
                aria-labelledby="dropdownLimitButton"
                class="
                absolute
                rounded-t bg-gray-200 py-2 px-4 whitespace-no-wrap 
                  hidden w-[10%]
                  z-10 mt-10 left-[87%]
                  rounded
                  divide-y divide-gray-100
                  shadow
                  dark:bg-gray-700
                "
              >
              <p>Select a limit</p>
              <input type="text" class="dark:bg-gray-700/70
                border-none
                flex-grow flex-wrap flex-shrink
                basis-0
                h-9
                w-full
                -mt-1
                pb-1
                focus:border-none focus:outline-none
                text-cool-blacky text-lg
                focus:transparent" v-model="limit">
              <div class=" bg-slate-900 rounded">
                <button class="w-full outline-none border-none" @click.prevent="toggleFilter()">Set</button>
              </div>
              </div>
        </div>
      </div>
    </form>
  </div>
  <div v-if="data.hits" class="relative w-full p-2">
    <div class="relative float-left text-white max-w-[53%] pb-2 pt-2">
      <Datatable :emails-received="data" @display-email="displayEmail" />
    </div>
    <div
      class="
        relative
        text-white
        h-full
        w-full
        float-right
        max-w-[43%]
        pl-0
        p-3.5
        whitespace-pre-line
      "
    >
      <p>{{ emailInside }}</p>
    </div>
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
    let emails = reactive<{ data: MatchedEmails }>({ data: {} });
    let searchTerm = ref("");
    let showEntries = ref([5, 10, 15, 20, 25, 50, 75, 100, 125, 150]);
    let limit = ref(25);
    let filteredEntries = ref([]);
    let currentPage = ref(1);
    let allPages = ref(1);
    let email = reactive<{ emailInside: any | undefined }>({
      emailInside: "",
    });

    const searchEmail = async (search: string): Promise<void> => {
      interface Termix {
        term: string;
      }

      let Term: Termix = {
        term: search,
      };

      email.emailInside = "";

      try {
        let queryLimit = limit.value
        const email = await emailsSearch(Term,queryLimit);
        emails.data = email;

        console.log("entries", emails);
      } catch (err) {
        console.log(err);
      }
    };

    const displayEmail = (body: any) => {
      email.emailInside = body;
      backToTop();
    };

    const backToTop = () => {
      document.body.scrollTop = 0;
      document.documentElement.scrollTop = 0;
    };

   const toggleFilter = () => {
    const limit = document.getElementById('dropdownLimit')
    if (limit === null) return
    
    if (limit.style.display === "none" || limit.style.display === "") {
      limit.style.display = "block"; 
      return
    }

    limit.style.display = "none"
  }

    return {
      searchEmail,
      displayEmail,
      backToTop,
      toggleFilter,
      ...toRefs(emails),
      ...toRefs(email),
      email,
      searchTerm,
      showEntries,
      limit,
      filteredEntries,
      currentPage,
      allPages,
    };
  },
});
</script>