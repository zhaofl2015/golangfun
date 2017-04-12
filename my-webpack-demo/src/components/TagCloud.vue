<template id='tagcloud'>
    <div id="tagscloud">
        <h4><strong>标签云</strong></h4>
        <a href={{Url}} class={{CssType}}>{{NameCount}}</a>
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
