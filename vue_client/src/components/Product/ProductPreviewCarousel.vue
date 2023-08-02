<template>
  <div class="relative max-w-3xl mx-auto text-center">
    <span
      class="absolute inset-x-0 h-px -translate-y-1/2 bg-black/10 top-1/2"
    ></span>

    <h2 class="relative inline-block px-4 text-2xl font-bold text-center">
      Preview for clothes
    </h2>
  </div>
  <swiper
    class="swiper"
    :modules="modules"
    :space-between="10"
    :slides-per-view="4"
    :centered-slides="true"
    :grab-cursor="true"
    :loop="true"
    :autoplay="{
      delay: 2500,
      disableOnInteraction: false,
    }"
    :breakpoints="{
      '640': {
        slidesPerView: 2,
        spaceBetween: 20,
      },
      '768': {
        slidesPerView: 3,
        spaceBetween: 15,
      },
      '1024': {
        slidesPerView: 5,
        spaceBetween: 10,
      },
    }"
    navigation
  >
    <swiper-slide v-for="slide in 10" :key="slide" class="slide">
      <ProductCarouselItem class="mx-5" />
    </swiper-slide>
  </swiper>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Autoplay } from "swiper";
import { Swiper, SwiperSlide } from "swiper/vue";
import "swiper/css";
import ProductCarouselItem from "./ProductCarouselItem.vue";

export default defineComponent({
  name: "ProductPreviewCarousel",
  title: "Centered slides",
  url: import.meta.url,
  components: {
    Swiper,
    SwiperSlide,
    ProductCarouselItem,
  },
  setup() {
    return {
      modules: [Autoplay],
    };
  },
});
</script>

<style lang="scss" scoped></style>
