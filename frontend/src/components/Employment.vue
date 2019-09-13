<template>
	<v-card
	class="mx-auto"
	dark
	style="margin-bottom: 10px"
	>
		<v-list-item three-line>
			<v-list-item-content>
				<v-list-item-title> {{Employer}}</v-list-item-title>
				<v-list-item-title> {{Title}}</v-list-item-title>
				<v-list-item-subtitle> {{Duration}}</v-list-item-subtitle>
				<v-list-item-subtitle> {{Blurb}}</v-list-item-subtitle>
			</v-list-item-content>
			<v-list-item-avatar
						tile
						size="80"
						color="white">
						<v-img :src="require('../assets/images/employer/' + Image)"></v-img>
			</v-list-item-avatar>
		</v-list-item>
	</v-card>
</template>
<script>
	export default {
		props:['Employer','Title','Duration','Blurb','Image'],
		data() {
			return {
				created:null,
			}
		},
		created() {
		}
	}
</script>
<style></style>