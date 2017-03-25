<template id='navi'>
    <!--<div class="list-group">-->
    	<!--<router-link class="list-group-item" to='/'>Home</router-link>-->
    	<!--<router-link class="list-group-item" to='/detail'>Detail</router-link>-->
    <!--</div>-->
    <ul class="nav nav-tabs">
      <router-link to="/"><li role="presentation" class="active">Home</li></router-link>
      <router-link to="/detail"><li role="presentation">Detail</li></router-link>
    </ul>
</template>

<script>
export default {
  name: 'navi'
}
</script>
