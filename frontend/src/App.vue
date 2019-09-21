<template>
  <v-app>
    <v-content>
      <v-container>
        <v-row>
          <v-col :cols="2">
            <v-navigation-drawer permanent>
              <v-list-item>
                <v-list-item-content>
                  <router-link to="/">
                    <v-img
                      id="page-logo"
                      height="75"
                      width="100"
                      position="center center"
                      src="./assets/logo.png"
                    ></v-img>
                  </router-link>
                </v-list-item-content>
              </v-list-item>
              <v-divider></v-divider>
              <v-list dense nav>
                <v-list-item
                  v-for="(subtitle, index) in subtitles"
                  :key="index"
                  link
                  :class="{active: isActiveLink(subtitle.url)}"
                  :to="subtitle.url"
                >
                  <v-list-item-content>{{subtitle.title}}</v-list-item-content>
                </v-list-item>
              </v-list>
            </v-navigation-drawer>
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
import { url } from "./mixins/url";
import { subtitles } from "./router";
import { mapMutations } from "vuex";

export default {
  name: "App",

  mixins: [url],

  data: function() {
    return {
      current: 0
    };
  },

  computed: {
  },

  methods: {
    ...mapMutations("title", ["setTitle"]),

    initTitle() {
      this.setTitle({
        newTitle: this.getTitleByUrl(this.$route.path)
      });
    },

    isActiveLink(url) {
      return this.$route.path == url;
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

<style lang="scss">
#page-logo {
  margin-bottom: 10px;
}

.active {
  background-color: #eeeeee;
}
</style>
