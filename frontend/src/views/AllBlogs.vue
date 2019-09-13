<template>
	<v-app style="background-color: white">
		<v-container grid-list-xl :fluid="true" style="padding-right: 7%;padding-left: 7%">
			<v-layout row wrap>
				<v-flex xs12>
					<p class="text-center"> Collection </p>
					<p class="text-center"> Blogs </p>
					<p class="text-center"> All Published Blogs & Tutorials </p>
				</v-flex>
			</v-layout>
			<v-layout row wrap>
				<v-flex v-for="blog in blogs" :key="blog.id" xs12>
					<BlogCards
					:title="blog.title" 
					:blurb="blog.blurb" 
					:url="blog.link"
					:publication="blog.publication"
					></BlogCards>
				</v-flex>
			</v-layout>
		</v-container>
	</v-app>

</template>
<script>
	import json from './../assets/blogs.json'
	import BlogCards from '../components/BlogCards.vue'
	export default {
		props:[''],
		components: {
			BlogCards,
		},
		data() {
			return {
				messages:null,
				blogs:json.blogs,
				windowWidth:null
			}
		},
		mounted() {
			this.$nextTick(function() {
				window.addEventListener('resize', this.getWindowWidth);
				window.addEventListener('resize', this.getWindowHeight);

				//Init
				this.getWindowWidth()
			})
		},
		methods: {
		    getWindowWidth() {
		      this.windowWidth = document.documentElement.clientWidth;
		    }
  		}
	}
</script>