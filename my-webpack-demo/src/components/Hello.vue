<template>
  <div class="hello">
    <div id="carousel-example-generic" class="carousel slide" data-ride="carousel">
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

    <div class="jumbotron">
      <h1>{{wall_blog.Title}}</h1>
      <p>{{wall_blog.Summary}} ...</p>
      <p><router-link :to="wall_blog.Detail"><a class="btn btn-primary btn-lg" role="button">Learn more</a></router-link></p>
    </div>


    <div class="row">
      <div class="col-sm-6 col-md-4" v-for="(blog,index) in window_blog">
          <div class="thumbnail imageShowT">
            <img v-bind:src="blog.img_url" class="imageShowT">
            <div class="caption">
              <h3>{{blog.Title}}</h3>
              <p>{{blog.Summary}}</p>
              <p><router-link v-bind:to="blog.Detail"><a class="btn btn-primary" role="button">查看</a></router-link>
                <router-link to="/detail"><a class="btn btn-default" role="button">换一个</a></router-link></p>
            </div>
          </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'hello',
  mounted: function () {
//    var xhrCors = 'withCredentials' in new XMLHttpRequest();
//    if(xhrCors) {
//
//      var xhr = new XMLHttpRequest();
//      var data;
//      xhr.open('GET', 'http://127.0.0.1:8081/');
//      xhr.onload = function (e) {
//        data = JSON.parse(to.response);
//      };
//      xhr.send();
//      console.log(this);
//      this.rotate_blog = data.rotate_blog;
//      this.wall_blog = data.wall_blog;
//      this.window_blog = data.window_blog;
//    }
    this.$http.get('/hello').then(
      function(response) {
        return response.json();
      }
    ).then(function(json){
      this.rotate_blog = json.rotate_blog;
      this.wall_blog = json.wall_blog;
      this.window_blog = json.window_blog;
    });
  },
  data: function () {
    return {
      msg: 'Welcome to Your Vue.js App',
      rotate_blog: [],
      activeNumber: 0,
      window_blog: [],
      wall_blog:{
        'Title': 'default title',
        'Summary': 'default summary'
      },
    }
  },
  methods: {
    changeOne: function(event) {
				this.$http.post('/change-one', {id: this.Id}).then(
					function(response){
						return response.json();
					}
				).then(function(json){
					vm.Title = json.Title;
					vm.Content = json.Content;
					vm.Id = json.Id;
					vm.all_ids = [json.Id];
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
