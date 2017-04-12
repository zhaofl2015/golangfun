<template id='archive'>
  <div class="sidebar-module">
    <h4>归档</h4>
    <ol class="list-unstyled">
      <li v-for="month in months">
        <a href="#">{{month}}</a>
      </li>
    </ol>
  </div>
</template>

<script>
export default {
  name: 'archive',
  mounted: function() {
    this.$http.get("/blog-months").then(
      function(response) {
        return response.json();
      }
    ).then(function(json){
      this.months = json;
    });
  },

  data() {
    return {
      months: []
    }
  }
}
</script>
