<template>
  <h3 class="text-2xl font-bold font-mono text-start m-2">
    Welcome to MarketPlace
  </h3>
  <el-row :gutter="20">
    <el-col
      :sm="8"
      :lg="6"
      :xl="4"
      v-for="product in products"
      :key="product.id"
      ><ProductCard :product="product"
    /></el-col>
  </el-row>
</template>

<script lang="ts" setup>
import ProductCard from "./ProductCard.vue";
import useProductsStore from "@/store/productsStore";
import { ref } from "vue";
const productStore = useProductsStore();
const products = ref(productStore.products);
</script>
<style lang="scss"></style>
