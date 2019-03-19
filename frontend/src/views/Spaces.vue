<template>
  <div class="spaces">
    <Toolbar/>
    <StudySpace/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import StudySpace from '@/components/StudySpace.vue'; // @ is an alias to /src
import Toolbar from '@/components/Toolbar.vue';
import axios from 'axios';

axios
  .get('http://localhost:3001/spaces')
  .then(response => {
    const spaces = response.data
    console.log(spaces)
  })

@Component({
  components: {
    StudySpace, Toolbar,
  }
})

export default class Spaces extends Vue {}
</script>