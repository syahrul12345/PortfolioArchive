<template>
	<v-app style="background-color: white">
		
		<v-container 
		grid-list-xl 
		:fluid="true" 
		class="mx-auto" 
		style="padding-right: 5%;padding-left: 5%">
			<v-layout row wrap>
				<v-flex xs12>
					<p class="text-center"> Collection </p>
					<p class="text-center"> DApps </p>
					<p class="text-center"> All DApps that I've created :) </p>
				</v-flex>
			</v-layout>
			<v-layout row wrap>
				<v-flex v-for="project in Projects" :key="project.id" xs12 sm12 md12 lg3>
				<router-link :to="`/dapp/` + project.id" style="text-decoration: none">
					<ProjectCard
					:title="project.name" 
					:blurb="project.blurb" 
					:url="project.url"
					:avatarSource="project.avatarName"
					:category="project.category"
					></ProjectCard>
				</router-link>
				</v-flex>
			</v-layout>
		</v-container>
	</v-app>

</template>
<script>
	import ProjectCard from '../components/ProjectCardV2'
	import json from '../assets/data.json'
	export default {
		name: 'App',
		components: {
			ProjectCard,
		},
		data: () => ({
			"Projects":json.Projects,
			windowWidth:null
		}),
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
<style>
	li a {
    text-decoration: none;
	}
</style>