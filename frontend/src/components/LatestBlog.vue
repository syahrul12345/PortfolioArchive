<template>
	<v-container grid-list-md>
      <v-layout shrink>
        <v-flex xs10 offset-xs1>
        	<v-card
        	class="mx-auto"
        	:flat="flat"
        	:outline="outline"
        	color="#ffffff">
        		<v-container grid-list-md :fluid="fluid">
        			<v-layout wrap>
        				<v-flex xs4>
        					<v-card
        					class="mx-auto"
        					:flat="flat"
        					:outline="outline"
        					elevation=true>
	        					<v-avatar
	        					class="profile"
	        					:size="$vuetify.breakpoint.smAndUp ? 328 : 256"
								color="gray"
								tile>
	        						<v-img src="../assets/profile.png"></v-img>
	        					</v-avatar>
	        				</v-card>
        				</v-flex>
        				<v-flex xs8>
        					<v-card
        					class="mx-auto"
        					>
        						<h1> Some Project</h1>
        						<h2> Project Blurb </h2>
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
	
	export default {
		data: () => ({
			fluid:false,
			flat:false,
			outline:true,
		})
	}

</script>
<style></style>