<template>
  <a
    href=""
    class="relative block overflow-hidden bg-center bg-no-repeat bg-cover rounded-xl"
    style="background-image: url(https://www.hyperui.dev/photos/beach-1.jpeg)"
    ><el-tooltip
      class="box-item"
      effect="dark"
      content="Add to shopping cart"
      placement="top"
    >
      <span
        class="absolute z-10 inline-flex items-center p-2 text-xs font-semibold text-white bg-amber-900 rounded-full right-4 top-4"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          aria-hidden="true"
          role="img"
          width="24px"
          height="24px"
          preserveAspectRatio="xMidYMid meet"
          viewBox="0 0 24 24"
        >
          <path
            fill="currentColor"
            d="M11 9h2V6h3V4h-3V1h-2v3H8v2h3v3zm-4 9c-1.1 0-1.99.9-1.99 2S5.9 22 7 22s2-.9 2-2s-.9-2-2-2zm10 0c-1.1 0-1.99.9-1.99 2s.89 2 1.99 2s2-.9 2-2s-.9-2-2-2zm-8.9-5h7.45c.75 0 1.41-.41 1.75-1.03l3.86-7.01L19.42 4l-3.87 7H8.53L4.27 2H1v2h2l3.6 7.59l-1.35 2.44C4.52 15.37 5.48 17 7 17h12v-2H7l1.1-2z"
          />
        </svg>
      </span>
    </el-tooltip>
    <div class="relative p-8 pt-40 text-white bg-black bg-opacity-40">
      <p class="text-sm">Chlaka</p>
      <h5 class="text-2xl font-bold">450£</h5>
    </div>
  </a>
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "ProductCarouselItem",
});
</script>
<style lang=""></style>
