<template>
	<div @click.stop="dialog = false">
	<v-container grid-list-xl text-centre v-bind:fluid="fluid" >
		<v-layout wrap>
			<v-flex xs12>
				<v-card
				class="mx-auto"
				:outline="outline"
				:hover="hover"
				:flat="flat">
					<v-container grid-list-xs text-centre v-bind:fluid="fluid">
						<v-layout shrink>
							<v-flex xs-10>
								<v-card-text>
									<h1 id="myName"> {{ name }} </h1>
									<h2 id="myBlurb"> {{ blurb }}</h2>
								</v-card-text>
							</v-flex>
							<v-flex xs-2>
								<v-avatar
								class = "profile"
								:size="$vuetify.breakpoint.smAndUp ? 164 : 128"
								color="gray"
								tile>
									<v-img src="../assets/profile.png"></v-img>
								</v-avatar>
							</v-flex>
						</v-layout>
					</v-container>
					<v-container grid-list-xs text-centre v-bind:fluid="fluid">
						<v-layout wrap>
							<v-card-actions>
								<v-flex xl12>
									<v-btn outlined style="font-size: 15px" @click.stop="dialog = !dialog">Employment History</v-btn>
								</v-flex>
								<v-flex xs12>
									<v-btn outlined style="font-size: 15px">Explore All my Dapps</v-btn>
								</v-flex>
							</v-card-actions>
						</v-layout>
					</v-container>
				</v-card>
				
			</v-flex>
		</v-layout>
		<!-- <v-dialog v-model="dialog" persistent max-width="600px">
          <v-container v-for="employer in employers">
            <Employment :Employer="employer.name" :Duration ="employer.duration" :Blurb="employer.blurb"></Employment>
          </v-container>
      	</v-dialog> -->
	</v-container>
</div>
</template>

<script>
	import Employment from './Employment'
	export default {
		props:['avatar','name' ,'blurb','employers'],
		
		data: () => ({
			profileLink: "../assets/profile.png",
			fluid:true,
			outline:false,
			hover:false,
			flat:true,
			dialog:false,
		}),
		
	}
</script>

<style>
#myName {
	font-size: 30px;
	padding-bottom: 25px;

}
#myCV {
	font-size: 25px
}
	
</style>