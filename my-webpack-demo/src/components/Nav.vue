<template id='navi'>
    <div class="list-group">
    	<a class="list-group-item" v-link="{ path: '/'}">Home</a>
    	<a class="list-group-item" v-link="{ path: '/about'}">About</a>
    </div>
</template>

<script>
export default {
  name: 'navi'
}
</script>
