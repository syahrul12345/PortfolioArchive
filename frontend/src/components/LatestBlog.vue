<template>
    <v-container grid-list-md style="padding-left:4%;padding-right:4%">
        <v-layout shrink>
            <v-flex xs12>
                <v-card
                class="mx-auto"
                :flat="true"
                :outline="false"
                color="rgba(51,51,51,0)"
                >
                    <v-container grid-list-md>
                        <v-layout wrap>
                            <v-flex xs4>
                                <v-card
                                class="mx-auto"
                                :flat="flat"
                                :outline="outline"
                                elevation=true>
                                    <v-img
                                    :min-height="251"
                                    :max-height="251"
                                    :aspect-ratio="2" 
                                    :src="require(`../assets/images/blogs/preview/` + latest.blogImage)"></v-img>
                                </v-card>
                            </v-flex>
                            <v-flex xs8>
                                <v-card
                                class="mx-auto"
                                :min-height="251"
                                :max-height="251"
                                >
                                    <h1 id="blogHeader">{{latest.title}}</h1>
                                    <h2 id="blogBlurb">{{latest.blurb}}</h2>
                                    <a 
                                    :href="`//${latest.link}`"
                                    >
                                        <p id="blogReadMore">Read more</p>
                                    </a>
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
        font-family: Amsi Pro Cond SemiBold,Noto Sans SC,sans-serif;
        padding-top:2%;
        padding-right:1%;
        padding-left:5%;
        font-size:25px;
    }
    #blogBlurb {
        font-family: Amsi Pro Cond SemiBold,Noto Sans SC,sans-serif;
        font-size: 20px;
        padding-top:2%;
        padding-left:5%;
    }
    #blogReadMore{
        font-family: Amsi Pro Cond SemiBold,Noto Sans SC,sans-serif;
        padding-top:2%;
        padding-left:5%;
    }
</style>