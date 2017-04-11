<template>
  <div class="bloglist">
    <spinner :show="loading"></spinner>
    <div class="row">
      <div class="col-md-8">
        <div class="article-single" v-for="blog in blogs">
          <article class="article">
              <div class="row">
                <header class="header">
                  <h3 class="title">
                    <label class="label label-primary">原</label>
                    <router-link v-bind:to="blog.Detail">{{blog.Title}}</router-link>
                  </h3>
                </header>
                <div v-html="blog.Summary" class="summary"></div>
                <aside class="aside clearfix">
                <!--<span class="withpadding">{{dateformat .CreateTime "2006-01-02 15:04:05"}}</span>-->
                  <a class="day moveright"><i class="fa fa-eye"></i> {{blog.ViewCount}}</a>
                  <a class="author">{{blog.Author}}</a>
                  <span class="tags">
                    <a class="tag" v-for="tag in blog.Tags">{{tag}}</a>
                  </span>
                  <router-link v-bind:to="blog.Detail" class="btn btn-primary btn-lg pull-right">查看详细</router-link>
                </aside>
              </div>
          </article>
        </div>
      </div>
      <div class="col-md-4">
        <right></right>
      </div>
    </div>
  </div>
</template>

<script>

  import Right from './Right.vue'
  import Spinner from './Spinner.vue'

  export default {
    name: 'bloglist',

    computed: {
      page() {
        return this.page;
      },
      maxPage () {
        return (this.total + this.per_page - 1)/ this.per_page;
      },
      hasMore () {
        return this.page < this.maxPage
      }
    },

    beforeMount: function() {
        this.loadBlogs(this.page)
    },

    methods: {
      loadBlogs(to = this.page, from = -1) {
        this.loading = true
        this.$http.get('/blog-list-api').then(
            function(response) {
            return response.json();
          }
        ).then(function(json){
          this.page = json.page;
          this.per_page = json.per_page;
          this.total = json.total;
          this.blogs = json.data;
          this.loading = false;
        });
      }
    },

    components: {
      Right,
      Spinner
    },
    
    data: function() {
      return {
        loading: false,
        page: 1,
        per_page: 10,
        total: 0,
        blogs: []
      }
    }
  }

</script>


<style>

  .summary {
    margin-left: 50px;
    font-size: 15px;
    color: #6199ca;
  }

  .moveright {
    margin-left: 50px;
  }
</style>
