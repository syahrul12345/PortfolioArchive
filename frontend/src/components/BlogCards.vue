<template>
		<v-card
		class="mx-auto"
		:hover="hover"
		:href="`//${url}`"
		>
			<v-list-item id="listSet" three-line>
				<v-list-item-avatar color="white">
					<v-img :src="require('../assets/images/blogs/avatars/' + publication + '.png' )"></v-img>
				</v-list-item-avatar>
				<v-list-item-content>
					<v-list-item-title>{{ title }}</v-list-item-title>
					<v-list-item-subtitle >{{ blurb }}</v-list-item-subtitle>
				</v-list-item-content>
			</v-list-item>
			<div v-bind:id="`${publication}`">
				<p id="category">{{ publication }}</p>
			</div>
		</v-card>
	
</template>

<script>
	export default {
		props:['title','blurb','url','publication'],
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
	#chainstack{
		background-color: #b0deff;
	}
	#coinmonks{
		background-color: #d2cc1e;
	}
	#hackernoon{
		background-color: #3bb57a;
	}
	
	li a {
    text-decoration: none;
	}
</style>