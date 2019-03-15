<template>
  <div class="spaces">
    <Toolbar/>
      <StudySpace/>
      <StudySpace/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import StudySpace from '@/components/StudySpace.vue'; // @ is an alias to /src
import Toolbar from '@/components/Toolbar.vue';

@Component({
  components: {
    StudySpace, Toolbar,
  },
})
export default class Spaces extends Vue {}
//@Props() private spaces!: Space[]
</script>