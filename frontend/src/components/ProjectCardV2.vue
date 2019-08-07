<template>
		<v-card
		class="mx-auto"
		:hover="hover"
		>
			<v-list-item id="listSet" three-line>
				<v-list-item-avatar color="grey">
					<v-img :src="require('../assets/images/avatar/' + avatarSource)"></v-img>
				</v-list-item-avatar>
				<v-list-item-content>
					<v-list-item-title>{{ title }}</v-list-item-title>
					<v-list-item-subtitle >{{ blurb }}</v-list-item-subtitle>
				</v-list-item-content>
			</v-list-item>
			<div v-bind:id="`${category}`">
				<p id="category">{{ category }}</p>
			</div>
		</v-card>
	
</template>

<script>
	export default {
		props:['title','blurb','url','avatarSource','category'],
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
	#category {
		padding-left: 10px;
	}
	#Games{
		background-color: #b0deff;
	}
	#Data{
		background-color: #ffc5a1;
	}
	#Frontend{
		background-color: #ffd19a;
	}
	#Identity{
		background-color: #fff8a6;
	}
	#Finance{
		background-color: #d2f3e0;
	}
	#Wallet{
		background-color: #ac8daf;
	}
	#Markets{
		background-color: #ddb6c6;
	}
	#Insurance{
		background-color: #f1d4d4;
	}
	#Tools{
		background-color: #72d6c9;
	}
	li a {
    text-decoration: none;
	}
</style>