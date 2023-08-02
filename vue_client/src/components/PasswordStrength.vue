<template>
  <el-progress :percentage="percentage" :color="customColors" />
</template>

<script lang="ts">
import { defineComponent, ref, Ref, toRefs, watch } from "vue";

interface Props {
  password: string;
}
export default defineComponent({
  name: "PasswordStrength",
  props: {
    password: {
      required: true,
      type: String,
    },
  },
  setup(props: Props) {
    const percentage = ref(0);

    const password: Ref<Props["password"]> = toRefs(props).password;

    const passwordSatisfied = ref({
      uppercase: false,
      lowercase: false,
      specialchars: false,
      length: false,
    });

    const customColors = [
      { color: "#bf212f", percentage: 0 },
      { color: "#f9a73e", percentage: 25 },
      { color: "#ffb347", percentage: 50 },
      { color: "#fdfd96", percentage: 75 },
      { color: "#27b376", percentage: 100 },
    ];
    const LengthStrength = (pass: string): void => {
      if (pass.length >= 10 && !passwordSatisfied.value.length) {
        console.log("length is ok");
        passwordSatisfied.value.length = true;
        percentage.value += 25;
        return;
      } else if (
        pass.length < 10 &&
        passwordSatisfied.value.length &&
        percentage.value > 25
      ) {
        console.log("length no good");
        passwordSatisfied.value.length = false;
        percentage.value -= 25;
        return;
      }
      return;
    };

    const UppercaseStrength = (pass: string): void => {
      const regx = /^(.* [A - Z]){ 3}.*$/;
      if (regx.test(pass) && !passwordSatisfied.value.uppercase) {
        console.log("uppercase is good");
        passwordSatisfied.value.uppercase = true;
        percentage.value += 25;
        return;
      } else if (
        !regx.test(pass) &&
        passwordSatisfied.value.uppercase &&
        percentage.value > 25
      ) {
        passwordSatisfied.value.uppercase = false;
        percentage.value -= 25;
        return;
      }
      return;
    };

    const LowercaseStrength = (pass: string): void => {
      const regx = /^(.* [a - z]){ 3}.*$/;
      if (regx.test(pass) && !passwordSatisfied.value.lowercase) {
        console.log("lowercase is good");
        passwordSatisfied.value.lowercase = true;
        percentage.value += 25;
        return;
      } else if (
        !regx.test(pass) &&
        passwordSatisfied.value.lowercase &&
        percentage.value > 25
      ) {
        passwordSatisfied.value.lowercase = false;
        percentage.value -= 25;
        return;
      }
      return;
    };

    const SpecialCharsStrength = (pass: string): void => {
      const regx = /^(?=.*[!@#)$%(^&*]{3,})$/;
      if (regx.test(pass) && !passwordSatisfied.value.specialchars) {
        console.log("special chars is good");
        passwordSatisfied.value.specialchars = true;
        percentage.value += 25;
        return;
      } else if (
        !regx.test(pass) &&
        passwordSatisfied.value.specialchars &&
        percentage.value > 25
      ) {
        passwordSatisfied.value.specialchars = false;
        percentage.value -= 25;
        return;
      }
      return;
    };

    watch(password, (newValue: Props["password"]) => {
      UppercaseStrength(newValue);
      LowercaseStrength(newValue);
      SpecialCharsStrength(newValue);
      LengthStrength(newValue);
    });

    return {
      percentage,
      customColors,
    };
  },
});
</script>
<style lang="scss" scoped></style>
