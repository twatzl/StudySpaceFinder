<template>
  <div class="md-layout-item md-medium-size-33 md-small-size-50 md-xsmall-size-100">
    <md-card>
      <md-card-media>
        <img v-if="space.photo" :src="require(`@/assets/photos/${space.photo}`)" >
         <img v-else :src="require(`@/assets/photos/library.png`)">
      </md-card-media>
      <md-card-header>
        <div class="md-title">{{ space.name }}</div>
        <div class="md-subhead" v-if="space.location2"> {{ space.building }}, {{ space.location2}}, floor {{ space.floor }} </div>
        <div class="md-subhead" v-else> {{ space.building }}, floor {{ space.floor }} </div> 
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
            <div v-if="space.reservable"> 
              <md-icon>today</md-icon>
              <a href="https://outlook.office.com/calendar/view/month">Reserve</a>
            </div>
            <div v-else> 
                non-reservable
            </div>
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
.md-card-header {
  padding: 8px;
}
.md-layout-item{
  flex: none;
}
.md-card-media + .md-card-header {
  padding-top: 8px;
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
  white-space: nowrap;
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
  location?: string;
}

@Component
export default class StudySpace extends Vue {
  @Prop({
      type: Object as () => Space,
      default: () => {
          return {
            id: 0,
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
  public space!: Space;
}
</script>