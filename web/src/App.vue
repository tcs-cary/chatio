<template>
  <div id="app">
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

#chatwindow {
  border: 2px solid black;
  margin: 50px 10px 0 10px;
  height: 83%;
  background-color: #a1a1a1;
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
