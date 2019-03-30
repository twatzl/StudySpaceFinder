<template>
  <div id="spaces">
    <Toolbar/>
    <div class="md-layout md-alignment-center">
      <StudySpace 
        v-for="item in spaces" 
        v-bind:key="item.id"
        v-bind:space="item"
      ></StudySpace>
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
import axios from 'axios';

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

axios
  .get('http://localhost:3001/spaces')
  .then(response => {
    const spaces = response.data
    console.log(spaces)
  })

@Component({
  components: {
    StudySpace, Toolbar,
  },
})

export default class Spaces extends Vue {
  private spaces: Space[] = []
  private mounted () {
    axios
      .get('http://localhost:3001/spaces')
      .then(response => {
        response.data.forEach(element => {
          const space = element as Space;
          this.spaces.push(space);
        });
      });
    console.log(this.spaces)
  }
}

</script>