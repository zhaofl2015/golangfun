<template>
  <div id="app">
    <!--img src="./assets/logo.png"-->
    <navi></navi>
    <!--<topnavi></topnavi>-->
    <router-view></router-view>
    <bottom></bottom>
  </div>
</template>

<script>
import Right from '@/components/Right'
import Navi from '@/components/Nav'
import TopNav from '@/components/TopNav'
import Bottom from '@/components/Bottom'

export default {
  name: 'app',
  components: {
    'right': Right,
    'navi': Navi,
    'topnavi': TopNav,
    'bottom': Bottom
  }
}

</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /*text-align: center;*/
  color: #2c3e50;
  margin-top: 60px;
}
</style>
