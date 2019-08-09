<template>
	<v-app>
		<Header></Header>
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