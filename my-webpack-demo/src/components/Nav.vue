<template id='navi'>
  <header class="headernav">
    <nav class="inner">
      <router-link to="/">Home</router-link>
      <!--<router-link to="/detail">Detail</router-link>-->
      <router-link to="/blog-list-api">Blogs</router-link>
    </nav>
  </header>
</template>

<script>
export default {
  name: 'navi'
}
</script>

<style lang="stylus">
.headernav
  background-color #6fa8dc
  position fixed
  z-index 999
  height 55px
  top 0
  left 0
  right 0
  .inner
    max-width 800px
    box-sizing border-box
    margin 0px auto
    padding 15px 5px
  a
    color rgba(255, 255, 255, .8)
    line-height 24px
    transition color .15s ease
    display inline-block
    vertical-align middle
    font-weight 300
    letter-spacing .075em
    margin-right 1.8em
    &:hover
      color #fff
    &.router-link-active
      color #fff
      font-weight 400
    &:nth-child(6)
      margin-right 0
</style>
