<template>
  <div id="message">
    <div id="info">{{ message.sender }} | {{ message.timestamp }}</div>
    <p id="text">{{ message.body }}</p>
  </div>
</template>

<script>
export default {
  name: "Message",
  props: {
    message: {
      type: Object,
      required: true
    }
  }
};
</script>

<style>
#info {
  font-weight: bold;
  margin: 0;
}

#message {
  width: 40%;
}

#text {
  background-color: green;
  margin: 0 0 0 10px;
  padding: 10px;
  border-radius: 5px;
}
</style>
