<template>
  <div id="spaces">
    <Toolbar v-bind:options="options"></Toolbar>
    <div class="md-layout md-alignment-center">
      <StudySpace v-for="item in spaces" v-bind:key="item.id" v-bind:space="item"></StudySpace>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.md-alignment-center {
  align-items: flex-start;
}
</style>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import StudySpace from "@/components/StudySpace.vue"; // @ is an alias to /src
import Toolbar from "@/components/Toolbar.vue";
import axios from "axios";

interface Space {
  id: number;
  name: string;
  building: string;
  open: string;
  people: string;
  reservable: boolean;
  photo?: string;
  floor?: string;
  info?: string;
  tags?: string[];
}

@Component({
  components: {
    StudySpace,
    Toolbar
  }
  // data () {
  //   return {
  //       value: [],
  //       options: this.options,
  //   }
  // },
  // methods: {
  //   addTag (newTag) {
  //       const tag = {
  //           name: newTag,
  //           code: newTag.substring(0, 2) + Math.floor((Math.random() * 10000000))
  //       }
  //       this.options.push(tag)
  //       this.value.push(tag)
  //   }
  // }
})
export default class Spaces extends Vue {
  private spaces: Space[] = [];
  private options: string[] = [];
  private mounted() {
    axios.get("http://localhost:3001/spaces").then(response => {
      response.data.forEach(element => {
        const space = element as Space;
        this.spaces.push(space);
        this.addSpaceToOptions(space);
      });
    });
  }

  private addSpaceToOptions(space: Space) {
    this.options.push(space.name);
    if(!this.options.includes(space.building)){
      this.options.push(space.building);
    }
    if (space.tags) {
      for (let tag of space.tags) {
        if (!this.options.includes(tag)) {
          this.options.push(tag);
        }
      }
    }
  }
}
</script>
