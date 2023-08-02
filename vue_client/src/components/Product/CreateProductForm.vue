<template>
  <div class="container flex h-fit justify-center items-center">
    <el-card class="shadow-md w-3/4 my-10">
      <template #header>
        <div class="card-header bg-amber-800 text-white p-3 rounded-md">
          <span class="text-xl font-bold tracking-widest">{{
            IsUpdate ? "Update product" : "Add a new product"
          }}</span>
        </div>
      </template>
      <el-form
        :model="model"
        :rules="rules"
        ref="form"
        @submit.prevent="createProduct"
        class="w-full"
      >
        <el-form-item prop="name_product">
          <el-input
            required
            class="py-2"
            v-model="model.name_product"
            placeholder="Product name"
            clearable
            size="large"
            :maxLength="100"
            show-word-limit
          ></el-input>
        </el-form-item>
        <el-form-item prop="description_product">
          <el-input
            v-model="model.description_product"
            placeholder="Product description"
            type="textarea"
            show-word-limit
            minLength="10"
            maxLength="200"
            clearable
          ></el-input>
        </el-form-item>
        <div class="flex">
          <el-form-item
            class="w-1/2"
            label="Quantity in stock"
            prop="quantity_product"
          >
            <el-input-number
              v-model="model.quantity_product"
              :min="0"
              :max="1000"
            />
          </el-form-item>
          <el-form-item
            label="Price (TND)"
            class="w-1/2"
            required
            prop="price_product"
          >
            <el-input-number
              v-model="model.price_product"
              :min="0"
              :precision="3"
              :step="0.01"
            />
          </el-form-item>
        </div>
        <el-form-item label="Category" required prop="category_id">
          <el-select
            class="w-full"
            v-model="model.category_id"
            placeholder="Category"
          >
            <el-option
              v-for="item in options"
              :key="item.id"
              :label="item.lable"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item prop="is_available">
          <el-checkbox
            v-model="model.is_available"
            label="Available"
            size="large"
          />
        </el-form-item>
        <el-form-item prop="product_image">
          <el-upload
            class="upload-image"
            drag
            @change="setImage"
            action="#"
            :auto-upload="false"
            :http-request="undefined"
            :multiple="false"
            :limit="1"
            :show-file-list="false"
          >
            <img
              v-if="cropImg"
              class="max-h-64"
              :src="cropImg"
              alt="Product image"
            />
            <div v-else>
              <el-icon class="el-icon--upload"
                ><svg
                  xmlns="http://www.w3.org/2000/svg"
                  aria-hidden="true"
                  role="img"
                  width="48"
                  height="48"
                  preserveAspectRatio="xMidYMid meet"
                  viewBox="0 0 24 24"
                >
                  <path
                    fill="currentColor"
                    d="M19 13a1 1 0 0 0-1 1v.38l-1.48-1.48a2.79 2.79 0 0 0-3.93 0l-.7.7l-2.48-2.48a2.85 2.85 0 0 0-3.93 0L4 12.6V7a1 1 0 0 1 1-1h7a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1ZM5 20a1 1 0 0 1-1-1v-3.57l2.9-2.9a.79.79 0 0 1 1.09 0l3.17 3.17l4.3 4.3Zm13-1a.89.89 0 0 1-.18.53L13.31 15l.7-.7a.77.77 0 0 1 1.1 0L18 17.21Zm4.71-14.71l-3-3a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-3 3a1 1 0 0 0 1.42 1.42L18 4.41V10a1 1 0 0 0 2 0V4.41l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"
                  />
                </svg>
              </el-icon>
              <div class="el-upload__text">
                Drop image here or <em>click to upload</em>
              </div>
            </div>
            <template #tip>
              <div class="el-upload__tip">
                jpg/png image with a size less than 1MB
              </div>
            </template>
          </el-upload>
          <el-dialog width="75%" v-model="dialogVisible">
            <div class="flex justify-start items-center">
              <vue-cropper
                ref="cropper"
                class="w-1/2"
                :src="imgSrc"
                preview=".preview"
                v-if="imgSrc"
              />
              <div class="preview" />
              <div class="cropped-image w-1/2">
                <img
                  v-if="cropImg"
                  :src="cropImg"
                  class="w-full"
                  alt="Cropped Image"
                />
                <div v-else class="crop-placeholder" />
              </div>
            </div>
            <template #footer>
              <span class="dialog-footer">
                <el-button @click.prevent="cropImage"> Crop </el-button>
                <el-button @click.prevent="reset"> Reset </el-button>
                <el-button
                  @click.prevent="
                    imgSrc = cropImg;
                    dialogVisible = false;
                  "
                >
                  Confirm
                </el-button>
              </span>
            </template>
          </el-dialog>
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
            ><h4 class="font-bold font-xl">
              {{ IsUpdate ? "Update product" : "Create product" }}
            </h4></el-button
          >
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>
<script lang="ts">
import { defineComponent, Ref, ref } from "vue";
import { ElMessage, ElNotification } from "element-plus";

