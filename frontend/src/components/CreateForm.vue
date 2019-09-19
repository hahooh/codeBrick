<template>
  <v-row justify="center">
    <v-dialog v-model="createForm" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">User Profile</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                
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
    }
  }
};
</script>