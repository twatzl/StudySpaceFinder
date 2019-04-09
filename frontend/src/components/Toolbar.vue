<template>
    <div>
        <md-toolbar>
            <div class="md-toolbar-row">
                <div class="md-toolbar-section-start">
                    <h3 class="md-title" style="flex: 1">Study spaces</h3>
                </div>

                <div>
                    <md-autocomplete
                            class="search"
                            v-model="searchItem"
                            :md-options="searchAutocompleteOptions"
                            :md-open-on-focus="false"
                            v-on:md-changed="emitSearchChanged"
                            v-on:md-selected="selectionChanged"
                            md-dense
                            md-layout="box">
                        <label>Search...</label>

                        <template slot="md-autocomplete-item" slot-scope="{ item, term }">
                            <!-- Building -->
                            <div v-if="item.type === 'building'">
                                <md-icon>home</md-icon><md-highlight-text :md-term="term">{{ item.name }}</md-highlight-text>
                            </div>

                            <!-- Tag -->
                            <div v-else-if="item.type === 'tag'">
                                <md-chip class="md-accent" md-clickable>
                                    <md-highlight-text :md-term="term">{{ item.name }}</md-highlight-text>
                                </md-chip>
                            </div>

                            <!-- else -->
                            <md-highlight-text v-else :md-term="term">{{ item.name }}</md-highlight-text>
                        </template>
                    </md-autocomplete>
                </div>


                <div class="md-toolbar-section-end">
                    <md-button class="md-icon-button">
                        <md-icon>more_vert</md-icon>
                    </md-button>
                </div>
            </div>
            <div class="md-toolbar-row">
                <div style="margin: 1em">
                    <md-chip v-for="tag in selectedTags" class="md-primary" md-deletable
                             v-on:md-delete="unselectTag(tag)" :key="tag">{{tag}}
                    </md-chip>
                </div>
            </div>
        </md-toolbar>
    </div>
</template>

<style lang="scss" scoped>
    .md-toolbar + .md-toolbar {
        margin: 16px 0;
    }

    .search {
        max-width: 500px;
    }
</style>

<script lang="ts">
    import {Component, Prop, Vue} from 'vue-property-decorator';

    @Component
    export default class StudySpace extends Vue {
        private searchItem: any = null;

        @Prop({default: []})
        private searchAutocompleteOptions!: string[];

        @Prop({
            default: () => {
                return [];
            },
        })
        private availableTags!: string[];

        private selectedTags: string[] = [];

        private emitSearchChanged() {
            const tb = this;
            this.$emit('search-changed', {searchTerm: tb.searchItem, selectedTags: tb.selectedTags});
        }

        private unselectTag(tag: string) {
            const index = this.selectedTags.indexOf(tag, 0);
            if (index > -1) {
                this.selectedTags.splice(index, 1);
                this.emitSearchChanged();
            }
        }

        private selectionChanged(value: any) {
            if (value.type === 'tag') {
                // todo
                this.searchItem = '';
                this.selectedTags.push(value.name);
            } else {
                this.searchItem = value.name;
            }
            // event here should be automatically emitted
        }
    }
</script>