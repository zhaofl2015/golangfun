<template>
  <div>
    <div class="row">
      <div class="col-md-8">
        <div class="page-header"><h1>{{ Title }}</h1></div>
        <div v-html="Content"></div>
        <img src="https://drscdn.500px.org/photo/197727369/q%3D80_m%3D1000/a0b25d7c3be6bd023f9d811048a17cf9" width="30%">
        <br>
        <button class="button-primary primary" v-on:click="change">换一个</button>
      </div>

      <div class="col-md-4">
        <aside class="cover">
          <div class="profile">
            <hr class="divider long">
            <p class="bio">
               雾霾天戴口罩
            </p>
            <hr class="divider short">
            <div class="navigation">
              <ul class="nav">
                <li role="presentation" class="active">
                  <a href="/">首页
                  </a>
                </li>
                <li role="presentation" class="">
                  <a href="/blog-list">博客
                  </a>
                </li>
                <li role="presentation" class="">
                  <a href="/news">最新
                  </a>
                </li>
              </ul>
            </div>
          </div>
        </aside>
      </div>
    </div>
  </div>
</template>


<script>
export default {
  name: 'detail',
  data () {
    return {
      Title: 'default title',
      Content: 'default content',
      Id: 'default id',
      all_ids: []
    }
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
    });
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
</style>
