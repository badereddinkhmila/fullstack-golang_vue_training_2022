<template>
  <el-container class="layout-container-demo">
    <el-header class="shadow-xl">
      <div class="h-full flex justify-between items-center">
        <router-link to="/" class="h-full flex items-center flex-no-shrink">
          <img height="54" class="" :src="require('@/assets/crown.svg')" />
          <span class="font-bold text-2xl tracking-tight pl-2"
            >MarketPlace</span
          >
        </router-link>
        <SearchInput />
        <div class="w-1/6 h-full flex items-center justify-between">
          <div class="flex items-center justify-between">
            <router-link to="/">
              <el-tooltip
                class="box-item"
                effect="dark"
                content="Home page"
                placement="bottom"
              >
                <el-icon size="24px"
                  ><svg
                    xmlns="http://www.w3.org/2000/svg"
                    aria-hidden="true"
                    role="img"
                    width="1em"
                    height="1em"
                    preserveAspectRatio="xMidYMid meet"
                    viewBox="0 0 24 24"
                  >
                    <g
                      fill="none"
                      stroke="currentColor"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                    >
                      <path
                        stroke-dasharray="15"
                        stroke-dashoffset="15"
                        d="M4.5 21.5h15"
                      >
                        <animate
                          fill="freeze"
                          attributeName="stroke-dashoffset"
                          dur="0.2s"
                          values="15;0"
                        />
                      </path>
                      <path
                        stroke-dasharray="15"
                        stroke-dashoffset="15"
                        d="M4.5 21.5V8M19.5 21.5V8"
                      >
                        <animate
                          fill="freeze"
                          attributeName="stroke-dashoffset"
                          begin="0.2s"
                          dur="0.2s"
                          values="15;0"
                        />
                      </path>
                      <path
                        stroke-dasharray="24"
                        stroke-dashoffset="24"
                        d="M9.5 21.5V12.5H14.5V21.5"
                      >
                        <animate
                          fill="freeze"
                          attributeName="stroke-dashoffset"
                          begin="0.4s"
                          dur="0.4s"
                          values="24;0"
                        />
                      </path>
                      <path
                        stroke-dasharray="30"
                        stroke-dashoffset="30"
                        stroke-width="2"
                        d="M2 10L12 2L22 10"
                      >
                        <animate
                          fill="freeze"
                          attributeName="stroke-dashoffset"
                          begin="0.5s"
                          dur="0.4s"
                          values="30;0"
                        />
                      </path>
                    </g>
                  </svg>
                </el-icon>
              </el-tooltip>
            </router-link>
            <el-divider direction="vertical" />
            <el-tooltip
              class="box-item"
              effect="dark"
              content="Notification"
              placement="bottom"
            >
              <el-icon size="24px"
                ><svg
                  xmlns="http://www.w3.org/2000/svg"
                  aria-hidden="true"
                  role="img"
                  width="1em"
                  height="1em"
                  preserveAspectRatio="xMidYMid meet"
                  viewBox="0 0 24 24"
                >
                  <g
                    fill="none"
                    stroke="currentColor"
                    stroke-linecap="round"
                    stroke-width="2"
                  >
                    <path
                      stroke-dasharray="4"
                      stroke-dashoffset="4"
                      d="M12 3V5"
                    >
                      <animate
                        fill="freeze"
                        attributeName="stroke-dashoffset"
                        dur="0.2s"
                        values="4;0"
                      />
                    </path>
                    <path
                      stroke-dasharray="28"
                      stroke-dashoffset="28"
                      d="M12 5C8.68629 5 6 7.68629 6 11L6 17C5 17 4 18 4 19H12M12 5C15.3137 5 18 7.68629 18 11L18 17C19 17 20 18 20 19H12"
                    >
                      <animate
                        fill="freeze"
                        attributeName="stroke-dashoffset"
                        begin="0.2s"
                        dur="0.4s"
                        values="28;0"
                      />
                    </path>
                    <path
                      stroke-dasharray="8"
                      stroke-dashoffset="8"
                      d="M10 20C10 21.1046 10.8954 22 12 22C13.1046 22 14 21.1046 14 20"
                    >
                      <animate
                        fill="freeze"
                        attributeName="stroke-dashoffset"
                        begin="0.6s"
                        dur="0.2s"
                        values="8;0"
                      />
                    </path>
                  </g>
                </svg>
              </el-icon>
            </el-tooltip>
            <el-divider direction="vertical" />
            <el-popover
              :width="300"
              popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;"
            >
              <template #reference>
                <span class="relative cursor-pointer">
                  <el-icon size="24px"><Goods /></el-icon>
                  <span class="top-1.5 left-2.5 absolute font-extrabold"
                    >0</span
                  >
                </span>
              </template>
              <template #default>
                <div
                  class="demo-rich-conent"
                  style="display: flex; gap: 16px; flex-direction: column"
                >
                  <el-avatar
                    :size="60"
                    src="https://avatars.githubusercontent.com/u/72015883?v=4"
                    style="margin-bottom: 8px"
                  />
                  <div>
                    <p
                      class="demo-rich-content__name"
                      style="margin: 0; font-weight: 500"
                    >
                      Element Plus
                    </p>
                    <p
                      class="demo-rich-content__mention"
                      style="
                        margin: 0;
                        font-size: 14px;
                        color: var(--el-color-info);
                      "
                    >
                      @element-plus
                    </p>
                  </div>

                  <p class="demo-rich-content__desc" style="margin: 0">
                    Element Plus, a Vue 3 based component library for
                    developers, designers and product managers
                  </p>
                </div>
              </template>
            </el-popover>
          </div>
          <el-divider direction="vertical" />
          <el-dropdown>
            <el-avatar
              src="https://cdn-icons-png.flaticon.com/512/4825/4825038.png"
            ></el-avatar>
            <template #dropdown>
              <el-dropdown-menu
                v-if="authStore.currentUser != null"
                class="p-2"
              >
                <div class="flex flex-col items-center w-full px-2">
                  <span class="text-start">
                    @{{ authStore.currentUser?.username }}</span
                  >
                  <span class="text-start">{{
                    authStore.currentUser?.email
                  }}</span>
                  <h5 class="text-start">{{ authStore.getRole }}</h5>
                </div>
                <el-dropdown-item
                  ><el-button
                    color="#9C4803"
                    class="w-full"
                    @click="authStore.logout()"
                    >Logout</el-button
                  ></el-dropdown-item
                >
              </el-dropdown-menu>
              <el-dropdown-menu
                v-if="authStore.currentUser === null"
                class="p-2"
              >
                <div class="flex justify-center items-center w-full px-2">
                  <router-link to="/login"
                    ><el-button color="#9C4803">Sign In</el-button></router-link
                  >
                  <el-divider direction="vertical" />
                  <router-link to="/signup"
                    ><el-button color="#9C4803">Sign Up</el-button></router-link
                  >
                </div>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
    </el-header>
    <el-container class="bg-gray-100">
      <AdminSidebarContent v-if="isAdmin" />
      <ProviderSidebarContent v-if="isProvider" />
      <CustomerSidebarContent v-if="isCustomer" />
      <el-main>
        <el-scrollbar>
          <Suspense>
            <template #default>
              <router-view class="h-full"></router-view>
            </template>
            <template #fallback>
              <div class="container w-full h-full">
                <img src="@/assets/loader.gif" class="m-auto" />
              </div>
            </template>
          </Suspense>
          <el-backtop :right="100" :bottom="100" />
        </el-scrollbar>
      </el-main>
    </el-container>
  </el-container>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from "vue";

