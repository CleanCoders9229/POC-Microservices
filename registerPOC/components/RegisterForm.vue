<template>
  <div class="container">
    <div class="row">
      <form v-on:submit.prevent="submitForm">
        <div class="mb-3">
          <label for="exampleInputEmail1" class="form-label"
            >Email address</label
          >
          <input
            type="email"
            class="form-control"
            id="exampleInputEmail1"
            aria-describedby="emailHelp"
            v-model="form.email"
          />
          <div id="emailHelp" class="form-text">
            We'll never share your email with anyone else.
          </div>
          
        </div>
        <div class="mb-3">
          <label for="exampleInputFullname" class="form-label">Fullname</label>
          <input
            type="text"
            class="form-control"
            id="fullname"
            v-model="form.fullname"
          />
        </div>
        <div class="mb-3">
          <label for="exampleInputPassword1" class="form-label">Password</label>
          <input
            type="password"
            class="form-control"
            id="exampleInputPassword1"
            v-model="form.password"
          />
        </div>
        <div class="mb-3">
          <label for="exampleInputPassword1" class="form-label"
            >Confirm Password</label
          >
          <input
            type="password"
            class="form-control"
            id="exampleInputPassword1"
            v-model="form.confirmPassword"
          />
        </div>
        <div class="d-grid gap-2">
          <button type="submit" class="btn btn-primary">Submit</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "RegisterForm",
  data() {
    return {
      form: {
        email: "",
        fullname: "",
        password: "",
        confirmPassword: "",
      },
    };
  },
  methods: {
    async submitForm() {
      await axios
        .post("http://localhost:8000/user/register", this.form)
        .then((res) => {
          console.log(res);
          window.location = res.data.nextURL
        })
        .catch((error) => {
          console.log("Error: ", error);
        });
    },
  },
};
</script>

<style>
</style>