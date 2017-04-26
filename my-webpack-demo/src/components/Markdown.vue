<template id='tagcloudnew'>
	<div style="height: 200px; width: 200px;">	
	    <h4><strong>标签云</strong></h4>
    	<div class="tagcloud">
    	    <a href="#" :class="tagitem.CssType" v-for="tagitem in all_tags">{{tagitem.NameCount}}</a>
    	</div>
	</div>
</template>


<script>

require('TagCloud');


export default {
  name: 'tagcloudnew',
  mounted: function() {
    this.$http.get("/blog-tags").then(
      function(response) {
        return response.json();
      }
    ).then(function(json){
      this.all_tags = json;
      tagcloud();
    });
  },
  beforeCreate: function() {
    this.$http.get("/blog-tags").then(
      function(response) {
        return response.json();
      }
    ).then(function(json){
      this.all_tags = json;
    });  
  },

  data() {
    return {
      all_tags: [
        {"CssType": "tagc1", "NameCount": "defaut"}
      ]
    }
  },
  methods: {
    test: function() {
      alert("haha");
    }
  }
}

</script>

