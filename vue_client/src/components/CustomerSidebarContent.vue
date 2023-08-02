<template>
  <el-aside class="shadow-lg bg-white">
    <div class="container flex flex-col w-full h-full p-3">
      <h4 class="text-left font-bold">
        <el-icon><PriceTag /></el-icon> Price (TND)
      </h4>
      <el-slider
        v-model="value"
        :format-tooltip="(val:number)=> `${val} TND`"
        :min="0"
        :max="3000"
        range
      />
      <div class="flex w-full justify-between">
        <el-input-number class="mx-2" v-model="value[0]" size="small" />
        <el-input-number
          class="mx-2"
          v-model="value[value.length - 1]"
          size="small"
        />
      </div>
      <el-divider />
      <h4 class="text-left font-bold">
        <el-icon size="32">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            aria-hidden="true"
            preserveAspectRatio="xMidYMid meet"
            viewBox="0 0 17 17"
          >
            <path
              fill="currentColor"
              d="M15.596 7.303a3.5 3.5 0 1 1 .707-.707l2.55 2.55a.5.5 0 0 1-.707.708l-2.55-2.55ZM16 4.5a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0Zm0 4.621V17.5a.5.5 0 0 1-.794.404L10 14.118l-5.206 3.786A.5.5 0 0 1 4 17.5v-13A2.5 2.5 0 0 1 6.5 2h3.258a4.484 4.484 0 0 0-.502 1H6.5A1.5 1.5 0 0 0 5 4.5v12.018l4.706-3.422a.5.5 0 0 1 .588 0L15 16.518V8.744c.15-.053.297-.114.44-.183l.56.56Z"
            />
          </svg>
        </el-icon>
        Category
      </h4>
      <el-checkbox-group
        class="flex flex-col"
        v-model="checkedCategory"
        :max="1"
      >
        <el-checkbox
          v-for="category in categories"
          :key="category.id"
          :label="category.lable"
          >{{ category.lable }}</el-checkbox
        >
      </el-checkbox-group>
      <el-button class="w-full mt-auto text-white" color="#9C4803"
        >Applay <el-icon><Refresh /> </el-icon
      ></el-button>
    </div>
  </el-aside>
</template>

<script lang="ts" setup>
import useProductsStore, { Category } from "@/store/productsStore";
import { ref } from "vue";
const productStore = useProductsStore();
const checkedCategory = ref<string>();
const value = ref([0, 3000]);
const categories = ref<Category[]>(productStore.categories);
</script>

<style lang="scss">
.el-aside {
  overflow: hidden !important;
}
.el-slider__bar {
  background-color: #9c4803 !important;
}
.el-slider__button {
  border-color: #9c4803 !important;
}
</style>
