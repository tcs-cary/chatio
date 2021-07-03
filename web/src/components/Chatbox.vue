<template lang="html">
  <div id="chatbox">
    <div id="left">
      <Message
        v-for="message in myMessages"
        :key="message.timestamp"
        :message="message"
      />
    </div>
    <div id="right">
      <Message
        v-for="message in otherMessages"
        :key="message.timestamp"
        :message="message"
      />
    </div>
  </div>
  <div id="messagebar">
    <input
      type="text"
      v-model="newMessage"
      v-on:keyup.enter="createMessage"
      placeholder="Type a Message..."
    />
    <button id="sendmessage" type="button" @click="createMessage">
      Send Message
    </button>
  </div>
</template>

<script>
import { defineComponent, computed } from "vue";
import Message from "@/components/Message.vue";
// import { useConnect } from "@/p2p/useP2P.js";
import { ref } from "vue";
import { getOrCreatePeerId } from "../p2p/peer-id.js";
import createLibp2p from "../p2p/libp2p.js"
import Chat from "../p2p/chat.js"


export default defineComponent({
  name: "Chatbox",
  components: {
    Message
  },
  props: {
    username: {
      type: String,
      required: true
    }
  },
  async setup(props) {



    const newMessage = ref("");
    const messages = ref([
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
    ]);



    // const connection = await useConnect("localhost:9000");

    const peerId = await getOrCreatePeerId();

    const node = await createLibp2p(peerId);

    const client = new Chat(node, "/chatio/1.0.0");

    client.on('message', (message) => {
        console.log(message);
        messages.value.push(JSON.parse(message.data));
      })
      // Listen for peer updates
      // client.on('peer:update', ({ id, name }) => {
      //   setPeers((peers) => {
      //     const newPeers = { ...peers }
      //     newPeers[id] = { name }
      //     return newPeers
      //   })
      // })
      // // Forward stats events to the eventBus
      client.on('stats', (stats) => console.log(stats));


    // connection.onMessage((msg) => {
    //   messages.value.push(msg);
    // })

    async function createMessage() {
      if (newMessage.value.trim() == "") return;
      const time = new Date();
      const timestamp = `${time.getHours()}:${time.getMinutes()}`;
      const newMsg = {
        timestamp: timestamp,
        sender: props.username,
        body: newMessage.value
      };
      messages.value.push(newMsg);
      await client.send(JSON.stringify(newMsg));
      newMessage.value = "";
    }
    const myMessages = computed(() => {
      return messages.value.filter(msg => msg.sender == props.username);
    });
    const otherMessages = computed(() => {
      return messages.value.filter(msg => msg.sender != props.username);
    });

    return {
      myMessages,
      otherMessages,
      newMessage,
      messages,
      createMessage,
    };
  }
});
</script>

<style lang="css" scoped>

#sendmessage {
  width: 15%;
  min-width: 50px;
  height: 50px;
  padding: 0;
  font-size: 20px;
  border: 1px solid black;
  border-left: none;
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

button:hover {
  cursor: pointer;
}

h1 {
  margin: 10px 0 0 0;
  font-size: 40px;
  flex-grow: 1;
  text-align: center;
}

#chatbox {
  border: 1px solid black;
  border-bottom: none;
  flex-grow: 2;
  width: 90%;
  background-color: white;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

#left,
#right {
  display: flex;
  flex-direction: column;
  width: 40%;
}
</style>
