<template>
	<router-link :to="`/dapp/` + dappId">
		<v-card
		class="mx-auto"
		:hover="hover"
		>
			<v-list-item id="listSet">
				<v-list-item-avatar color="grey">
					<v-img :src="mediaSource"></v-img>
				</v-list-item-avatar>
				<v-list-item-content>
					<v-list-item-title>{{ title }}</v-list-item-title>
					<v-list-item-subtitle >{{ blurb }}</v-list-item-subtitle>
				</v-list-item-content>
			</v-list-item>
			<div id="categoryBar">
				<p id="category">{{ category }}</p>
			</div>
		</v-card>
	</router-link>
	
</template>

<script>
	export default {
		//this props are binded to and values is passed from the parent App.vue
		props:['title','blurb','url','mediaSource','dappId','category'],
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
	#listSet {
		padding-top: 30px;
		padding-bottom: 30px;
	}
	#categoryBar {
		background-color: #9cf196;

	}
	#category {
		padding-left: 10px;
	}
</style>