import AdminSidebarContent from "@/components/AdminSidebarContent.vue";
import ProviderSidebarContent from "@/components/ProviderSidebarContent.vue";
import CustomerSidebarContent from "@/components/CustomerSidebarContent.vue";
import useAuthenticationStore from "@/store/authStore";
import SearchInput from "@/components/SearchInput.vue";
import useProductsStore from "@/store/productsStore";

const authStore = useAuthenticationStore();
const productStore = useProductsStore();

const isAdmin = computed(() => authStore.getRole === "ROLE_SUPERADMIN");
const isProvider = computed(() => authStore.getRole === "ROLE_PROVIDER");
const isCustomer = computed(
  () =>
    authStore.getRole != "ROLE_PROVIDER" &&
    authStore.getRole != "ROLE_SUPERADMIN"
);

onMounted(() => {
  productStore.getAllProducts();
  productStore.getCategories();
});
</script>

<style lang="scss" scoped>
.layout-container-demo {
  height: 100vh;
  width: 100vw;

  .el-header {
    text-align: right;
    font-size: 12px;
    top: 0;
    width: 100%;
  }
  .el-aside {
    height: 75vh;
    width: 200px;
    overflow-y: hidden;
    .el-menu {
      border: none;
    }
  }

  .el-container {
    align-items: center;
    border-right: none;
  }
  .el-main {
    padding: 5px;
    height: calc(100vh - 60px);
    overflow: hidden;
  }
  .toolbar {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    height: 100%;
    right: 20px;
  }
}
</style>
