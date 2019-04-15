<template>
  <div id="spaces">
    <div>
      <Toolbar
        v-bind:searchAutocompleteOptions="autocompleteOptions"
        v-on:search-changed="onSearchChanged"
      />
    </div>
    <div class="md-layout md-alignment-center">
      <div v-for="item in filteredSpaces">
        <StudySpace v-bind:key="item.id" v-bind:space="item"></StudySpace>
      </div>
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

interface AutocompleteOption {
  type: string;
  name: string;
}

@Component({
  components: {
    StudySpace,
    Toolbar
  }
})
export default class Spaces extends Vue {
  private static isNullOrUndefined(obj) {
    return obj == null;
  }
  private searchString: string = "";

  private spaces: Space[] = [];

  private filteredSpaces: Space[] = [];
  private autocompleteOptions: AutocompleteOption[] = [];
  private availableTags: string[] = [];
  private selectedTags: string[] = [];
  private selectedBuildings: string[] = [];
  private selectedFloors: string[] = [];

  private mounted() {
    this.availableTags = [];
    axios.get("http://localhost:3001/spaces").then(response => {
      response.data.forEach(element => {
        const space = element as Space;
        this.spaces.push(space);
        this.addSpaceToSearchIndex(space);
      });
      this.updateFilteredSpaces();
    });
    // console.log(this.spaces)
  }

  private addSpaceToSearchIndex(space: Space) {
    const nameType: string = "name";
    const buildingType: string = "building";
    const floorType: string = "floor";
    this.addAutocompleteOption(nameType, space.name);
    this.addAutocompleteOption(buildingType, space.building);
    this.addAutocompleteOption(floorType, space.floor);
    if (!Spaces.isNullOrUndefined(space.tags)) {
      const tags = space.tags as string[];
      this.addMultipleAutocompleteOptions("tag", tags);
      this.availableTags = this.availableTags.concat(tags);
    }
  }

  private addMultipleAutocompleteOptions(type: string, names: string[]) {
    for (const idx in names) {
      const name = names[idx];
      this.addAutocompleteOption(type, name);
    }
  }

  private addAutocompleteOption(type: string, name: string) {
    // check to not insert twice
    for (const idx in this.autocompleteOptions) {
      const opt = this.autocompleteOptions[idx];
      if (opt.name === name) {
        return;
      }
    }

    this.autocompleteOptions.push({
      type: type,
      name: name
    });
  }

  private updateFilteredSpaces() {
    const fSpaces: Space[] = [];
    for (const sName in this.spaces) {
      const s = this.spaces[sName];
      if (this.spaceMatchesSearch(s)) {
        fSpaces.push(s);
      }
    }
    this.filteredSpaces = fSpaces;
  }

  private onSearchChanged(search: {
    searchTerm: string;
    selectedTags: string[];
    selectedBuildings: string[];
    selectedFloors: string[];
  }) {
    this.searchString = search.searchTerm;
    this.selectedTags = search.selectedTags;
    this.selectedBuildings = search.selectedBuildings;
    this.selectedFloors = search.selectedFloors;
    this.updateFilteredSpaces();
  }

  private spaceMatchesSearch(space: Space) {
    // if (space == null) return true; // debug

    if (this.selectedTags.length > 0) {
      Spaces.matchAllTags(this.selectedTags, space.tags);
    }

    if (
      this.selectedBuildings.length > 0 &&
      this.selectedBuildings.indexOf(space.building) === -1
    ) {
      return false;
    }

    if (
      this.selectedFloors.length > 0 &&
      this.selectedFloors.indexOf(space.floor) === -1
    ) {
      return false;
    }

    if (this.searchString != null && this.searchString.length > 0) {
      const lowerSearch = this.searchString.toLowerCase();

      return (
        (space.name != null &&
          space.name.toLowerCase().includes(lowerSearch)) ||
        (space.building != null &&
          space.building.toLowerCase().includes(lowerSearch)) ||
        (space.floor != null && space.floor.toLowerCase().includes(lowerSearch))
      );
    }
    return true;
  }

  private static matchAllTags(
    selection: string[],
    studySpaceTags: string[]
  ): boolean {
    for (const selectedIdx in selection) {
      const selectedTag = selection[selectedIdx];
      let found = false;
      for (const tagIdx in studySpaceTags) {
        const tag = studySpaceTags[tagIdx];
        if (tag === selectedTag) {
          found = true;
          break;
        }
      }

      // do not display a study space if any of the search tags does not apply
      if (!found) {
        return false;
      }
    }
    return true;
  }
}
</script>