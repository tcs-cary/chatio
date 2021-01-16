<template>
  <div id="app">
    <div v-if="visible" id="usernameField">
      Choose a username to Begin Chatting: <input type="text" name="" placeholder="Type a username..." value="" v-model="username" >
      <button type="button" @click="startChatting">Confirm</button>
    </div>
    <div v-if="!visible" id="usernameField">
      <p>Your username is: {{ username }}</p>
      <button type="button" @click="visible = !visible; username = ''">Change Username</button>
    </div>
    <div id="chatbox">
      <h1><u>CHAT ROOM</u></h1>
      <div id="chatwindow">

      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "App",
  data: function() {
    return {
      username: '',
      visible: true
    }
  },
  methods: {
    startChatting: function() {
      if (this.username === '') {
        return
      }
      this.visible = !this.visible;
    }
  }
};
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Yusei+Magic&display=swap');

#chatbox {
  width: 60%;
  height: 70%;
  background-color: grey;
  text-align: center;
  margin: 3% auto;
  box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;
}

#usernameField {
  text-align: center;
  margin-top: 15px;
}

#chatwindow {
  border: 2px solid black;
  margin: 50px 10px 0 10px;
  height: 83%;
  background-color: #a1a1a1;
}

button {
  border: 1px black solid;
  background-color: grey;
  font-family: 'Yusei Magic', sans-serif;
}

button:hover {
  cursor: pointer;
}

input[type='text'] {

}

html, body {
  background-image: url("https://www.transparenttextures.com/patterns/asfalt-dark.png");
  background-color: #9c9c9c;
  font-family: 'Yusei Magic', sans-serif;
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
}

#app {
  height: 100%;
  width: 100%;
}
</style>