import useAuthenticationStore from "@/store/authStore";
import { useRouter } from "vue-router";
import useProductsStore, { Product } from "@/store/productsStore";

import VueCropper from "vue-cropperjs";
import "cropperjs/dist/cropper.css";

export default defineComponent({
  name: "CreateProductForm",
  components: { VueCropper },
  setup() {
    const authStore = useAuthenticationStore();
    const productStore = useProductsStore();
    const router = useRouter();
    const options = productStore.categories;
    const IsUpdate: Ref = ref<boolean>(
      router.currentRoute.value.params.id != null
    );

    const prodToUpdate: any = productStore.filterProductById(
      router.currentRoute.value.params.id as string
    );

    const model = IsUpdate.value
      ? ref({
          name_product: prodToUpdate.name_product,
          description_product: prodToUpdate.description_product,
          price_product: prodToUpdate.priceProduct,
          quantity_product: prodToUpdate.priceProduct,
          is_available: prodToUpdate.isAvailable,
          product_image: prodToUpdate.imageProduct,
          category_id: prodToUpdate.imageProduct,
        })
      : ref({
          name_product: "",
          description_product: "",
          price_product: 0.0,
          quantity_product: 0,
          is_available: false,
          product_image: "",
          category_id: null,
        });

    const dialogVisible: Ref = ref(false);
    const imgSrc: Ref = ref<string>("");
    const cropImg: Ref = ref<string>("");

    return {
      model,
      authStore,
      productStore,
      router,
      dialogVisible,
      IsUpdate,
      options,
      imgSrc,
      cropImg,
    };
  },

  data() {
    return {
      rules: {
        name_product: [
          {
            required: true,
            message: "Product name is required",
            trigger: "blur",
          },
          {
            min: 5,
            message: "Product name length should be at least 5 characters!",
            trigger: "blur",
          },
        ],
        description_product: [
          {
            required: true,
            message: "Product description is required!",
            trigger: "blur",
          },
          {
            min: 10,
            message: "Product name length should be at least 10 characters!",
            trigger: "blur",
          },
        ],
        quantity_product: [
          {
            required: true,
            message: "Product quantity is required!",
            trigger: "blur",
          },
        ],
      },
    };
  },
  methods: {
    cropImage() {
      // get image data for post processing, e.g. upload or setting image src
      const cropper: any = this.$refs.cropper;
      this.cropImg = cropper.getCroppedCanvas().toDataURL();
    },
    reset() {
      const cropper: any = this.$refs.cropper;
      cropper.reset();
    },
    setImage(file: any) {
      const isIMAGE =
        file.raw.type === "image/jpeg" || file.raw.type === "image/png";
      const isLt1M = file.size / (1024 * 1024) < 1;
      if (!isIMAGE) {
        ElMessage({
          showClose: true,
          message: "Oops, only upload jpg/png picture!",
          type: "error",
          duration: 3000,
        });
        return false;
      }
      if (!isLt1M) {
        ElMessage({
          showClose: true,
          message: "Oops, upload file size cannot exceed 1 MB!",
          type: "error",
          duration: 3000,
        });
        return false;
      }
      const reader: FileReader = new FileReader();
      reader.readAsDataURL(file.raw);
      reader.onload = (event: any) => {
        this.imgSrc = event.target.result as string;
        this.dialogVisible = true;
      };
    },

    showFileChooser() {
      const input: any = this.$refs.input;
      input.click();
    },
    async createProduct() {
      const form: any = this.$refs.form;
      let valid = await form.validate();
      if (!valid) {
        return;
      }
      this.model.product_image = this.imgSrc;
      this.productStore
        .createProduct(this.model)
        .then(() => {
          ElNotification({
            title: "Success",
            message: "Product successfully created!",
            type: "success",
          });
          this.router.push({ name: "list_products" });
        })
        .catch((err) => {
          ElNotification({
            title: "Error",
            message: err,
            type: "error",
          });
        });
    },
  },
});
</script>

<style lang="scss">
.upload-image {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
