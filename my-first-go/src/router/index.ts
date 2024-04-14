// import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
// import HomeView from "../views/HomeView.vue";

// const routes: Array<RouteRecordRaw> = [
//   {
//     path: "/",
//     name: "home",
//     component: HomeView,
//   },
//   {
//     path: "/about",
//     name: "about",
//     // route level code-splitting
//     // this generates a separate chunk (about.[hash].js) for this route
//     // which is lazy-loaded when the route is visited.
//     component: () =>
//       import(/* webpackChunkName: "about" */ "../views/AboutView.vue"),
//   },
// ];

// const router = createRouter({
//   history: createWebHistory(process.env.BASE_URL),
//   routes,
// });
import { createApp, ref } from "vue";

// Header component
function HeaderComponent() {
  return {
    template: `
      <header class="py-6 text-2xl font-bold text-gray-800"> Mamuro Email </header>
    `,
  };
}

// Search component
function SearchComponent() {
  const searchTerm = ref("");

  function handleSearch() {
    // Implement search functionality here
    console.log("Searching for:", searchTerm.value);
  }

  return {
    searchTerm,
    handleSearch,
    template: `
      <div class="flex flex-col items-center justify-center w-full my-8">
        <input v-model="searchTerm" type="text" placeholder="Search" class="w-64 px-4 py-2 text-gray-700 border rounded shadow">
        <button @click="handleSearch" class="mt-4 px-4 py-2 bg-blue-500 text-white rounded shadow hover:bg-blue-600">Search</button>
      </div>
    `,
  };
}

// EmailTable component
function EmailTableComponent() {
  return {
    props: {
      emails: {
        type: Array,
      },
    },
    template: `
      <div class="w-full mx-auto">
        <table class="w-full">
          <thead>
            <tr>
              <th class="px-4 py-2 text-left">Subject</th>
              <th class="px-4 py-2 text-left">From</th>
              <th class="px-4 py-2 text-left">To</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="email in emails" :key="email.id">
              <td class="border px-4 py-2">{{ email.subject }}</td>
              <td class="border px-4 py-2">{{ email.from }}</td>
              <td class="border px-4 py-2">{{ email.to }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    `,
  };
}

// LoremIpsumText component
function LoremIpsumTextComponent() {
  return {
    template: `
      <div class="w-full mx-auto mt-8 text-center">
        <p class="text-gray-700">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis dapibus est et dui vehicula, a pretium nisi faucibus.</p>
      </div>
    `,
  };
}

// Create the Vue app
const app = createApp({
  setup() {
    const emails = ref([
      {
        id: 1,
        subject: "Sample Subject 1",
        from: "sender@example.com",
        to: "recipient@example.com",
      },
      {
        id: 2,
        subject: "Sample Subject 2",
        from: "sender@example.com",
        to: "recipient@example.com",
      },
      {
        id: 3,
        subject: "Sample Subject 3",
        from: "sender@example.com",
        to: "recipient@example.com",
      },
    ]);

    return { emails };
  },
});

// Mount the app
app.component("HeaderComponent", HeaderComponent());
app.component("SearchComponent", SearchComponent());
app.component("EmailTableComponent", EmailTableComponent());
app.component("LoremIpsumTextComponent", LoremIpsumTextComponent());
app.mount("#app");
