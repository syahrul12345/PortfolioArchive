<template>
        <v-container grid-list-xl class="mx-auto" :fluid ="true">
            <v-layout wrap>
                <v-flex xs12  sm12 md4 offset-md1>
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
                <v-flex xs12 sm 12 md6>
                    <v-card
                    class="mx-auto"
                    >
                        <v-card-title>{{latest.title}}</v-card-title>
                        <v-card-text>{{latest.blurb}}</v-card-text>
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