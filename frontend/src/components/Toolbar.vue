<template>
  <div>
    <md-toolbar class="md-primary">
      <div class="md-toolbar-row">
        <div class="md-toolbar-section-start">
          <h3 class="md-title">Study spaces</h3>
        </div>
      </div>
      <div class="md-toolbar-row">
        <div>
          <multiselect
            v-model="value"
            :options="filterOptions"
            :multiple="true"
            :close-on-select="false"
            :clear-on-select="false"
            :preserve-search="true"
            placeholder="Search or filter by ..."
            :show-labels="false"
            :preselect-first="false"
          >
            <template slot="selection" slot-scope="{ values, search, isOpen }">
              <span
                class="multiselect__single"
                v-if="values.length &amp;&amp; !isOpen"
              >{{ values.length }} options selected</span>
            </template>
          </multiselect>
          <pre class="language-json"><code>{{ value  }}</code></pre>
        </div>
        <!-- <multiselect
          v-model="value"
          :options="options"
          :multiple="true"
          :taggable="true"
          track-by="code"
          placeholder="Search..."
          tag-placeholder="Select"
          label="name"
          @tag="addTag"
        ></multiselect>
        <pre class="language-json"><code>{{ value  }}</code></pre> -->
      </div>
    </md-toolbar>
  </div>
</template>

<style lang="scss" scoped>
.md-chip {
  height: 28px;
  border-radius: 28px;
  line-height: 28px;
  background-color: powderblue;
}
.md-toolbar + .md-toolbar {
  margin: 16px 0;
}
.md-toolbar .md-field.md-inline.md-autocomplete-box {
  box-shadow: 0 3px 1px -2px rgba(0, 0, 0, 0.2), 0 2px 2px 0 rgba(0, 0, 0, 0.14),
    0 1px 5px 0 rgba(0, 0, 0, 0.12);
}
.md-toolbar .md-autocomplete.md-theme-default.md-autocomplete-box label {
  color: rgba(0, 0, 0, 0.87);
}
.search {
  max-width: 500px;
  margin-left: 4px;
}
</style>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import Multiselect from 'vue-multiselect'; 

interface Option {
  name: string,
}

@Component
export default class Toolbar extends Vue {
  @Prop({ type: Array })
  public options!: string[];

  // props: {
  //   options: {
  //     type: string[],
  //   }
  // }
  components: {
    Multiselect,
  }
  data () {
    
    return {
      value: null,
      filterOptions: this.options,
    }
  }
  // methods: {
  //   addTag (newTag: string):void {

  //     const tag = {
  //       name: newTag,
  //       code: newTag.substring(0, 2) + Math.floor((Math.random() * 10000000))
  //     }
  //     this.options.push(tag)
  //     this.value.push(tag)
  //   }
  // }
  
}
</script>
<style src="vue-multiselect/dist/vue-multiselect.min.css"></style>