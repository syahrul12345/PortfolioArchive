<template>
	<v-app>
		<Header ></Header>
		<v-container grid-list-xl>
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
	import Header from '../components/Header'
	export default {
		props:[''],
		components: {
			BlogCards,
			Header
		},
		data() {
			return {
				messages:null,
				blogs:json.blogs
			}
		},
		created() {
			
		}
	}
</script>