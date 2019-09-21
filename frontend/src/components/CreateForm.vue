<template>
  <v-row justify="center">
    <v-dialog v-model="createForm" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">User Profile</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row v-for="(header, index) in headers" :key="index">
              <v-col cols="12">
                <div v-if="header.type === 'text'">
                  <v-text-field
                    :label="getInputLabel(header.text, header.required)"
                    :rules="header.rules"
                  ></v-text-field>
                </div>
                <div v-if="header.type === 'checkbox'">
                  <v-checkbox :label="getInputLabel(header.text, true)"></v-checkbox>
                </div>
                <div v-if="header.type === 'select'">
                  <v-select :items="header.items" :label="getInputLabel(header.text, true)"></v-select>
                </div>
              </v-col>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <div class="flex-grow-1"></div>
          <v-btn color="blue darken-1" text @click="closeForm()">Close</v-btn>
          <v-btn color="blue darken-1" text @click="saveForm()">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
export default {
  props: ["createForm", "headers"],

  methods: {
    saveForm() {},

    closeForm() {
      this.$emit("createForm", false);
    },

    getInputLabel(inputName, isRequired) {
      return `${isRequired ? "* " : ""}${inputName}`;
    }
  }
};
</script>