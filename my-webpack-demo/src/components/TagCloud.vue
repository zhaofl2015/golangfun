<template id='tagcloud'>
    <div id="tagscloud">
        <h4><strong>标签云</strong></h4>
        <a href="#" :class="tagitem.CssType" v-for="tagitem in all_tags">{{tagitem.NameCount}}</a>
    </div>
</template>

<script>

export default {
  name: 'tagcloud',
  beforeCreate: function() {
    this.$http.get("/blog-tags").then(
      function(response) {
        return response.json();
      }
    ).then(function(json){
      this.all_tags = json;
    });
  },
  mounted: function() {
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
      all_tags: [{"CssType": "tagc1", "NameCount": "defaut"}]
    }
  },
  methods: {
    test: function() {
      alert("haha");
    }
  }
}
</script>

<style>
  /* tagscloud */
#tagscloud {
  width:250px;
  height:260px;
  position:relative;
  font-size:12px;
  color:#333;
  margin:20px auto 0;
  text-align:center;
}
#tagscloud a{
  position:absolute;
  top:0px;
  left:0px;
  color:#333;
  font-family:Arial;
  text-decoration:none;
  margin:0 10px 15px 0;
  line-height:18px;
  text-align:center;
  font-size:12px;
  padding:1px 5px;
  display:inline-block;
  border-radius:3px;
}
#tagscloud a.tagc1{
  background:#38D66F;
  color:#fff;
}
#tagscloud a.tagc2{
  background:#A432B3;
  color:#fff;
}
#tagscloud a.tagc3{
  background:#948D10;
  color:#fff;
}
#tagscloud a:hover{
  color:#fff;
  background:#0099ff;
}
</style>
