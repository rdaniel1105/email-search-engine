<template>
  <div id="wholePage">
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
            max-w-[70%]
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
                @click.prevent="
                  (showFilter = false),
                    (currentPage = 1),
                    (showEntries = true),
                    (allPages = 1),
                    searchEmail(searchTerm)
                "
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
              type="button"
              viewBox="0 0 20 20"
              @click="showFilter = !showFilter"
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
            :class="showFilter ? 'block' : 'hidden'"
            class="
              absolute
              rounded-t
              py-1
              px-4
              whitespace-no-wrap
              w-[10%]
              z-10
              mt-10
              left-[83%]
              rounded
              divide-y divide-gray-100
              shadow
              dark:bg-gray-600
            "
          >
            <p>Select a limit</p>
            <input
              type="text"
              class="
                dark:bg-slate-400
                rounded-md
                border-none
                flex-grow flex-wrap flex-shrink
                basis-0
                h-9
                w-full
                mt-1
                mb-1.5
                pb-1
                focus:border-none focus:outline-none
                text-cool-blacky text-lg
                focus:transparent
              "
              v-model="limit"
            />
            <div class="rounded">
              <button
                type="button"
                @click="
                  (showFilter = false),
                    (currentPage = 1),
                    (showEntries = true),
                    (allPages = 1),
                    searchEmail(searchTerm)
                "
                class="w-full outline-none border-none bg-violet-700"
              >
                Set
              </button>
            </div>
          </div>
        </div>
      </div>
    </form>
  </div>
  <div
    id="entries"
    :class="showEntries ? 'block' : 'hidden'"
    @click="showFilter = false"
    class="mx-2 my-2 text-white relative p-1 pb-0 rounded-xl"
  >
    <div>Showing {{ currentPage }} of {{ allPages }}</div>
    <ul class="flex w-[200px] divide-x divide-gray-300 rounded-lg border-2">
      <li
        :class="[{ 'pointer-events-none': currentPage === 1 }]"
        class="bg-white"
      >
        <a
          href="#"
          type="button"
          @click.prevent="(currentPage = 1), searchEmail(searchTerm)"
          class="text-cool-blacky space-x-0 hover:bg-cyan-600 px-1"
        >
          First
        </a>
      </li>
      <li
        :class="[{ 'pointer-events-none': currentPage === 1 }]"
        class="bg-white"
      >
        <a
          href="#"
          type="button"
          @click.prevent="
            currentPage < 1 ? (currentPage = 1) : (currentPage -= 1),
              searchEmail(searchTerm)
          "
          class="text-cool-blacky space-x-0 hover:bg-cyan-600 px-1"
        >
          Previous
        </a>
      </li>
      <li
        :class="[{ 'pointer-events-none': currentPage === allPages }]"
        class="bg-white"
      >
        <a
          href="#"
          type="button"
          @click.prevent="
            currentPage > allPages
              ? (currentPage = allPages)
              : (currentPage += 1),
              searchEmail(searchTerm)
          "
          class="text-cool-blacky space-x-0 hover:bg-cyan-600 px-1"
        >
          Next
        </a>
      </li>
      <li
        :class="[{ 'pointer-events-none': currentPage === allPages }]"
        class="bg-white"
      >
        <a
          href="#"
          type="button"
          @click.prevent="(currentPage = allPages), searchEmail(searchTerm)"
          class="text-cool-blacky space-x-0 hover:bg-cyan-600 px-1"
        >
          Last
        </a>
      </li>
    </ul>
  </div>
  <div v-if="data.hits" @click="showFilter = false" class="relative w-full p-2">
    <div class="relative float-left text-white w-full max-w-[53%] p-3 pr-0">
      <Datatable :emails-received="data" @display-email="displayEmail" />
    </div>
    <div
      @click="showFilter = false"
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
  <div
    :class="noEmailsMatched ? 'block' : 'hidden'"
    class="text-white text-center"
  >
    <p>Sorry, we could not find any emails with this term :(</p>
  </div>
</template>
  
<script lang="ts">
import { defineComponent, reactive, ref, toRefs } from "vue";
import Datatable from "@/views/EmailsTable/DataTable.vue";
import { emailsSearch } from "@/services/emailsAPI";
import { MatchedEmails, TermType } from "@/types/interface";

export default defineComponent({
  components: {
    Datatable,
  },
  setup() {
    let matchedEmails = reactive<{ data: MatchedEmails }>({ data: {} }),
      searchTerm = ref(""),
      limit = ref(25),
      paginatedEntries = reactive<{ paginatedPages: Array<number> }>({
        paginatedPages: [],
      }),
      currentPageReactive = reactive<{ currentPage: number }>({
        currentPage: 1,
      }),
      allPagesReactive = reactive<{ allPages: number }>({ allPages: 1 }),
      email = reactive<{ emailInside: string | undefined }>({
        emailInside: "",
      }),
      styleObject = reactive<{
        showEntries: boolean;
      }>({ showEntries: false }),
      filterDisplay = reactive<{
        showFilter: boolean;
      }>({ showFilter: false }),
      emptyResponse = reactive<{
        noEmailsMatched: boolean;
      }>({ noEmailsMatched: false });

    const searchEmail = async (search: string): Promise<void> => {
      let Term: TermType = {
        term: search,
      };

      email.emailInside = "";

      try {
        emptyResponse.noEmailsMatched = false;
        paginatedEntries.paginatedPages = [];
        let queryLimit = limit.value;
        let queryFrom = (currentPageReactive.currentPage - 1) * limit.value;

        if (currentPageReactive.currentPage === 1) {
          queryFrom = 1;
        }

        const emails = await emailsSearch(Term, queryLimit, queryFrom);

        if (typeof emails === "string") {
          throw emails;
        }

        matchedEmails.data = emails;

        calculateEntries(matchedEmails.data.total?.value);
      } catch (err) {
        paginatedEntries.paginatedPages = [];
        styleObject.showEntries = false;
        matchedEmails.data = {};
        emptyResponse.noEmailsMatched = true;
      }
    };

    const displayEmail = (body: string) => {
      email.emailInside = body;
      backToTop();
    };

    const backToTop = () => {
      document.body.scrollTop = 0;
      document.documentElement.scrollTop = 0;
    };

    const calculateEntries = (total: number | undefined) => {
      if (total == undefined) {
        return;
      }

      if (total === 0) {
        paginatedEntries.paginatedPages = [];
        styleObject.showEntries = false;
        emptyResponse.noEmailsMatched = true;
      }

      allPagesReactive.allPages = total / limit.value;
      if (allPagesReactive.allPages % limit.value != 0) {
        allPagesReactive.allPages += 1 - (allPagesReactive.allPages % 1);
      }

      for (let i = 0; i < allPagesReactive.allPages; i++) {
        paginatedEntries.paginatedPages[i] = i + 1;
      }
    };

    return {
      searchEmail,
      displayEmail,
      backToTop,
      ...toRefs(matchedEmails),
      ...toRefs(email),
      ...toRefs(allPagesReactive),
      ...toRefs(paginatedEntries),
      ...toRefs(currentPageReactive),
      ...toRefs(styleObject),
      ...toRefs(filterDisplay),
      ...toRefs(emptyResponse),
      email,
      searchTerm,
      limit,
    };
  },
});
</script>