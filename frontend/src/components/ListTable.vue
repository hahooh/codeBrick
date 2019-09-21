<template>
  <v-card>
    <v-card-title>{{getTitle}}</v-card-title>
    <v-container>
      <div v-if="!getItemsError">
        <v-row>
          <v-col cols="12">
            <v-text-field v-model="searchTerm" label="Search" single-line hide-details></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12">
            <v-data-table
              :loading="loading"
              :loading-text="loadingMessage"
              show-select
              v-model="selectedItems"
              :search="searchTerm"
              :headers="headers"
              :items="formatItems('Y', 'N')"
              item-key="Id"
              :items-per-page="15"
            ></v-data-table>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12">
            <v-btn @click="createForm=true" raised small class="buttons">+</v-btn>
            <v-btn @click="deleteItems" raised small class="buttons">-</v-btn>
            <v-btn @click="chooseFile()" raised small class="buttons">File Upload</v-btn>
            <input id="file-input" type="file" hidden />
          </v-col>
        </v-row>
      </div>
      <div>
        <v-alert type="error" :value="getItemsError">Fetching data error. Please contact IT.</v-alert>
      </div>
      <CreateForm v-on:createForm="toggleFormDialog" :createForm="createForm" :headers="headers"></CreateForm>
    </v-container>
  </v-card>
</template>

<script>
import CreateForm from "./CreateForm";
import { mapActions, mapGetters, mapMutations } from "vuex";
import { url } from "../mixins/url";

export default {
  components: { CreateForm },

  mixins: [url],

  data: function() {
    return {
      headers: [],
      searchTerm: "",
      selectedItems: [],
      getItemsError: false,
      createForm: false,
      loading: false,
      loadingMessage: ""
    };
  },

  computed: {
    ...mapGetters("list", ["getItems"]),
    ...mapGetters("title", ["getTitle"])
  },

  methods: {
    ...mapMutations("list", ["saveItems"]),
    ...mapActions("list", ["fetchItems", "deleteItem"]),

    initItems() {
      this.startLoading("Fetching items");
      this.saveItems({ items: [] });
      this.headers = [];
      const path = this.$route.path;
      this.fetchItems({ path })
        .then(response => {
          this.headers = this.getHeadersByUrl(path);
          this.getItemsError = false;
          this.stopLoading();
        })
        .catch(error => {
          this.getItemsError = true;
          this.stopLoading();
        });
    },

    formatItems(trueValue, falseValue) {
      const newItems = [...this.getItems];
      if (newItems.length === 0) {
        return [];
      }

      const item = newItems[0];
      const itemKeys = Object.keys(item);
      const booleanKeys = itemKeys.filter(
        key => typeof item[key] === typeof true
      );

      return newItems.map(item => {
        booleanKeys.forEach(key => {
          item[key] = item[key] ? trueValue : falseValue;
        });
        return item;
      });
    },

    chooseFile() {
      document.getElementById("file-input").click();
    },

    toggleFormDialog(isOpen) {
      this.createForm = isOpen;
    },

    deleteItems() {
      this.startLoading();
      Promise.all(
        this.selectedItems.map(item => {
          return this.deleteItem({ path: this.$route.path, id: item.Id });
        })
      )
        .then(responese => {
          this.stopLoading();
        })
        .catch(error => {
          console.error("Promise.all failed.", error);
          this.stopLoading();
        });
    },

    startLoading(loadingMessage) {
      this.loadingMessage = loadingMessage;
      this.loading = true;
    },

    stopLoading() {
      this.loading = false;
    }
  },

  watch: {
    $route: "initItems"
  },

  created() {
    this.initItems();
  }
};
</script>

<style lang="scss">
.buttons {
  margin-right: 10px;
}
</style>