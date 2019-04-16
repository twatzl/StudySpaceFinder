<template>
  <div>
    <md-toolbar class="md-primary">
      <div class="md-toolbar-row">
        <div class="md-toolbar-section-start">
          <h3 class="md-title">Study spaces</h3>
        </div>
      </div>
      <div class="md-toolbar-row">
        <md-autocomplete
          class="search"
          v-model="searchItem"
          :md-options="searchAutocompleteOptions"
          :md-open-on-focus="false"
          v-on:md-changed="emitSearchChanged"
          v-on:md-selected="selectionChanged"
          md-dense
          md-layout="box"
        >
          <label>Search...</label>

          <template slot="md-autocomplete-item" slot-scope="{ item, term }">
            <!-- Building -->
            <!-- <div v-if="item.type === 'building'">
              <md-icon>home</md-icon>
              <md-highlight-text :md-term="term">{{ item.name }}</md-highlight-text>
            </div> -->

            <!-- Floor -->
            <div v-if="item.type === 'floor'">
              <md-chip md-clickable>
                <md-highlight-text :md-term="term">floor {{ item.name }}</md-highlight-text>
              </md-chip>
            </div>

            <!-- Tag -->
            <div v-else-if="item.type === 'tag'">
              <md-chip md-clickable>
                <md-highlight-text :md-term="term">{{ item.name }}</md-highlight-text>
              </md-chip>
            </div>

            <!-- else -->
            <md-highlight-text v-else :md-term="term">{{ item.name }}</md-highlight-text>
          </template>
        </md-autocomplete>
      </div>
      <div class="md-toolbar-row">
        <div style="margin: 1em">
          <md-chip
            v-for="tag in selectedBuildings"
            md-deletable
            v-on:md-delete="unselectBuilding(tag)"
            :key="tag"
          >
            <md-icon>home</md-icon>
            {{tag}}
          </md-chip>
          <md-chip
            v-for="tag in selectedFloors"
            md-deletable
            v-on:md-delete="unselectFloor(tag)"
            :key="tag"
          >floor {{tag}}</md-chip>
          <md-chip
            v-for="tag in selectedTags"
            md-deletable
            v-on:md-delete="unselectTag(tag)"
            :key="tag"
          >{{tag}}</md-chip>
        </div>
      </div>
    </md-toolbar>
  </div>
</template>

<style lang="scss" scoped>
.md-chip {
  height: 28px;
  border-radius: 28px;
  line-height: 28px;
  background-color: powderblue;
}
.md-toolbar + .md-toolbar {
  margin: 16px 0;
}
.md-toolbar .md-field.md-inline.md-autocomplete-box {
  box-shadow: 0 3px 1px -2px rgba(0, 0, 0, 0.2), 0 2px 2px 0 rgba(0, 0, 0, 0.14),
    0 1px 5px 0 rgba(0, 0, 0, 0.12);
}
.md-toolbar .md-autocomplete.md-theme-default.md-autocomplete-box label {
  color: rgba(0, 0, 0, 0.87);
}
.search {
  max-width: 500px;
  margin-left: 4px;
}
</style>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";

@Component
export default class Toolbar extends Vue {
  private searchItem: any = null;

  @Prop({ default: [] })
  private searchAutocompleteOptions!: string[];

//   @Prop({
//     default: () => {
//       return [];
//     }
//   })
  private availableTags!: string[];

  private selectedTags: string[] = [];
  private selectedFloors: string[] = [];
  private selectedBuildings: string[] = [];

  private emitSearchChanged() {
    const tb = this;
    this.$emit("search-changed", {
      searchTerm: tb.searchItem,
      selectedBuildings: tb.selectedBuildings,
      selectedFloors: tb.selectedFloors,
      selectedTags: tb.selectedTags
    });
  }

  private unselectTag(tag: string) {
    const index = this.selectedTags.indexOf(tag, 0);
    if (index > -1) {
      this.selectedTags.splice(index, 1);
      this.emitSearchChanged();
    }
  }

  private unselectBuilding(tag: string) {
    const index = this.selectedBuildings.indexOf(tag, 0);
    if (index > -1) {
      this.selectedBuildings.splice(index, 1);
      this.emitSearchChanged();
    }
  }

  private unselectFloor(tag: string) {
    const index = this.selectedFloors.indexOf(tag, 0);
    if (index > -1) {
      this.selectedFloors.splice(index, 1);
      this.emitSearchChanged();
    }
  }

  private selectionChanged(value: any) {
    if (value.type === "tag") {
      this.searchItem = "";
      this.selectedTags.push(value.name);
    } else if (value.type == "building") {
      this.searchItem = "";
      this.selectedBuildings.push(value.name);
    } else if (value.type == "floor") {
      this.searchItem = "";
      this.selectedFloors.push(value.name);
    } else {
      this.searchItem = value.name;
    }
    // event here should be automatically emitted
  }
}
</script>