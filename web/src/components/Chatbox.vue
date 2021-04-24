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
</template>

<script>
import { defineComponent, computed } from 'vue';
import Message from "@/components/Message.vue";

export default defineComponent({
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
  setup(props) {
    const myMessages = computed(() => {
      return props.messages.filter(msg => msg.sender == props.username);
    })
    const otherMessages = computed(() => {
      return props.messages.filter(msg => msg.sender != props.username);
    })

    return {
      myMessages,
      otherMessages
    }
  }
});
</script>

<style lang="css" scoped>
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
