<template>
  <div class="container flex h-full justify-center items-center">
    <el-card class="shadow-md w-3/4">
      <template #header>
        <div class="card-header bg-amber-800 text-white p-3 rounded-md">
          <span class="text-xl font-bold tracking-widest"
            >Add a new category</span
          >
        </div>
      </template>
      <el-form
        :model="model"
        :rules="rules"
        ref="categoryForm"
        @submit.prevent="createCategory"
      >
        <el-form-item :class="lableAvailable && 'is-available'" prop="lable">
          <el-input
            class="py-2"
            v-model="model.lable"
            placeholder="Lable"
            clearable
            size="large"
          ></el-input>
        </el-form-item>
        <el-form-item prop="description">
          <el-input
            v-model="model.description"
            placeholder="Category description"
            type="textarea"
            clearable
            size="large"
          ></el-input>
        </el-form-item>
        <el-divider />
        <el-form-item>
          <el-button
            :loading="productStore.isLoading"
            size="large"
            color="#9C4803"
            class="w-1/2 text-white m-auto"
            native-type="submit"
            block
            ><h4 class="font-bold font-xl">Create category</h4></el-button
          >
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>
<script setup lang="ts">
import useProductsStore from "@/store/productsStore";
import { checkLableAvailability } from "@/api/services/categories.service";
import { ref } from "vue";
import { ElNotification } from "element-plus";
import { useRouter } from "vue-router";

const productStore = useProductsStore();
const lableAvailable = ref<boolean>(false);
const categoryForm = ref<any>();
const router = useRouter();
const model = ref({
  lable: "",
  description: "",
});
const validateLable = async (rule: any, lable: string, callback: any) => {
  if (!lableAvailable.value) {
    await checkLableAvailability(model.value.lable).then((resp: boolean) => {
      lableAvailable.value = resp;
    });
  }
  if (!lableAvailable.value) {
    return callback(new Error("This Lable is taken, please try another one!"));
  }
  callback();
};
const rules = {
  lable: [
    {
      required: true,
      message: "Category lable is required",
      trigger: "blur",
    },
    {
      min: 5,
      message: "Category label length should be at least 5 characters!",
      trigger: "blur",
    },
    { validator: validateLable, trigger: "blur" },
  ],
  description: [
    {
      required: true,
      message: "Category description is required!",
      trigger: "blur",
    },
    {
      min: 10,
      message: "Category description length should be at least 10 characters!",
      trigger: "blur",
    },
  ],
};

const createCategory = async () => {
  const valid = await categoryForm.value.validate();
  if (!valid) {
    return;
  }
  await productStore
    .createCategory(model.value)
    .then(() => {
      ElNotification({
        message: "New category was added successfuly !",
        type: "success",
        duration: 3000,
      });
      router.push({ name: "list_categories" });
    })
    .catch(() =>
      ElNotification({
        message: "Oops, Something went wrong creating the new category !",
        type: "error",
        duration: 3000,
      })
    );
};
</script>
<style lang="scss"></style>
