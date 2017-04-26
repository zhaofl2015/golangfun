<template>
  <div class="container full-heighta" id="editblog">
    <div class="form-inline">
      <div class="checkbox">
        <label><input v-model="show" type="checkbox">show</label>
      </div>
      <div class="checkbox">
        <label><input v-model="html" type="checkbox">html</label>
      </div>
      <div class="checkbox">
        <label><input v-model="breaks" type="checkbox">breaks</label>
      </div>
      <div class="checkbox">
        <label><input v-model="linkify" type="checkbox">linkify</label>
      </div>
      <div class="checkbox">
        <label><input v-model="emoji" type="checkbox">emoji</label>
      </div>
      <div class="checkbox">
        <label><input v-model="typographer" type="checkbox">typographer</label>
      </div>
      <div class="checkbox">
        <label><input v-model="toc" type="checkbox">toc</label>
      </div>
    </div>
    <div id="toc"></div>
    <div class="row full-heighta">
      <div class="col-xs-6 full-heighta box1">
        <textarea class="source full-heighta" v-model="source">
        </textarea>
      </div>
      <section class="col-xs-6 full-heighta box2">
        <vue-markdown class="result-html full-heighta" :watches="['show','html','breaks','linkify','emoji','typographer','toc']"
          :source="source" :show="show" :html="html" :breaks="breaks" :linkify="linkify"
          :emoji="emoji" :typographer="typographer" :toc="toc" v-on:rendered="allRight"
          v-on:toc-rendered="tocAllRight" toc-id="toc"></vue-markdown>
      </section>
    </div>
    <div class="gh-ribbon"><a href="https://github.com/zhaofl2015/golangfun" target="_blank">Fork me on GitHub</a></div>
  </div>
</template>



<script>
    require('vue-markdown')

    import VueMarkdown from 'vue-markdown'
    import "../../static/markdown/main.css"
    import "../../static/markdown/prism.css"

export default {
  name: 'editblog',
  data () {
    return {
        source: "",
        show: true,
        html: false,
        breaks: true,
        linkify: false,
        emoji: true,
        typographer: true,
        toc: false
    }
  },
  components: {
    VueMarkdown
  },
  methods: {
    allRight: function (htmlStr) {
        console.log("markdown is parsed !");
    },
    tocAllRight: function (tocHtmlStr) {
        console.log("toc is parsed :", tocHtmlStr);
    }
  },
  mounted: function () {
    console.log(this.$route.path);
  },

  computed: {
  }
}

    $('.box1').scroll(function(){
        var a=$(this).scrollTop();
        var b=$(this).scrollLeft();
        $('.box2').
        animate({'scrollTop':a,'scrollLeft':b},0)
    })
    $('.box2').scroll(function(){
        var a=$(this).scrollTop();
        var b=$(this).scrollLeft();
        $('.box1').
        animate({'scrollTop':a,'scrollLeft':b},0)
    })
</script>

<style scoped>
.full-heighta {
  height: 100%;
}

</style>