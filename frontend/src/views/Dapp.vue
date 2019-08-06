<template>
 <v-app>
 	<!-- <h1> {{ $route.params.id }} </h1> -->
 	<v-container grid-list-md>
 		<v-layout wrap>
 			<v-flex xs12>
 				<v-card
 				class="mx-auto"
 				:hover="hover">
	 				<div id="dappHeader">
	 					<p> {{project.name}} </p>
 					</div>
 					<v-container grid-list-xs>
 						<v-layout wrap>
 							<v-flex xs7 offset-xs1>
 								<v-card
 								:hover="false"
 								:outline="true"
 								:flat="true">
 										<v-img :src="require('../assets/' + myFileName)"></v-img>
 								</v-card>
 								<v-flex xs12>
 								<v-card
 								:hover="false"
 								:outline="true"
 								:flat="true">
 									{{ project.blurb}}
 								</v-card>
 							</v-flex>
 							</v-flex>
 							<v-flex xs4>
 								<v-card
 								:hover="false"
 								:outline="true"
 								:flat="true">
 									<!-- Buttons -->
 									<v-container grid-list-xs>
	 									<v-layout justify-center wrap>
	 										<v-flex xs12 text-center>
	 											<v-card
				 								:hover="true"
				 								:outline="true"
				 								:flat="false"
				 								:text="true"
				 								text-center
				 								:href="`//`+project.website">
				 									Visit the website!
				 								</v-card>
	 										</v-flex>
	 										<v-flex xs12 text-center>
	 											<v-card
				 								:hover="true"
				 								:outline="true"
				 								:flat="false"
				 								:text="true"
				 								text-center
				 								:href="`//`+project.github">
				 									Github
				 								</v-card>
	 										</v-flex>
	 										<v-flex xs12>
	 											Contributors: Syahrul Nizam
	 										</v-flex>
	 										<v-flex xs12>
	 											Platform: {{project.platform}}
	 										</v-flex>
	 										<v-flex xs12>
	 											Frameworks used:
	 												<v-layout justify-center wrap>
	 													<v-flex xs3 v-for=" tool in project.tools">
	 														<p>{{tool}}</p>
	 													</v-flex>
	 												</v-layout>
	 										</v-flex>
	 										<v-flex xs12>
	 											Languages:
	 											<v-layout justify-center wrap>
	 												<v-flex xs3 v-for="languages in project.languages">
	 													<p>{{languages}}</p>
	 												</v-flex>
	 											</v-layout>
	 										</v-flex>
	 									</v-layout>
 									</v-container>
 								</v-card>
 							</v-flex>

 						</v-layout>
 					</v-container>
 				</v-card>
 			</v-flex>
 		</v-layout>
 	</v-container>
 </v-app>
</template>

<script>
	import json from '../assets/data.json'
	export default {
		data: () => ({
			hover:false,
			id:null,
			project:null,
			myFileName:null,
			
		}),
		created() {
			this.id = this.$route.params.id
			json.Projects.find(item => {
				if(item.id == this.id){
					this.project = item;
					this.myFileName = item.myFileName

				}
			})

		}
	}
</script>
<style>
#dappHeader {
	font-family: Amsi Pro Cond SemiBold,Noto Sans SC,sans-serif;
	font-size: 20px;
	padding-left: 15px;
	background-color: #9cf196;

}

</style>