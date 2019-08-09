<template>
		<div class="header">
			<v-container grid-list-md :fluid="true">
				<v-tabs dark v-model="activeTab">
					<router-link 
					v-for="button in buttons" 
					:key="button.id" 
					:to= "button.link"
					style="text-decoration: none"
					text-center>
						<v-tab xs1 text-center style="padding-top: 13px">
							<v-card
							:flat="true"
							:outline="true"
							:hover="true"
							text-center
							> {{button.name}}</v-card>
					</v-tab>	
					</router-link>
					
				</v-tabs>
			</v-container>
		</div>
	<!-- 	<v-card color="grey lighten-4" flat>
			<v-toolbar :color="primary" :dark="true">
				<v-toolbar-title>Title</v-toolbar-title>
			</v-toolbar>
		</v-card> -->
</template>
<script>
	export default {
		props:[''],
		data() {
			return {
				created:null,
				tab:null,
				activeTab:null,
				buttons: [
					{"id":1,name:"Home","link":"/"},
					{"id":2,name:"All DApps","link":"/dapps"},
					{"id":3,name:"All Articles","link":"/allblogs"},
					{"id":4,name:"Resume","link":"/resume"},
					{"id":5,name:"Contact Me","link":"/contact"},
				]
			}
		},
		created()  {
			this.buttons.filter((item) => {
				if(this.$route.path == item.link){
					this.activeTab = item.id - 1

				}
			})
		}
	}
</script>
<style>
	.header{
		font-size:20px;
		height:3%;
		background-color: dark;
		margin-bottom: 2%
	}
	li a {
		text-decoration: none;
	}
</style>