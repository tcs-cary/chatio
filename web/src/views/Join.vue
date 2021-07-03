<template>
  <div class="join">
    <h1>Pick a Username</h1>
    <input type="text" v-model="username" placeholder="Type a Username" />
    <button type="button" @click="checkUsername">Start Chatting!</button>
  </div>
</template>

<script>
import { ref } from "vue";
import { useRouter } from "vue-router";

export default {
  name: "Join",
  setup() {
    const router = useRouter();

    const username = ref("");

    function checkUsername() {
      localStorage.setItem("username", username.value);
      router.push("/");
    }

    return {
      username,
      checkUsername
    };
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
</style>
