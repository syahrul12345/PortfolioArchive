<template>
	<v-container grid-list-xl text-centre>
		<v-layout wrap>
			<v-flex xs12>
				<v-card
				class="mx-auto"
				:outline=false
				:hover=true>
					<v-list-item three-line>
						<v-list-item-content>
							<v-list-item-title id="myName">Syahrul Nizam </v-list-item-title>
							<v-list-item-subtitle id="myCV"> Developer advocate in the nascent space of blockchain. I create blockchain Decentralized Applications (DApps). 
							</v-list-item-subtitle>
						</v-list-item-content>
					</v-list-item>
					<v-container grid-list-md>
						<v-layout wrap>
							<v-card-actions>
								<v-flex xl10>
									<v-btn outlined style="font-size: 20px">Read More</v-btn>
								</v-flex>
							</v-card-actions>
						</v-layout>
					</v-container>
				</v-card>
				
			</v-flex>
	</v-layout>
	</v-container>
</template>

<script>
	export default {
		props:['avatar'],
		data: () => ({
			
		})
	}
</script>

<style>
#myName {
	font-size: 30px
}
#myCV {
	font-size: 25px
}
	
</style>