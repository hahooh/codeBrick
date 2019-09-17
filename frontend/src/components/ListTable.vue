<template>
  <v-data-table :headrs="headers" :items="items" :items-per-page="100" class="elevateion-1"></v-data-table>
</template>

<script>
import { mapState, mapActions, mapGetters } from "vuex";
import { url } from "../mixins/url";

export default {
  mixins: [url],

  computed: {
    ...mapState("list", ["items"])
  },

  methods: {
    ...mapActions("list", ["fetchItems"]),
    initItems() {
      const path = this.$route.path;
      this.fetchItems({ path })
        .then(response => {
          this.headers = this.getHeadersByUrl(path);
        })
        .catch(() => {
          console.error("fetch failed");
        });
    }
  },

  watch: {
    $route: "initItems"
  },

  data: function() {
    return {
      headers: []
    };
  },

  created() {
    this.initItems();
  }
};
</script>