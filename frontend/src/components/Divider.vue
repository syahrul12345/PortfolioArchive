<template>
	<v-container grid-list-xs :fluid="fluid">
		<v-layout wrap>
			<v-flex xs12>
				<v-card
				class="mx-auto"
				:outline=false
				:flat=true
				color="rgba(51,51,51,0)">
					<v-card-text id="divider"> 
						{{ DividerText }}     <a :href="link">{{linkText}}</a>
					</v-card-text>

				</v-card>
			</v-flex>
			
		</v-layout>
	</v-container>
</template>
<script>
	export default {
		props:['DividerText','fluid','linkText','link'],
		data:()=> ({

		})
	}	
</script>
<style>
	#divider {
		font-size: 20px;
		font-family: Amsi Pro Cond SemiBold,Noto Sans SC,sans-serif;
	}
	a{
		font-size:15px;
	}
</style>