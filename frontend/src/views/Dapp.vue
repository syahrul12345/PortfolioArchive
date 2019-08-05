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
	 					<p class= "dappName"> HELLO </p>
 					</div>
 					<v-container grid-list-xl>
 						<v-layout wrap>
 							<v-flex xs7 offset-xs1>
 								<v-card
 								:hover="false"
 								:outline="true"
 								:flat="true">
 										<v-img src="../assets/Multis.jpg"></v-img>
 								</v-card>
 								<v-flex xs12>
 								<v-card
 								:hover="false"
 								:outline="true"
 								:flat="true">
 									"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.""Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
 								</v-card>
 							</v-flex>
 							</v-flex>
 							<v-flex xs3>
 								<v-card
 								:hover="false"
 								:outline="true"
 								:flat="true">
 									"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.""Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
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
	export default {
		data: () => ({
			hover:false,
		}),
		mounted() {
			console.log('The id is: ' + this.$route.params.id);
		},
	}
</script>
<style>
#dappHeader {
	background-color: green;
}
.dappName {
	font-family: Amsi Pro Cond SemiBold,Noto Sans SC,sans-serif;
	font-size: 20px;
	padding-left: 2px;
}

</style>