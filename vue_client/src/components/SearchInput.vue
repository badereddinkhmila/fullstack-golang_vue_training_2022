<template>
  <el-autocomplete
    v-model="state"
    :fetch-suggestions="querySearchAsync"
    placeholder="Search Products..."
    clearable
    @select="handleSelect"
  >
    <template #suffix>
      <el-icon :size="18" class="el-input__icon" @click="handleIconClick">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          aria-hidden="true"
          role="img"
          width="1em"
          height="1em"
          preserveAspectRatio="xMidYMid meet"
          viewBox="0 0 20 20"
        >
          <path
            fill="currentColor"
            d="M11.3 2.48a3.5 3.5 0 0 0-2.6 0L2.943 4.785A1.5 1.5 0 0 0 2 6.176v7.646a1.5 1.5 0 0 0 .943 1.393L8.7 17.518a3.5 3.5 0 0 0 2.6 0l.097-.039a4.514 4.514 0 0 1-.975-.696a1.5 1.5 0 0 1-.983-.047l-.296-.12a2.656 2.656 0 0 1-.071-.026l-5.758-2.303A.5.5 0 0 1 3 13.822V6.238l6.5 2.6v2.598c.253-.49.593-.926 1-1.29V8.838l6.5-2.6v4.433c.625.773 1 1.757 1 2.829V6.176a1.5 1.5 0 0 0-.943-1.392L11.3 2.48Zm-2.228.93a2.5 2.5 0 0 1 1.857 0l5.225 2.09l-2.279.91l-6.154-2.46l1.35-.54ZM6.375 4.487l6.154 2.461L10 7.961L3.846 5.499l2.529-1.011Zm9.928 11.108a3.5 3.5 0 1 0-.707.707l2.55 2.55a.5.5 0 0 0 .708-.707l-2.55-2.55ZM16 13.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"
          />
        </svg>
      </el-icon>
    </template>
  </el-autocomplete>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";

const state = ref("");

interface LinkItem {
  value: string;
  link: string;
}

const links = ref<LinkItem[]>([]);

const loadAll = () => {
  return [
    { value: "vue", link: "https://github.com/vuejs/vue" },
    { value: "element", link: "https://github.com/ElemeFE/element" },
    { value: "cooking", link: "https://github.com/ElemeFE/cooking" },
    { value: "mint-ui", link: "https://github.com/ElemeFE/mint-ui" },
    { value: "vuex", link: "https://github.com/vuejs/vuex" },
    { value: "vue-router", link: "https://github.com/vuejs/vue-router" },
    { value: "babel", link: "https://github.com/babel/babel" },
  ];
};

let timeout: any;
const querySearchAsync = (queryString: string, cb: (arg: any) => void) => {
  const results = queryString
    ? links.value.filter(createFilter(queryString))
    : links.value;

  clearTimeout(timeout);
  timeout = setTimeout(() => {
    cb(results);
  }, 3000 * Math.random());
};
const createFilter = (queryString: string) => {
  return (restaurant: LinkItem) => {
    return (
      restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
    );
  };
};

const handleSelect = (item: LinkItem) => {
  console.log(item);
};

onMounted(() => {
  links.value = loadAll();
});
const handleIconClick = () => {
  console.log("clicked");
};
</script>
