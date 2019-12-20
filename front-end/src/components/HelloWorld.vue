<template>
  <div class="hello">
    <el-form ref="form" :model="sizeForm" label-width="80px" size="mini">
      <el-form-item label="username">
        <el-input v-model="form.username"></el-input>
      </el-form-item>
      <el-form-item label="password">
        <el-input v-model="form.password"></el-input>
      </el-form-item>
      <el-form-item size="large">
        <el-button type="primary" @click="onSubmit">登陆</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
/* eslint-disable */
import axios from "axios";
const instance = axios.create({
  baseURL: "http://127.0.0.1:3000"
});
export default {
  name: "HelloWorld",
  data() {
    return {
      form: {
        username: "123",
        password: "456"
      }
    };
  },
  methods: {
    onSubmit() {
      instance.post('/login',this.form)
        .then((response) => {
          alert(response.data.msg)
        })
        .catch((error) => {
          alert(error);
        });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.hello {
  margin: auto;
  text-align: center;
}
.el-form {
  position: relative;
  margin: auto;
  width: 50%;
  padding: 10px;
}
</style>
