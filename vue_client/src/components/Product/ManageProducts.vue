<template>
  <el-card class="h-full w-full justify-center shadow-2xl">
    <template #header>
      <div
        class="card-header text-left m-0 bg-amber-800 text-white p-3 rounded-md"
      >
        <span class="text-xl font-bold tracking-widest">Manage products</span>
      </div>
    </template>
    <el-table
      :data="filterTableData"
      :border="true"
      rowKey="id"
      style="width: 100%"
    >
      <el-table-column label="Name" prop="name_product" width="50" />
      <el-table-column sortable label="Description" prop="description_product">
      </el-table-column>
      <el-table-column sortable label="Price" prop="price_product">
      </el-table-column>
      <el-table-column sortable label="Quantity" prop="quantity_product">
      </el-table-column>
      <el-table-column label="Image">
        <template #default="scope">
          <img :src="scope.row.product_image" class="w-20 h-20" />
        </template>
      </el-table-column>

      <el-table-column label="Actions" align="center" width="150">
        <template #default="scope">
          <el-button type="primary" @click="handleEdit(scope.$index, scope.row)"
            ><el-icon><Edit /></el-icon
          ></el-button>
          <el-divider direction="vertical" />
          <el-button
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
            ><el-icon><Delete /></el-icon
          ></el-button> </template
      ></el-table-column>
    </el-table>
  </el-card>
</template>
<script setup lang="ts">
import router from "@/router";
import useProductsStore, { Product } from "@/store/productsStore";
import { onMounted, ref } from "vue";

const productStore = useProductsStore();
const handleDelete = (index: number, row: any) => {
  console.log();
};
const handleEdit = (index: number, row: any) => {
  router.push({ name: "update_product", params: { id: row.id } });
};
onMounted(() => {
  productStore.getAllProducts();
});

const filterTableData = ref<Product[]>(productStore.products);
</script>
<style lang="scss"></style>
