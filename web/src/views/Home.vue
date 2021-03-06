<template>
  <div class="home">
    <div class="header">
      <h2 id="usernameDisplay">Username: {{ username }}</h2>
      <h1>Chatio</h1>
      <button @click="changeUsername" id="changeusername" type="button">
        Change Username
      </button>
    </div>
    <hr />
    <Chatbox :username="username" :messages="messages" />
    <div id="messagebar">
      <input type="text" v-model="newMessage" v-on:keyup.13="createMessage" placeholder="Type a Message..." />
      <button id="sendmessage" type="button" @click="createMessage">
        Send Message
      </button>
    </div>
  </div>
</template>

<script>
import Chatbox from "@/components/Chatbox.vue";

export default {
  name: "Home",
  components: {
    Chatbox
  },
  methods: {
    changeUsername: function() {
      this.$router.push("/join");
    },
    createMessage: function() {
      if (this.newMessage.trim() == "") {
        return;
      }
      const newMsg = {
        timestamp: Date.now(),
        sender: this.username,
        body: this.newMessage
      };
      this.messages.push(newMsg);
      this.newMessage = "";
    }
  },
  data: () => {
    return {
      username: localStorage.getItem("username"),
      messages: [
        {
          timestamp: "3:18PM",
          sender: "Arul",
          body: "Hello this is a test message."
        },
        {
          timestamp: "3:19PM",
          sender: "Carson",
          body: "message 2"
        },
        {
          timestamp: "3:20PM",
          sender: "Arul",
          body: "message 3"
        }
      ],
      newMessage: ""
    };
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
