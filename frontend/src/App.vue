<template>
 <v-app>
 	<v-app-bar id = "app-bar-desktop'" v-if="windowWidth > 600">
 		<v-tabs v-model="activeTab">
 			<v-tab v-for="button in buttons"
 			:key="button.id"
 			:to="button.link"
 			style="text-decoration: none"
 			text-center>
 				{{button.name}}
 			</v-tab>
			
		</v-tabs>
	</v-app-bar>
	<v-app-bar id = "app-bar-mobile'" v-if="windowWidth < 600">
		<v-app-bar-nav-icon>
			<v-menu>
	            <template v-slot:activator="{ on }">
	              <v-btn
	                icon
	                v-on="on"

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
							<v-list-item xs1 text-center>
								<v-list-item-title> {{button.name}}</v-list-item-title>
							</v-list-item>

						</router-link>
	            </v-list>
          </v-menu>
		</v-app-bar-nav-icon>
        <v-toolbar-title class="px-0">Portfolio</v-toolbar-title>
        <div class="flex-grow-1"></div>
        <v-btn icon></v-btn>	
	</v-app-bar>
  	<v-content id="animated">
  		<router-view/>
  	</v-content>
 </v-app>
  
</template>
<script>
	export default {
		components: {
				
		},
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
				],
				windowWidth:null
			}
		},
		created()  {
			this.buttons.filter((item) => {
				if(this.$route.path == item.link){
					this.activeTab = item.id - 1

				}
			})
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
<style>
</style>
