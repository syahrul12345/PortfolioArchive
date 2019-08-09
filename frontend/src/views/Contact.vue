<template>
	<v-app>
		<Header></Header>
		<v-container bg fill-height grid-list-md text-xs-center text-center>
			<v-layout row wrap align-center>
				<v-flex xs4>
					<a href="//www.linkedin.com/in/syahrul-nizam-41b984135/">
					<v-icon
					x-large>fab fa-linkedin-in
					</v-icon></a>
				</v-flex>
				<v-flex xs4>
					<a href="mailto:syahrul.nizam@u.nus.edu">
					<v-icon
					x-large>fas fa-envelope</v-icon></a>
				</v-flex>
				<v-flex xs4>
					<a href="//github.com/syahrul12345">
					<v-icon
					x-large>fab fa-github</v-icon></a>
					
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