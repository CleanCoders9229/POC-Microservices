<template>
  <div class="index-container">
    <navbar></navbar>
    <div class="container">
      <div class="row">
        <div class="col-lg col-sm-12">
          <login-form></login-form>
        </div>
        <div class="col-lg col-md-12">
          <register-form></register-form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import LoginForm from "../components/LoginForm.vue";
import Navbar from "../components/Navbar.vue";
import RegisterForm from "../components/RegisterForm.vue"
export default {
  components: {
    LoginForm,
    Navbar,
    RegisterForm
  },
};
</script>
