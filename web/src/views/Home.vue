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
    <Suspense>
    <Chatbox :username="username" />
    </Suspense>

  </div>
</template>

<script>
import { defineComponent, defineAsyncComponent } from "vue";
import { useRouter } from "vue-router";
import { ref } from "vue";

export default defineComponent({
  name: "Home",
  components: {
    Chatbox: defineAsyncComponent(() => import("@/components/Chatbox.vue"))
  },
  setup() {
const router = useRouter();
const username = ref(localStorage.getItem("username"));

function changeUsername() {
  router.push("/join");
}

return {
  username,
  changeUsername
}
  }
});
</script>

<style scoped>
#usernameDisplay {
  color: black;
  font-size: 20px;
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
