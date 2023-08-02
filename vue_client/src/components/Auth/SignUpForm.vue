<template>
  <div class="container flex h-full justify-center items-center">
    <el-card class="shadow-md w-3/4">
      <template #header>
        <div class="card-header bg-amber-800 text-white p-3 rounded-md">
          <span class="text-xl font-bold tracking-widest"
            >Create an account</span
          >
        </div>
      </template>
      <el-form
        :model="model"
        :rules="rules"
        ref="form"
        @submit.prevent="register"
      >
        <el-form-item
          :class="usernameAvailable && 'is-available'"
          prop="username"
        >
          <el-input
            required
            class="py-2"
            v-model="model.username"
            placeholder="Username"
            :prefix-icon="usernameAvailable ? 'Check' : 'User'"
            clearable
            size="large"
          ></el-input>
        </el-form-item>
        <el-form-item :class="emailAvailable && 'is-available'" prop="email">
          <el-input
            v-model="model.email"
            placeholder="Email"
            type="email"
            :prefix-icon="emailAvailable ? 'Check' : 'Message'"
            clearable
            size="large"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            class="py-2"
            v-model="model.password"
            placeholder="Password"
            show-password
            clearable
            size="large"
          >
          </el-input>
          <PasswordStrength class="w-full" :password="model.password" />
        </el-form-item>
        <el-form-item required prop="confirm_password">
          <el-input
            class="py-2"
            v-model="model.confirm_password"
            placeholder="Confirm Password"
            show-password
            clearable
            size="large"
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
            ><h4 class="font-bold font-xl">SignUp</h4></el-button
          >
        </el-form-item>
        <router-link to="/login"> Already have an account ? </router-link>
      </el-form>
    </el-card>
  </div>
</template>
<script lang="ts">
import { defineComponent, ref } from "vue";
import PasswordStrength from "../PasswordStrength.vue";
import { ElNotification } from "element-plus";

import useAuthenticationStore from "@/store/authStore";
import {
  checkEmailAvailability,
  checkUsernameAvailability,
} from "@/api/services/auth.service";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "SignUpForm",
  components: {
    PasswordStrength,
  },
  setup() {
    const authStore = useAuthenticationStore();
    const router = useRouter();
    const usernameAvailable = ref(false);
    const emailAvailable = ref(false);
    return {
      authStore,
      router,
      usernameAvailable,
      emailAvailable,
    };
  },
  data() {
    return {
      model: {
        username: "",
        email: "",
        password: "",
        confirm_password: "",
      },
      rules: {
        username: [
          {
            required: true,
            message: "Username is required",
            trigger: "blur",
          },
          {
            min: 4,
            message: "Username length should be at least 5 characters!",
            trigger: "blur",
          },
          { validator: this.availableUsername, trigger: "blur" },
        ],
        email: [
          {
            required: true,
            message: "Email is required!",
            trigger: "blur",
          },
          { validator: this.validEmail, trigger: "blur" },
          { validator: this.availableEmail, trigger: "blur" },
        ],
        password: [
          { required: true, message: "Password is required!", trigger: "blur" },
          {
            min: 5,
            message: "Password length should be at least 5 characters",
            trigger: "blur",
          },
        ],
        confirm_password: [
          { required: true, message: "Password is required!", trigger: "blur" },
          {
            validator: this.validPasswordMatch,
            trigger: "blur",
          },
        ],
      },
    };
  },
  methods: {
    validEmail(rule: any, email: string, callback: any) {
      const re =
        /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      if (re.test(email)) {
        callback();
      } else {
        return callback(new Error("Please enter a valid email !"));
      }
    },
    async availableUsername(rule: any, username: string, callback: any) {
      if (!this.usernameAvailable) {
        await checkUsernameAvailability(this.$data.model.username).then(
          (resp) => {
            this.usernameAvailable = resp;
          }
        );
      }
      if (!this.usernameAvailable) {
        return callback(
          new Error("This username is taken, please try another one!")
        );
      }
      callback();
    },
    async availableEmail(rule: any, email: string, callback: any) {
      if (!this.emailAvailable) {
        await checkEmailAvailability(this.$data.model.email).then((resp) => {
          this.emailAvailable = resp;
        });
      }
      if (!this.emailAvailable) {
        return callback(
          new Error("This email is taken, please try another one!")
        );
      }
      callback();
    },

    validPasswordMatch(rule: any, confirm_password: string, callback: any) {
      if (confirm_password === this.$data.model.password) {
        callback();
      } else {
        return callback(new Error("Passwords doesn't match !"));
      }
    },

    async register() {
      const form: any = this.$refs.form;
      let valid = await form.validate();
      if (!valid) {
        return;
      }
      this.authStore
        .register({
          username: this.model.username,
          email: this.model.email,
          password: this.model.password,
        })
        .then(() => {
          ElNotification({
            title: "Success",
            message: "Registered successfully",
            type: "success",
          });
          this.router.push({ name: "login" });
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
