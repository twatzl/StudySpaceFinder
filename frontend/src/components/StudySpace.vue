<template>
  <div class="study-space">
  
    <md-card>
<<<<<<< HEAD
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
=======
      <md-card-media>
        <img src="@/assets/gurula.jpg" alt="Gurula">
      </md-card-media>
>>>>>>> 9d1240ae8327ff752365b2f63368887f2945f251

      <md-card-header>
        <div class="md-title">{{ space.name }}</div>
        <div class="md-subhead">{{ space.building }}</div>
      </md-card-header>

      <md-card-expand>
        <md-card-actions md-alignment="space-between">
          <div class="info-container">
            <md-icon>access_time</md-icon>
            <span>{{ space.open }}</span>
          </div>
          <div class="info-container">
            <md-icon>people</md-icon>
            <span>{{ space.people }}</span>
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
          <md-card-content>
            <div class="tag-container">
              <md-chip v-for="tag in space.tags" :key="tag">{{ tag }}</md-chip>
            </div>
            {{ space.info }}
            </md-card-content>
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
.md-card-content {
  padding: 0 12px 12px;
}
.md-card-media + .md-card-header {
  padding-top: 16px;
}
.md-chip {
  padding: 0 8px;
  height: 28px;
  border-radius: 28px;
  line-height: 28px;
  background-color: powderblue;
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
.tag-container {
  padding: 4px 0;
}
img {
  max-height: 210px;
}
</style> 

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

interface Space {
  name: string;
  building: string;
  open: string;
  people: string;
  floor?: string;
  info?: string;
  tags?: string[];
  reservable?: boolean;
}

@Component
<<<<<<< HEAD
  export default class StudySpace extends Vue { }
    
    
=======
export default class StudySpace extends Vue {
  @Prop({
      type: Object,
      default: () => {
          return {
            name: 'Gurula',
            building: 'Exactum',
            floor: 'K',
            open: '8-19',
            people: '20',
            info: 'Gurula is the den of TKO-äly, the association of computer science students',
            tags: ['sofas', 'coffee', 'computers', 'help', 'food'],
          };
      },
  })
  private space!: Space;
}
>>>>>>> 9d1240ae8327ff752365b2f63368887f2945f251
</script>