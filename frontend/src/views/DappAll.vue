<template>
	<v-app>
		<v-container grid-list-xl>
			<v-layout row wrap>
				<v-flex xs12>
					<p class="text-center"> Collection </p>
					<p class="text-center"> DApps </p>
					<p class="text-center"> All DApps that I've created :) </p>
				</v-flex>
			</v-layout>
			<v-layout row wrap>
				<v-flex v-for="project in Projects" xs3>
		          <router-link :to="`/dapp/` + project.id">
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
			ProjectCard
		},
		data: () => ({
			"Projects":json.Projects
		})
	}
</script>
<style>
	li a {
    text-decoration: none;
	}
</style>