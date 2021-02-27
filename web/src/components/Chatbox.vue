<template lang="html">
  <div id="chatbox">
    <Message v-for="message in myMessages()" :key="message.timestamp" :message="message"/>
  </div>
</template>

<script>
import Message from "@/components/Message.vue";
export default {
  name: "Chatbox",
  components: {
    Message
  },
  props: {
    messages: {
      type: Array,
      required: true
    },
    username: {
      type: String,
      required: true
    }
  },
  computed: {
    myMessages: function() {
      return this.messages.filter(msg => msg.sender == this.username)
    },
    otherMessages: function() {
      return this.messages.filter(msg => msg.sender != this.username)
    }
  }
}
</script>

<style lang="css" scoped>
#chatbox {
  border: 1px solid black;
  border-bottom: none;
  flex-grow: 2;
  width: 90%;
  background-color: white;
  display: flex;
  flex-direction: column;
}
</style>
