<template>
  <div id="app">
    <div id="nav">
      <!-- <router-link to="/">Home</router-link> |
      <router-link to="/join">join</router-link> -->
    </div>
    <router-view />
  </div>
</template>

<style>

* {
  box-sizing: border-box;
  color: black;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

html,
body {
  background-image: url("https://www.transparenttextures.com/patterns/asfalt-dark.png");
  background-color: #9c9c9c;
  font-family: "Yusei Magic", sans-serif;
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
}
</style>
