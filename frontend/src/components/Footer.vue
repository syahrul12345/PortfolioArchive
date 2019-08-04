<template>
	<v-container grid-list-xs>
		<v-layout wrap>
			<v-flex xs-1 offset-xs5>
				<v-footer padless>
					<v-card
					flat
					tile
					class="indigo lighten-1 white--text text-center">
						<v-card-text>
							<v-btn
							v-for="icon in icons"
							v-key="icon"
							class="mx-4 white--text"
							icon>
								<v-icon size="24px">{{ icon }}</v-icon>
							</v-btn>
						</v-card-text>
						<v-card-text>
							Made by Syahrul Nizam using VueJS
						</v-card-text>
					</v-card>
				</v-footer>
			</v-flex>
		</v-layout>
	</v-container>
	
</template>
<script>
	export default {
		data: () => ({
			icons: [
				'fab fa-github',
		        'fab fa-linkedin',
		        'fab fa-facebook',
			]
		})
	}
</script>
<style>
	
</style>