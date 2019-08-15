<template>
        <v-container grid-list-xl class="mx-auto" :fluid ="true">
            <v-layout shrink>
                <v-flex xs4 offset-xs1>
                    <v-card
                    :flat="flat"
                    :outline="outline"
                    elevation=true
                    >
                        <v-img
                        :aspect-ratio="2" 
                        :src="require(`../assets/images/blogs/preview/` + latest.blogImage)"></v-img>
                    </v-card>
                </v-flex>
                <v-flex xs6>
                    <v-card
                    class="mx-auto"
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
        padding-bottom:2%;
    }
</style>