<template>
  <el-aside class="shadow-lg bg-white">
    <h3 class="font-bold p-2">Marketplace</h3>
    <el-divider />
    <el-scrollbar>
      <el-menu>
        <el-sub-menu index="1">
          <template #title>
            <el-icon><Box /></el-icon>Products
          </template>
          <el-divider />
          <el-menu-item-group>
            <template #title>Group 1</template>
            <router-link to="/products/create">
              <el-menu-item index="1-1">
                <el-icon
                  ><svg
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
                      d="M19.4 9H16V5.6c0-.6-.4-.6-1-.6s-1 0-1 .6V9h-3.4c-.6 0-.6.4-.6 1s0 1 .6 1H14v3.4c0 .6.4.6 1 .6s1 0 1-.6V11h3.4c.6 0 .6-.4.6-1s0-1-.6-1zm-12 0H.6C0 9 0 9.4 0 10s0 1 .6 1h6.8c.6 0 .6-.4.6-1s0-1-.6-1zm0 5H.6c-.6 0-.6.4-.6 1s0 1 .6 1h6.8c.6 0 .6-.4.6-1s0-1-.6-1zm0-10H.6C0 4 0 4.4 0 5s0 1 .6 1h6.8C8 6 8 5.6 8 5s0-1-.6-1z"
                    />
                  </svg>
                </el-icon>
                New Product
              </el-menu-item>
            </router-link>
            <el-divider />
            <el-menu-item index="1-2">
              <el-icon
                ><svg
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
                    d="M2 3.5A1.5 1.5 0 0 1 3.5 2h1A1.5 1.5 0 0 1 6 3.5v1A1.5 1.5 0 0 1 4.5 6h-1A1.5 1.5 0 0 1 2 4.5v-1ZM3.5 3a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-1ZM2 9.5A1.5 1.5 0 0 1 3.5 8h1A1.5 1.5 0 0 1 6 9.5v1A1.5 1.5 0 0 1 4.5 12h-1A1.5 1.5 0 0 1 2 10.5v-1ZM3.5 9a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-1ZM2 15.5A1.5 1.5 0 0 1 3.5 14h1A1.5 1.5 0 0 1 6 15.5v1A1.5 1.5 0 0 1 4.5 18h-1A1.5 1.5 0 0 1 2 16.5v-1Zm1.5-.5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-1ZM8 4.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Zm0 6a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Zm0 6a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Z"
                  />
                </svg>
              </el-icon>
              All Products
            </el-menu-item>
          </el-menu-item-group>
        </el-sub-menu>
        <el-divider />
        <el-sub-menu index="2">
          <template #title>
            <el-icon
              ><svg
                xmlns="http://www.w3.org/2000/svg"
                aria-hidden="true"
                role="img"
                width="1em"
                height="1em"
                preserveAspectRatio="xMidYMid meet"
                viewBox="0 0 24 24"
              >
                <path
                  fill="currentColor"
                  d="M10 3H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zM9 9H5V5h4v4zm11-6h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zm-1 6h-4V5h4v4zm-9 4H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1zm-1 6H5v-4h4v4zm8-6c-2.206 0-4 1.794-4 4s1.794 4 4 4s4-1.794 4-4s-1.794-4-4-4zm0 6c-1.103 0-2-.897-2-2s.897-2 2-2s2 .897 2 2s-.897 2-2 2z"
                /></svg></el-icon
            >Categories
          </template>
          <el-divider />
          <el-menu-item-group>
            <template #title>Group 1</template>
            <el-divider />
            <router-link :to="{ name: 'list_categories' }">
              <el-menu-item index="2-1">
                <el-icon
                  ><svg
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
                      d="M2 3.5A1.5 1.5 0 0 1 3.5 2h1A1.5 1.5 0 0 1 6 3.5v1A1.5 1.5 0 0 1 4.5 6h-1A1.5 1.5 0 0 1 2 4.5v-1ZM3.5 3a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-1ZM2 9.5A1.5 1.5 0 0 1 3.5 8h1A1.5 1.5 0 0 1 6 9.5v1A1.5 1.5 0 0 1 4.5 12h-1A1.5 1.5 0 0 1 2 10.5v-1ZM3.5 9a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-1ZM2 15.5A1.5 1.5 0 0 1 3.5 14h1A1.5 1.5 0 0 1 6 15.5v1A1.5 1.5 0 0 1 4.5 18h-1A1.5 1.5 0 0 1 2 16.5v-1Zm1.5-.5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-1ZM8 4.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Zm0 6a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Zm0 6a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Z"
                    />
                  </svg>
                </el-icon>
                All Categories
              </el-menu-item>
            </router-link>
          </el-menu-item-group>
        </el-sub-menu>
        <el-divider />
        <el-divider />
      </el-menu>
    </el-scrollbar>
  </el-aside>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "ProviderSidebarContent",
});
</script>

<style scoped></style>
