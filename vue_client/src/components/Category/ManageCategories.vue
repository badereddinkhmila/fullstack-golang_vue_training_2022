<template>
  <el-card class="h-full w-full justify-center shadow-2xl">
    <template #header>
      <div
        class="card-header text-left m-0 bg-amber-800 text-white p-3 rounded-md"
      >
        <span class="text-xl font-bold tracking-widest">Manage categories</span>
      </div>
    </template>
    <el-table
      :data="filterTableData"
      :border="true"
      rowKey="id"
      style="width: 100%"
    >
      <el-table-column label="Id" prop="id" width="50" />
      <el-table-column sortable label="Lable" prop="lable"> </el-table-column>
      <el-table-column sortable label="Description" prop="description">
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
import useProductsStore, { Category } from "@/store/productsStore";
import { onMounted, ref } from "vue";

const productStore = useProductsStore();
const handleDelete = (index: number, row: any) => {
  console.log(row.description);
};
const handleEdit = (index: number, row: any) => {
  console.log(index);
};
onMounted(() => {
  productStore.getCategories();
});

const searchDesc = ref<string>("");
const searchLab = ref<string>("");
const filterTableData = ref<Category[]>(productStore.categories);
</script>
<style lang="scss"></style>
