<template>
  <div class="home">
    <h1>Chatio</h1>
    <hr>
    <div id="chatbox">

    </div>
    <input type="text" placeholder="Type a Message...">
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "Home"
};
</script>

<style scoped>

input {
  position: absolute;
  bottom: 0;
  margin-bottom: 15px;
  height: 50px;
  width: 90%;
  font-size: 20px;
  left: 4.8%;
  padding: 0
}

#chatbox {
  border: 1px solid black;
  height: 75%;
  width: 90%;
  position: absolute;
  left: 50%;
  top: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
  background-color: white;
}

hr {
  border: 1px solid black;
  width: 75%;
}

h1 {
  margin: 10px 0 0 0;
  font-size: 40px;
}

.home {
  position: absolute;
  left: 50%;
  top: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
  border: 1px solid black;
  height: 75%;
  background-color: #cfcfcf;
  width: 75%;
}

</style>
