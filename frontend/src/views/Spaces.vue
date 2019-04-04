<template>
    <div id="spaces">
        <Toolbar v-bind:searchAutocompleteOptions="autocompleteOptions" v-on:search-changed="onSearchChanged"/>
        <div class="md-layout md-alignment-center">
            <div v-for="item in filteredSpaces">
                <StudySpace
                        v-bind:key="item.id"
                        v-bind:space="item"
                ></StudySpace>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
    .md-alignment-center {
        align-items: flex-start;
    }
</style>

<script lang="ts">
    import {Component, Vue} from 'vue-property-decorator';
    import StudySpace from '@/components/StudySpace.vue'; // @ is an alias to /src
    import Toolbar from '@/components/Toolbar.vue';
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
        .then((response) => {
            const spaces = response.data;
            // console.log(spaces)
        });

    @Component({
        components: {
            StudySpace, Toolbar,
        },
    })

    export default class Spaces extends Vue {

        private static isNullOrUndefined(obj) {
            return obj == null;
        }

        private searchString: string = '';
        private spaces: Space[] = [];
        private filteredSpaces: Space[] = [];
        private autocompleteOptions: string[] = [];
        private availableTags: string[] = [];
        private selectedTags: string[] = [];

        private mounted() {
            this.availableTags = [];
            axios
                .get('http://localhost:3001/spaces')
                .then((response) => {
                    response.data.forEach((element) => {
                        const space = element as Space;
                        this.spaces.push(space);
                        this.addSpaceToSearchIndex(space);
                    });

                    this.updateFilteredSpaces();
                });
            // console.log(this.spaces)
        }

        private addSpaceToSearchIndex(space: Space) {
            this.autocompleteOptions.push(space.name + ', ' + space.building + ', ' + space.floor);
            if (!Spaces.isNullOrUndefined(space.tags)) {
                const tags = space.tags as string[];
                this.autocompleteOptions = this.autocompleteOptions.concat(tags);
                this.availableTags = this.availableTags.concat(tags);
            }
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

        private onSearchChanged(searchString: string) {
            this.searchString = searchString;
            this.updateFilteredSpaces();
        }

        private spaceMatchesSearch(space: Space) {
            // if (space == null) return true; // debug
            /*
            if (this.selectedTags.length > 0) {
                for (let sTIdx in this.selectedTags) {
                    let selectedTag = this.selectedTags[sTIdx];
                    let found = false;
                    for (let tIdx in space.tags) {
                        let tag = space.tags[tIdx];
                        if (tag == selectedTag) {
                            found = true;
                            break;
                        }
                    }

                    // do not display a study space if any of the search tags does not apply
                    if (!found) return false;
                }
            }
            */

            if (this.searchString != null && this.searchString.length > 0) {
                const lowerSearch = this.searchString.toLowerCase();

                return space.name != null && space.name.toLowerCase().includes(lowerSearch) ||
                    space.building != null && space.building.toLowerCase().includes(lowerSearch) ||
                    space.floor != null && space.floor.toLowerCase().includes(lowerSearch);
            }
            return true;
        }
    }
</script>