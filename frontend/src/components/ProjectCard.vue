<template>
	<router-link :to="`/dapp/` + dappId" style="text-decoration: none" >
		<v-card
		class="mx-auto"
		:hover="hover"
		:max-width="360">
			<v-img
			v-if="media"
			class="white--text"
			max-height="189"
			:aspect-ratio="1.77"
			:src="require('../assets/images/preview/' + previewFileName)"
			>
			</v-img>
			<v-list-item three-line>
				<v-list-item-content>
					<v-list-item-title>{{ title }}</v-list-item-title>
					<v-list-item-subtitle >{{ blurb }}</v-list-item-subtitle>
				</v-list-item-content>
			</v-list-item>
		</v-card>
	</router-link>

</template>

<script>
	export default {
		//this props are binded to and values is passed from the parent App.vue
		props:['title','blurb','previewFileName','dappId'],
		data: () => ({
			media:true,
			hover:true,
			actions:true,
			projectTitle:null,
			projectBlurd:null,
			
		})
	}

</script>
<style>
</style>