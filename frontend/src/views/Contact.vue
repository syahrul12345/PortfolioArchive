<template>
	<v-app>
		<Header></Header>
		<v-container grid-list-xs>
			<v-layout wrap>
				<v-flex xs>
					
				</v-flex>
			</v-layout>
		</v-container>
	</v-app>
	
</template>
<script>
	import Header from '../components/Header'
	export default {
		props:[''],
		components: {
			Header
		},
		data() {
			return {
				created:null,
			}
		}
	}
</script>
<style></style>