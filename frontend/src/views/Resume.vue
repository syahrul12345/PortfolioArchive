<template>
<v-app>
	<Header></Header>
	<v-container grid-list-md>
		<v-layout wrap>
			<v-flex xs-12>
				<v-card
				:outline="true"
				:hover="false"
				:text="true">
					<v-container grid-list-xs>
						<v-layout wrap>
							<v-flex xs3>
								<v-card
								:outline="true"
								:hover="true"
								:text="true"
								:max-width="256">
									<v-img :src="require('../assets/profile.png')"></v-img>
								</v-card>
								
							</v-flex>
							<v-flex xs9>
								<v-card
								:flat="true"
								:hover="false">
									<v-card-text>
										<div>About me</div>
										<p> {{personal.info}} </p>
										<p> {{personal.hobbies}}</p>
									</v-card-text>
								</v-card>
							</v-flex>
						</v-layout>
					</v-container>
				</v-card>
			</v-flex>
			<v-flex xs4>
				<SkillBox></SkillBox>
			</v-flex>
			<v-flex xs8>
				<Employment 
				v-for="employer in personal.employers" 
				:key="employer.id"
				:Employer="employer.name"
				:Title="employer.position"
				:Duration="employer.period"
				:Image="employer.image">
					
				</Employment>
			</v-flex>
		</v-layout>
	</v-container>
</v-app>
</template>
<script>
	import Header from '../components/Header'
	import json from '../assets/personal.json'
	import SkillBox from '../components/Skills'
	import Employment from '../components/Employment'
	export default {
		props:[''],
		components: {
			Header,
			SkillBox,
			Employment
		},
		data() {
			return {
				created:null,
				personal:json
			}
		},
		created() {
			console.log("CREATED RESUME PAGE")
		}
	}
</script>
<style>
</style>