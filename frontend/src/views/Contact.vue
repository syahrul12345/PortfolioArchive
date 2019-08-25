<template>
	<v-app>
		<Header id="header" v-if="windowWidth > 600 "> "></Header>
    <HeaderSmall id="headerSmall" v-if="windowWidth < 600"></HeaderSmall>
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
	import HeaderSmall from '../components/HeaderSmall'
	export default {
		props:[''],
		components: {
			Header,
			HeaderSmall
		},
		data() {
			return {
				created:null,
				windowWidth:null,
			}
		},
		mounted() {
			this.$nextTick(function() {
				window.addEventListener('resize', this.getWindowWidth);
				window.addEventListener('resize', this.getWindowHeight);

				//Init
				this.getWindowWidth()
			})
		},
		methods: {
		    getWindowWidth() {
		      this.windowWidth = document.documentElement.clientWidth;
		    }
  		}
	}
</script>
<style></style>