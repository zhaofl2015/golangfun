<template>
  <div class="hello">
    <spinner :show="loading"></spinner>
    <div id="carousel-example-generic" v-if="not_loading" class="carousel slide" data-ride="carousel">
        <!-- Indicators -->
        <ol class="carousel-indicators">
          <li data-target="#carousel-example-generic"  v-for="(blog, index) in rotate_blog" v-bind:data-slide-to="index" v-bind:class="activeNumber === index ? 'active': ''"></li>
        </ol>

        <!-- Wrapper for slides -->
        <div class="carousel-inner" role="listbox">
          <div v-for="(blog, index) in rotate_blog" v-bind:class="activeNumber === index? 'item peopleCarouselImg active': 'item peopleCarouselImg'">
            <router-link v-bind:to="blog.Detail">
              <img :src="blog.img_url" class="img-responsive center-block peopleCarouselImg">
            </router-link>
            <div class="carousel-caption">{{blog.Title}}</div>
          </div>
        </div>


      <!-- Controls -->
      <a class="left carousel-control" href="#carousel-example-generic" role="button" data-slide="prev">
        <span class="glyphicon glyphicon-chevron-left" aria-hidden="true"></span>
        <span class="sr-only">Previous</span>
      </a>
      <a class="right carousel-control" href="#carousel-example-generic" role="button" data-slide="next">
        <span class="glyphicon glyphicon-chevron-right" aria-hidden="true"></span>
        <span class="sr-only">Next</span>
      </a>
    </div>

    <div class="jumbotron" v-if="not_loading" >
      <h1>{{wall_blog.Title}}</h1>
      <p>{{wall_blog.Summary}} ...</p>
      <p><router-link :to="wall_blog.Detail"><a class="btn btn-primary btn-lg" role="button">Learn more</a></router-link></p>
    </div>


    <div class="row" v-if="not_loading" >
      <div class="col-sm-6 col-md-4" v-for="(blog,index) in window_blog">
          <div class="thumbnail imageShowT">
            <img v-bind:src="blog.img_url" class="imageShowT">
            <div class="caption">
              <h3>{{blog.Title}}</h3>
              <p>{{blog.Summary}}</p>
              <p><router-link v-bind:to="blog.Detail"><a class="btn btn-primary" role="button">查看</a></router-link>
                <a class="btn btn-default" role="button" @click="changeOne(blog.Id)">换一个</a></p>
            </div>
          </div>
      </div>
    </div>
  </div>
</template>

<script>

  import Spinner from './Spinner.vue'

export default {
  name: 'hello',

  beforeMount: function() {
    this.loadFirstPage()
  },

  components: {
    Spinner
  },

  computed: {
    not_loading() {
      return !this.loading;
    }
  },

  data: function () {
    return {
      loading: false,
      msg: 'Welcome to Your Vue.js App',
      rotate_blog: [],
      activeNumber: 0,
      window_blog: [],
      wall_blog:{
        'Title': '',
        'Summary': ''
      },
      all_ids: []
    }
  },
  methods: {
    loadFirstPage: function() {
      this.loading = true
      this.$http.get('/hello').then(
        function(response) {
          return response.json();
        }
      ).then(function(json){
        this.rotate_blog = json.rotate_blog;
        this.wall_blog = json.wall_blog;
        this.window_blog = json.window_blog;
        this.all_ids = [];
        for(var i=0;i < this.rotate_blog.length; i++) {
          this.all_ids.push(this.rotate_blog[i].Id);
        }
        for(var i=0;i < this.window_blog.length; i++) {
          this.all_ids.push(this.window_blog[i].Id);
        }
        this.all_ids.push(this.wall_blog.Id);
        this.loading = false
      });
    },
    changeOne: function(id) {
      console.log(this.all_ids);
      console.log(id);
				this.$http.post('/change-one', {ids: this.all_ids, id:id}).then(
					function(response){
						return response.json();
					}
				).then(function(json){
          for(var i=0; i<this.window_blog.length;i ++) {
            if(this.window_blog[i].Id == id) {
//              this.window_blog[i] = json; // why does not work ? todo
              this.window_blog[i].Id = json.Id;
              this.window_blog[i].Content = json.Content;
              this.window_blog[i].Title = json.Title;
              this.window_blog[i].Summary = json.Summary;
              this.window_blog[i].Detail = json.Detail;

              break;
            }
          }

          // replace all_ids info
          for(var i=0;i<this.all_ids;i ++) {
            if(this.all_ids[i] == id) {
              this.all_ids[i] = json.Id;
              break;
            }
          }
				});
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

.peopleCarouselImg img {
	  width: auto;
	  height: 500px;
	  max-height: 500px;
}
.imageShowT img {
		width: auto;
		height: 230px;
		max-height: 230px;
}
</style>
