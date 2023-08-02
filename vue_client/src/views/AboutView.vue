<template>
  <div class="about">
    <h1>This is an about page</h1>
  </div>
</template>
<script setup lang="ts">
import useAuthenticationStore from "@/store/authStore";

const authStore = useAuthenticationStore();

await authStore.getAllUsers().then(() => {
  console.log("request is sent");
});
</script>
