<template>
  <div>
    <md-toolbar class="md-primary">
      <h3 class="md-title" style="flex: 1">Study spaces</h3>
      <md-button class="md-icon-button">
        <md-icon>more_vert</md-icon>
      </md-button>
    </md-toolbar>
  </div>
</template>

<style lang="scss" scoped>
  .md-toolbar + .md-toolbar {
    margin: 16px 0;
  }
</style>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class StudySpace extends Vue {}
</script>