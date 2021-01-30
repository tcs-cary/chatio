<template>
  <div class="join">
    <h1>Pick a Username</h1>
    <input type="text" v-model="username" placeholder="Type a Username" />
    <button type="button" @click="checkUsername">Start Chatting!</button>
  </div>
</template>

<script>
export default {
  name: "Join",
  data: function() {
    return {
      username: ""
    };
  },
  methods: {
    checkUsername: function() {
      localStorage.setItem("username", this.username);
      this.$router.push("/");
    }
  }
};
</script>

<style scoped>
input {
  height: 15%;
  font-size: 15px;
  width: 50%;
  float: left;
  margin-left: 6%;
}

button {
  font-size: 20px;
  height: 18%;
  width: 30%;
}

button:hover {
  cursor: pointer;
}

.join {
  position: absolute;
  left: 50%;
  top: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
  border: 1px solid black;
  height: 20%;
  background-color: #cfcfcf;
  width: 30%;
}

html,
body {
  background-image: url("https://www.transparenttextures.com/patterns/asfalt-dark.png");
  background-color: #9c9c9c;
  font-family: "Yusei Magic", sans-serif;
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
}
</style>
