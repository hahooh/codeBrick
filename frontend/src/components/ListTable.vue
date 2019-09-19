<template>
  <v-card>
    <v-card-title>{{title}}</v-card-title>
    <v-container>
      <div v-if="items.length > 0">
        <v-row>
          <v-col cols="12">
            <v-text-field v-model="searchTerm" label="Search" single-line hide-details></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12">
            <v-data-table
              show-select
              v-model="selectedItems"
              :search="searchTerm"
              :headers="headers"
              :items="getItems('Y', 'N')"
              item-key="Id"
              :items-per-page="100"
            ></v-data-table>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12">
            <v-btn @click="createForm=true" raised small class="buttons">+</v-btn>
            <v-btn raised small class="buttons">-</v-btn>
            <v-btn @click="chooseFile()" raised small class="buttons">File Upload</v-btn>
            <input id="file-input" type="file" hidden />
          </v-col>
        </v-row>
      </div>
      <div>
        <v-alert type="error" :value="getItemsError">Fetching data error. Please contact IT.</v-alert>
      </div>
      <CreateForm v-on:createForm="toggleFormDialog" :createForm="createForm"></CreateForm>
    </v-container>
  </v-card>
</template>

<script>
import CreateForm from "./CreateForm";
import { mapState, mapActions, mapGetters, mapMutations } from "vuex";
import { url } from "../mixins/url";

export default {
  components: { CreateForm },

  mixins: [url],

  computed: {
    ...mapState("list", ["items"]),
    ...mapState("title", ["title"])
  },

  methods: {
    ...mapMutations("list", ["saveItems"]),
    ...mapActions("list", ["fetchItems"]),

    initItems() {
      this.saveItems({ items: [] });
      this.headers = [];
      const path = this.$route.path;
      this.fetchItems({ path })
        .then(response => {
          this.headers = this.getHeadersByUrl(path);
          this.getItemsError = false;
        })
        .catch(error => {
          console.error(error);
          this.getItemsError = true;
        });
    },

    getItems(trueValue, falseValue) {
      const newItems = [...this.items];
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
    }
  },

  watch: {
    $route: "initItems"
  },

  data: function() {
    return {
      headers: [],
      searchTerm: "",
      selectedItems: [],
      getItemsError: false,
      createForm: false
    };
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