<template>
	<v-card
	class="mx-auto"
	:hover="hover">
		<v-img
        v-if="media"
        class="white--text"
        v-bind:src="mediaSource"
      	>
		</v-img>
		<v-list-item three-line>
			<v-list-item-content>
				<v-list-item-title>{{ title }}</v-list-item-title>
				<v-list-item-subtitle >{{ blurb }}</v-list-item-subtitle>
			</v-list-item-content>
		</v-list-item>
		<v-container grid-list-xl>
			<v-layout wrap>
				<v-card-actions v-if="actions">
					<v-flex xs10>
						<v-btn outlined style="font-size: 10px">Read More</v-btn>
					</v-flex>
					<v-flex xs10 offset-xs7>
						<v-btn
						:text=true
						:href="url">
							<v-icon>fab fa-github</v-icon>
						</v-btn>
					</v-flex>
					
				</v-card-actions>
			</v-layout>
		</v-container>
		
		
	</v-card>
	
</template>

<script>
	export default {
		//this props are binded to and values is passed from the parent App.vue
		props:['title','blurb','url','mediaSource'],
		data: () => ({
			media:true,
			hover:true,
			actions:true,
			projectTitle:null,
			projectBlurd:null
		})
	}

</script>
<style></style>