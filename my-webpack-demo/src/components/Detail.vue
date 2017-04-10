<template>
  <div class="layout-content">
    <div class="row">
      <div class="col-md-8">
        <div class="page-header"><h1>{{ Title }}</h1></div>
        <div v-html="Content"></div>
        <img src="https://drscdn.500px.org/photo/197727369/q%3D80_m%3D1000/a0b25d7c3be6bd023f9d811048a17cf9" width="30%">
        <br>
        <button class="button-primary primary" v-on:click="change">换一个</button>
      </div>

      <div class="col-md-4">
        <right></right>
      </div>
    </div>
  </div>
</template>


<script>

  import Right from './Right.vue'

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
  components: {
    Right
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

.layout-content{
      min-height: 200px;
      margin: 15px;
      overflow: hidden;
      background: #fff;
      border-radius: 4px;
}
</style>
