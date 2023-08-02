<template>
  <div class="container flex h-full justify-center items-center">
    <el-card class="shadow-md w-3/4">
      <template #header>
        <div class="card-header bg-amber-800 text-white p-3 rounded-md">
          <span class="text-xl font-bold tracking-widest">Authenticate</span>
        </div>
      </template>
      <el-form :model="model" :rules="rules" ref="form" @submit.prevent="login">
        <el-form-item prop="username">
          <el-input
            class="py-3 rounded-full"
            v-model="model.usernameOrEmail"
            placeholder="Username or email"
            suffix-icon="User"
            clearable
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            class="py-3 rounded-full"
            show-password
            clearable
            v-model="model.password"
            placeholder="Password"
            type="password"
          >
          </el-input>
        </el-form-item>
        <el-divider />
        <el-form-item>
          <el-button
            :loading="authStore.isLoading"
            size="large"
            color="#9C4803"
            class="w-1/2 text-white m-auto"
            native-type="submit"
            block
            ><h4 class="font-bold font-xl">Login</h4></el-button
          >
        </el-form-item>
        <div class="flex flex-col">
          <router-link to="/signup">Don't have an account ?</router-link>
          or
          <a class="forgot-password" href="https://oxfordinformatics.com/"
            >Forgot password ?</a
          >
        </div>
      </el-form>
    </el-card>
  </div>
</template>
<script lang="ts">
import { defineComponent, ref } from "vue";
import { ElNotification } from "element-plus";
import useAuthenticationStore from "@/store/authStore";
import { Router, useRouter } from "vue-router";

const showPassword = ref(false);

export default defineComponent({
  name: "LoginForm",
  setup() {
    const authStore = useAuthenticationStore();
    const router: Router = useRouter();
    return {
      showPassword,
      authStore,
      router,
    };
  },
  data() {
    return {
      model: {
        usernameOrEmail: "",
        password: "",
      },
      rules: {
        usernameOrEmail: [
          {
            required: true,
            message: "Username or email is required !",
            trigger: "blur",
          },
        ],
        password: [
          {
            required: true,
            message: "Password is required !",
            trigger: "blur",
          },
        ],
      },
    };
  },
  methods: {
    async login() {
      const form: any = this.$refs.form;
      let valid = await form.validate();
      if (!valid) {
        return;
      }
      this.authStore
        .login({
          usernameOrEmail: this.model.usernameOrEmail,
          password: this.model.password,
        })
        .then(() => {
          ElNotification({
            title: "Success",
            message: "Logged in successfully",
            type: "success",
          });
          this.router.push({ name: "home" });
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
<style lang="scss"></style>
