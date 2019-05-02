<template>
  <div id="spaces">
    <Toolbar v-bind:options="options" @input="updateFilters"></Toolbar>
    <div class="md-layout md-alignment-center">
      <StudySpace v-for="item in filteredSpaces" v-bind:key="item.id" v-bind:space="item"></StudySpace>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.md-layout.md-alignment-center {
  align-items: flex-start;
}
</style>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import StudySpace from "@/components/StudySpace.vue"; // @ is an alias to /src
import Toolbar from "@/components/Toolbar.vue";
import axios from "axios";
import BackToTop from "vue-backtotop";

Vue.use(BackToTop);

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

})
export default class Spaces extends Vue {
  private spaces: Space[] = [];
  private filters: string[] = [];
  private filteredSpaces: Space[] = [];
  private options: string[] = [];

  private mounted() {
    axios.get("http://localhost:3001/spaces").then(response => {
      response.data.forEach(element => {
        const space = element as Space;
        this.spaces.push(space);
        this.addSpaceToOptions(space);
      });
    });
    this.updateSpaceSelection();
  }

  private updateFilters(value: string[]) {
    this.filters = value;
    this.updateSpaceSelection();
  }

  private updateSpaceSelection() {
    if (this.filters.length != 0) {
      let toFilter = this.spaces;
      let filtered: Space[];
      for (let filter of this.filters) {
        filtered = [];
        for (let space of toFilter) {
          if (space.name.includes(filter)) {
            filtered.push(space);
            continue;
          }
          if (space.building.includes(filter)) {
            filtered.push(space);
            continue;
          }
          if (space.tags.includes(filter)) {
            filtered.push(space);
          }
        }
        toFilter = filtered;
      }
      this.filteredSpaces = filtered;
      return;
    }
    this.filteredSpaces = this.spaces;
  }

  private addSpaceToOptions(space: Space) {
    this.options.push(space.name);
    if (!this.options.includes(space.building)) {
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
