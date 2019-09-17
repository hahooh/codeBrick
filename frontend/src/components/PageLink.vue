<template>
  <router-link :to="to">{{title}}</router-link>
</template>

<script>
export default {
  props: ["to", "title"],
};
</script>