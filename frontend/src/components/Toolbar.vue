<template>
    <div>
        <md-toolbar class="md-primary">
            <h3 class="md-title" style="flex: 1">Study spaces</h3>

            <md-autocomplete
                    class="search"
                    v-model="searchStringInternal"
                    :md-options="searchAutocompleteOptions"
                    :md-open-on-focus="false"
                    v-on:md-changed="searchStringChanged()"
                    md-layout="box">
                <label>Search...</label>
            </md-autocomplete>


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
    .search {
        max-width: 500px;
    }
</style>

<script lang="ts">
    import {Component, Prop, Vue} from 'vue-property-decorator';

    @Component
    export default class StudySpace extends Vue {
        private searchStringInternal: any = null;

        @Prop({default: []})
        private searchAutocompleteOptions!: string[];

        @Prop({default: []})
        private availableTags!: string[];

        private searchStringChanged() {
            const tb = this;
            this.$emit('search-changed', tb.searchStringInternal);
        }
    }
</script>