<template>
  <div class="study-space">
    <md-card>
      <md-card-header>
        <div class="md-title">Group study rooms</div>
        <div class="md-subhead">Library, Physicum</div>
      </md-card-header>

      <md-card-media-actions>
        <md-card-media>
          <img src="@/assets/library.png" alt="Gurula" style="max-height: 157.5px">
        </md-card-media>

        <md-card-actions>
          <md-button class="md-icon-button">
            <md-icon>favorite</md-icon>
          </md-button>

          <md-button class="md-icon-button">
            <md-icon>bookmark</md-icon>
          </md-button>

          <md-button class="md-icon-button">
            <md-icon>share</md-icon>
          </md-button>
        </md-card-actions>
      </md-card-media-actions>

      <md-card-expand>
        <md-card-actions md-alignment="space-between">
          <!-- <div>
            <md-button>See map</md-button>
            <md-button>Share</md-button>
          </div>-->
          <md-card-expand-trigger>
            <md-button class="md-icon-button">
              <md-icon>keyboard_arrow_down</md-icon>
            </md-button>
          </md-card-expand-trigger>
        </md-card-actions>

        <md-card-expand-content>
          <md-card-content>Pre-bookable rooms for 3+ people. Some rooms have screens or whiteboards.</md-card-content>
        </md-card-expand-content>
      </md-card-expand>
    </md-card>

    <md-card>

      <md-card-media-actions>
        <md-card-media>
          <img src="@/assets/gurula.jpg" alt="Gurula">
        </md-card-media>

        <md-card-actions>
          <md-button class="md-icon-button">
            <md-icon>favorite</md-icon>
          </md-button>

          <md-button class="md-icon-button">
            <md-icon>bookmark</md-icon>
          </md-button>

          <md-button class="md-icon-button">
            <md-icon>share</md-icon>
          </md-button>
        </md-card-actions>
      </md-card-media-actions>

      <md-card-header>
        <div class="md-title">Gurula</div>
        <div class="md-subhead">Exactum</div>
      </md-card-header>

      <md-card-expand>
        <md-card-actions md-alignment="space-between">
          <!-- <div>
            <md-button>See map</md-button>
            <md-button>Share</md-button>
          </div>-->
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
.md-header {
  text-align: left;
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