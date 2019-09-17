<template>
  <v-app>
    <v-content>
      <v-container>
        <v-row>
          <v-col :cols="1">
            <router-link to="/">
              <v-img height="50" position="center center" src="./assets/logo.png"></v-img>
            </router-link>
          </v-col>
          <v-col :col="11">
            <h4>{{title}}</h4>
          </v-col>
        </v-row>
        <v-row>
          <v-col :cols="2">
            <ul>
              <li v-for="(subtitle, index) in subtitles" :key="index">
                <PageLink :to="subtitle.url" :title="subtitle.title"></PageLink>
              </li>
            </ul>
          </v-col>
          <v-col :cols="10">
            <router-view />
          </v-col>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import PageLink from "./components/PageLink";
import { url } from "./mixins/url";
import { subtitles } from "./router";
import { mapState, mapMutations } from "vuex";

export default {
  name: "App",

  mixins: [url],

  computed: {
    ...mapState("title", ["title"])
  },

  components: {
    PageLink
  },

  methods: {
    ...mapMutations("title", ["setTitle"]),
    initTitle() {
      this.setTitle({
        newTitle: this.getTitleByUrl(this.$route.path)
      });
    }
  },

  watch: {
    $route: "initTitle"
  },

  created() {
    this.initTitle();
  }
};
</script>
