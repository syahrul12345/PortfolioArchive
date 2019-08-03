<template>
	<v-container grid-list-xs text-centre>
		<v-layout wrap>
			<v-flex xs10 offset-xs1>
				<v-card
				class="mx-auto"
				:outline=false
				:hover=true>
					<v-list-item three-line>
						<v-list-item-content>
							<v-list-item-title>Syahrul Nizam </v-list-item-title>
							<v-list-item-subtitle> I'm a solidity developer </v-list-item-subtitle>
						</v-list-item-content>
						<v-list-item-avatar
						tile
						size="80"
						color="grey">
						</v-list-item-avatar>
					</v-list-item>
				</v-card>
			</v-flex>
	</v-layout>
	</v-container>
</template>

<script>
	export default {
		data: () => {
			messages:null
		}
	}
</script>

<style>
	
</style>