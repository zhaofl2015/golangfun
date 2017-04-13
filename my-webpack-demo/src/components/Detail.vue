<template>
  <div class="container">
    <div class="row">
      <div class="col-md-8 blog-main">
        <div class="blog-header">
          <h1 class="blog-title">{{Title}}</h1>
        </div>

        <div class="">
            <span class="">阅读({{ViewCount}})</span>
            <span class="">|</span>
            <span class="">评论({{CommentNum}})</span>
            <span class="">|</span>
            <span class=""><a href="#" title="编辑"><span class="glyphicon glyphicon-edit" aria-hidden="true"></span></a></span>
            <span class="">|</span>
            <span class=""><a id='btnDelete' title="删除"><span class="glyphicon glyphicon-trash" aria-hidden="true"></span></a></span>
            <button class="button-primary primary" v-on:click="change">换一个</button>
        </div>

        <div class="blog-post">
          <p class="blog-post-meta">{{CreateTime}} by {{Author}} &nbsp;&nbsp; <mark>{{VisibleText}}</mark></p>
          <span class="glyphicon glyphicon-tags" aria-hidden="true"></span>
          <span class="label label-info" v-for="tag in Tags">{{tag}}</span>
          <br/><br/>
          <div class="well" v-html="Content"></div>
        </div>

        <nav>
          <ul class="pager">
              <li><a href="" title="上一篇"><span class="glyphicon glyphicon-arrow-left" aria-hidden="true"></span></a></li>
              <!--<li>已经是第一篇了</li>-->
              <li><a href="" title="下一篇"><span class="glyphicon glyphicon-arrow-right" aria-hidden="true"></span></a></li>
              <!--<li>已经是最后一篇了</li>-->
          </ul>
        </nav>
      </div>

      <div class="col-md-3 col-sm-offset-1">
        <archive></archive>
        <tagCloud></tagCloud>
      </div>
    </div>

  </div>
</template>


<script>

  import Right from './Right.vue'
  import Archive from './Archive'
  import TagCloud from './TagCloud.vue'

export default {
  name: 'detail',
  data () {
    return {
      Title: '',
      Content: '',
      Id: '',
      all_ids: [],
      ViewCount: 0,
      Comment: [],
      CreateTime: '',
      Author: '',
      Visible: 0,
      VisibleText: '',
      Tags: []
    }
  },
  components: {
    Right,
    Archive,
    TagCloud
  },
  methods: {
    change: function(event) {
				this.$http.post('/change-one', {id: this.Id}).then(
					function(response){
						return response.json();
					}
				).then(function(json){
					this.Title = json.Title;
					this.Content = json.Content;
					this.Id = json.Id;
					this.all_ids = [json.Id];
          this.ViewCount = json.ViewCount;
          this.Comment = json.Comment;
          this.CreateTime = json.CreateTime;
          this.Author = json.Author;
          this.Visible = json.Visible;
          this.VisibleText = json.VisibleText;
          this.Tags = json.Tags;
				});
			}
  },
  mounted: function () {
    console.log(this.$route.path);
    this.$http.get("/getdetail/" + this.$route.params["id"]).then(
      function(response) {
        return response.json();
      }
    ).then(function(json){
      this.Title = json.Title;
      this.Content = json.Content;
      this.Id = json.Id;
      this.all_ids = [json.Id];
      this.ViewCount = json.ViewCount;
      this.Comment = json.Comment;
      this.CreateTime = json.CreateTime;
      this.Author = json.Author;
      this.Visible = json.Visible;
      this.VisibleText = json.VisibleText;
      this.Tags = json.Tags;
    });
  },

  computed: {
    CommentNum() {
      return this.Comment.length;
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

.layout-content{
      min-height: 200px;
      margin: 15px;
      overflow: hidden;
      background: #fff;
      border-radius: 4px;
}
</style>
