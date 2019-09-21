<template>
  <v-row justify="center">
    <v-dialog v-model="createForm" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">Create {{title}}</span>
        </v-card-title>
        <v-card-text>
          <small>* indicates required field</small>
          <v-container>
            <v-row v-for="(header, index) in headers" :key="index">
              <v-col cols="12">
                <div v-if="header.type === 'text'">
                  <v-text-field
                    :label="getInputLabel(header.text, header.required)"
                    :rules="header.rules"
                    @input="input => onInputChange(input, header.value)"
                  ></v-text-field>
                </div>
                <div v-if="header.type === 'checkbox'">
                  <v-checkbox
                    :label="getInputLabel(header.text, true)"
                    @change="input => onInputChange(input, header.value)"
                  ></v-checkbox>
                </div>
                <div v-if="header.type === 'select'">
                  <v-select
                    :items="header.items"
                    :label="getInputLabel(header.text, true)"
                    @change="input => onInputChange(input, header.value)"
                  ></v-select>
                </div>
              </v-col>
            </v-row>
            <v-row v-if="inputError">
              <v-col cols="12">
                <v-alert type="error">{{inputErrorMessage}}</v-alert>
              </v-col>
            </v-row>
          </v-container>
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
import { mapState } from "vuex";
import { INPUT_FIELD_CHECKBOX } from "../constants/inputFieldType";

export default {
  props: ["createForm", "headers"],

  data: function() {
    return {
      inputFields: {},
      inputError: false,
      inputErrorMessage: ""
    };
  },

  computed: {
    ...mapState("title", ["title"])
  },

  methods: {
    saveForm() {
      const invalidInput = this.validateInputs();
      if (invalidInput) {
        this.inputErrorMessage = `Invalid input ${invalidInput.text}`;
        this.inputError = true;
        return;
      }
      console.log(this.inputFields);
    },

    closeForm() {
      this.$emit("createForm", false);
    },

    getInputLabel(inputName, isRequired) {
      return `${isRequired ? "* " : ""}${inputName}`;
    },

    onInputChange(input, inputName) {
      this.inputError = false;
      this.inputFields[inputName] = input;
    },

    validateInputs() {
      const invalidInput = this.headers
        .filter(header => header.input && header.required)
        .find(requiredFieldHeader => {
          const inputRules = requiredFieldHeader.rules;
          let inputValue = this.inputFields[requiredFieldHeader.value];
          console.log(requiredFieldHeader.type);
          if (
            !inputValue &&
            requiredFieldHeader.type === INPUT_FIELD_CHECKBOX
          ) {
            this.onInputChange(false, requiredFieldHeader.value);
            inputValue = false;
          }

          if (inputValue === undefined) {
            return true;
          }

          let isValidInput = 1;
          if (inputRules) {
            isValidInput = Math.min.apply(
              null,
              inputRules.map(rule => {
                return rule(inputValue) === true ? 1 : 0;
              })
            );
          }

          return isValidInput === 0 ? true : false;
        });
      return invalidInput;
    }
  }
};
</script>