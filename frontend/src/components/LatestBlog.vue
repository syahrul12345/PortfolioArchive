<template>
    <v-container grid-list-md>
        <v-layout shrink>
            <v-flex xs12>
                <v-card
                class="mx-auto"
                :flat="true"
                :outline="false"
                color="rgba(51,51,51,0)"
                :href="`//${latest.link}`">
                    <v-container grid-list-md :fluid="fluid">
                        <v-layout wrap>
                            <v-flex xs5>
                                <v-card
                                class="mx-auto"
                                :flat="flat"
                                :outline="outline"
                                elevation=true>
                                    <v-img
                                    :min-height="251"
                                    :aspect-ratio="2" 
                                    :src="require(`../assets/images/blogs/preview/` + latest.blogImage)"></v-img>
                                </v-card>
                            </v-flex>
                            <v-flex xs7>
                                <v-card
                                class="mx-auto"
                                :min-height="251"
                                :max-height="251"
                                >
                                    <h1 id="blogHeader">{{latest.title}}</h1>
                                    <h2 id="blogBlurb">{{latest.blurb}}</h2>
                                </v-card>
                            </v-flex>
                        </v-layout>
                    </v-container>
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>

</template>
<script>
	import json from './../assets/blogs.json'
	export default {
		data: () => ({
			fluid:false,
			flat:false,
			outline:false,
            latest:null,
		}),
        created() {
            this.latest = json.blogs[json.blogs.length-1]
        }

	}

</script>
<style>
    #blogHeader {
        font-size:25px;
    }
    #blogBlurb {
        font-size: 20px;
    }
</style>