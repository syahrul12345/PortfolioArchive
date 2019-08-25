<template>
		<div class="header mb-4">
			<v-container grid-list-md :fluid="true">
				<v-layout wrap>
					<v-flex xs12>
						<v-card height="50"
						:outline="true"
						:flat="false"
						dark>
							<v-card-title class="py-0"> 
          						<v-flex xs1 offset-xs11>
								<v-menu bottom left>
						            <template v-slot:activator="{ on }">
						              <v-btn
						                dark
						                icon
						                v-on="on"
						                class="mt-1"
						              >
						             <v-icon>mdi-equal-box</v-icon>
						              </v-btn>
						            </template>
						            <v-list>
						            	<router-link 
											v-for="button in buttons" 
											:key="button.id" 
											:to= "button.link"
											style="text-decoration: none"
											text-center>
												<v-list-item xs1 text-center style="padding-top: 13px">
													<v-list-item-title> {{button.name}}</v-list-item-title>
												</v-list-item>

											</router-link>
						            </v-list>
						          </v-menu>
						      </v-flex>
							</v-card-title>
						
						</v-card>
					</v-flex>
					
				</v-layout>
			</v-container>
		</div>
	
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