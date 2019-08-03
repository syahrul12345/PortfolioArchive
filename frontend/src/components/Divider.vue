<template>
	<v-container grid-list-xs>
		<v-layout wrap>
			<v-flex xs12>
				<v-card
				class="mx-auto"
				:outline=false
				:flat=true>
					<v-card-text id="divider"> {{ DividerText }} </v-card-text>				
				</v-card>
			</v-flex>
			
		</v-layout>
	</v-container>
</template>
<script>
	export default {
		props:['DividerText'],
		data:()=> ({

		})
	}	
</script>
<style>
	#divider {
		font-size: 20px;
		font-family: montserrat;
	}
</style>