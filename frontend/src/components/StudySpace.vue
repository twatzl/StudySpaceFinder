<template>
  <div class="study-space">
    <md-card>
      <md-card-media>
        <img src="@/assets/gurula.jpg" alt="Gurula">
      </md-card-media>

      <md-card-header>
        <div class="md-title">Gurula</div>
        <div class="md-subhead">Exactum</div>
      </md-card-header>

      <md-card-expand>
        <md-card-actions md-alignment="space-between">
          <div class="info-container">
            <md-icon>access_time</md-icon>
            <span>8-19</span>
          </div>
          <div class="info-container">
            <md-icon>people</md-icon>
            <span>20</span>
          </div>
          <div class="info-container">
            <md-icon>today</md-icon>
            <a href="https://outlook.office.com/calendar/view/month">Reserve</a>
          </div>
          <md-card-expand-trigger>
            <md-button class="md-icon-button">
              <md-icon>keyboard_arrow_down</md-icon>
            </md-button>
          </md-card-expand-trigger>
        </md-card-actions>

        <md-card-expand-content>
          <md-card-content>Gurula is the den of TKO-äly, the accosiation of computer science students.</md-card-content>
        </md-card-expand-content>
      </md-card-expand>
    </md-card>
  </div>
</template>

 <style lang="scss" scoped>
.md-card {
  width: 320px;
  margin: 4px;
  display: inline-block;
  vertical-align: top;
}
.md-card-actions {
  padding-top: 0;
}
.md-card-actions .md-button {
    margin: 0 8px;
}
.md-card-media + .md-card-header {
  padding-top: 16px;
}
.md-header {
  text-align: left;
}
.md-icon {
  margin: 8px;
}
.info-container {
  padding: 0 4px;
}
img {
  max-height: 210px;
}
</style> 

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class StudySpace extends Vue {}
</script>