<template>
  <div class="home">
    <div class="header">
      <h2 id="usernameDisplay">Username: {{ username }}</h2>
      <h1>Chatio</h1>
      <button @click="changeUsername" id="changeusername" type="button">Change Username</button>
    </div>
    <hr>
    <div id="chatbox">
    </div>
    <div id="messagebar">
      <input type="text" placeholder="Type a Message...">
      <button id="sendmessage" type="button">Send Message</button>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "Home",
  methods: {
    changeUsername: function() {
      this.$router.push("/join");
    }
  },
  data: () => {
    return {
      username: localStorage.getItem("username")
    }
  }
};
</script>

<style scoped>

#usernameDisplay {
  color: black;
  font-size: 20px;
}

#sendmessage {
  width: 15%;
  min-width: 50px;
  height: 50px;
  padding: 0;
  font-size: 20px;
  border: 1px solid black;
  border-left: none;
}

#changeusername {
  border: 1px solid black;
  font-size: 18px;
  padding: 5px;
  margin-top: 10px;
}

button:hover {
  cursor: pointer;
}

.header {
  display: flex;
  width: 70%;
}

input {
  flex-grow: 1;
  height: 50px;
  font-size: 20px;
  padding: 0;
  border: 1px solid black;
  border-right: none;
}

#messagebar {
  display: flex;
  justify-content: center;
  padding: 0;
  width: 90%;
  margin-bottom: 18px;
}

#chatbox {
  border: 1px solid black;
  border-bottom: none;
  flex-grow: 2;
  width: 90%;
  background-color: white;
  display: flex;
}

hr {
  border: 1px solid black;
  width: 75%;
}

h1 {
  margin: 10px 0 0 0;
  font-size: 40px;
  flex-grow: 1;
  text-align: center;
}

.home {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border: 1px solid black;
  height: 80%;
  background-color: #cfcfcf;
  width: 80%;
}

</style>
