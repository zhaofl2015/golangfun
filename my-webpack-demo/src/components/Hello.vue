<template>
  <div class="hello">
    <div id="carousel-example-generic" class="carousel slide" data-ride="carousel">
        <!-- Indicators -->
        <ol class="carousel-indicators">
          <li data-target="#carousel-example-generic"  v-for="(blog, index) in rotate_blog" v-bind:data-slide-to="index" v-bind:class="activeNumber === index ? 'active': ''"></li>
        </ol>

        <!-- Wrapper for slides -->
        <div class="carousel-inner" role="listbox">
          <div  v-for="(blog, index) in rotate_blog" v-bind:class="activeNumber === index? 'item peopleCarouselImg active': 'item peopleCarouselImg'">
            <router-link to="/detail">
              <img :src="blog.img_url" class="img-responsive center-block peopleCarouselImg">
            </router-link>
            <div class="carousel-caption">{{blog.Title}}</div>
          </div>
        </div>
    </div>

    <h1>{{ msg }}</h1>
    <h2>Essential Links</h2>
    <ul>
      <li><a href="https://vuejs.org" target="_blank">Core Docs</a></li>
      <li><a href="https://forum.vuejs.org" target="_blank">Forum</a></li>
      <li><a href="https://gitter.im/vuejs/vue" target="_blank">Gitter Chat</a></li>
      <li><a href="https://twitter.com/vuejs" target="_blank">Twitter</a></li>
      <br>
      <li><a href="http://vuejs-templates.github.io/webpack/" target="_blank">Docs for This Template</a></li>
    </ul>
    <h2>Ecosystem</h2>
    <ul>
      <li><a href="http://router.vuejs.org/" target="_blank">vue-router</a></li>
      <li><a href="http://vuex.vuejs.org/" target="_blank">vuex</a></li>
      <li><a href="http://vue-loader.vuejs.org/" target="_blank">vue-loader</a></li>
      <li><a href="https://github.com/vuejs/awesome-vue" target="_blank">awesome-vue</a></li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'hello',
  mounted: function () {
    this.$http.post('http://127.0.0.1:8081/').then(
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
      activeNumber: 0
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
