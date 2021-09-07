<template>
  <div class="index-container">
    <navbar></navbar>
    <div class="container">
      <div class="row">
          <div class="col"></div>
        <div class="col-lg col-sm-12">
          <login-form></login-form>
        </div>
        <div class="col"></div>
      </div>
    </div>
  </div>
</template>

<script>
import LoginForm from "../components/LoginForm.vue";
import Navbar from "../components/Navbar.vue";
export default {
  components: {
    LoginForm,
    Navbar,
  },
};
</script>
