<template>
	<v-card
	class="mx-auto"
	>
		<v-list-item three-line>
			<v-list-item-content>
				<v-list-item-title> {{Employer}}</v-list-item-title>
				<v-list-item-subtitle> {{Duration}}</v-list-item-subtitle>
				<v-list-item-subtitle> {{Blurb}}</v-list-item-subtitle>
			</v-list-item-content>
			<v-list-item-avatar
						tile
						size="80"
						color="grey">
			</v-list-item-avatar>
		</v-list-item>
	</v-card>
</template>
<script>
	export default {
		props:['Employer','Duration','Blurb'],
		data: () => {
		}
	}
</script>
<style></style>