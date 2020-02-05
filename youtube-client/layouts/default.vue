<template>
  <div id="app">
    <TheHeader/>

    <div class="columns">
      <TheSidebar/>

      <div class="column content">
        <nuxt/>
      </div>
    </div>
  </div>
</template>

<script>
  import TheSidebar from "../components/TheSidebar";
  import TheHeader from "../components/TheHeader";

  export default {
    components: {
      TheSidebar, TheHeader
    }
  }
</script>

<style scoped>
  #app {
    min-height: 100vh;
    background: #FAFAFA;
  }

  .columns {
    min-height: 100vh;
  }
</style>
