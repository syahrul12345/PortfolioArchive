<template>
	<v-container grid-list-md>
      <v-layout shrink>
        <v-flex xs10 offset-xs1>
        	<v-card
        	class="mx-auto"
        	:flat="true"
        	:outline="false"
        	color="rgba(51,51,51,0)">
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
        						<h1> {{latestBlogTitle}}</h1>
        						<h2> {{latestBlogBlurb}} </h2>
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
        props:['latestBlogTitle','latestBlogBlurb'],
		data: () => ({
			fluid:false,
			flat:false,
			outline:false,
		})
	}

</script>
<style></style>