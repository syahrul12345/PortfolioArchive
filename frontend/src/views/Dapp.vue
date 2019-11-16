<template>
<v-app>
<v-container id="mainCard"
grid-list-md 
:fluid="true" 
style="">
		<v-layout wrap>
			<v-flex xs12>
			<v-card
			class="mx-auto"
			:hover="hover">
				<div id="dappHeader">
					<p> {{project.Name}} </p>
				</div>
				<v-container grid-list-xs>
					<v-layout wrap>
						<v-flex xs12 lg7>
							<v-card
							:hover="false"
							:outline="true"
							:flat="true">
									<v-img :src="require('../assets/images/preview/' + myFileName)"></v-img>
							</v-card>
							<v-flex xs12>
							<v-card
							:hover="false"
							:outline="true"
							:flat="true">
								{{ project.Blurb}}
							</v-card>
						</v-flex>
						</v-flex>
						<v-flex xs12 lg5>
							<v-card
							:hover="false"
							:outline="true"
							:flat="true">
								<!-- Buttons -->
								<v-container grid-list-xs>
									<v-layout justify-center wrap>
										<v-flex xs12>
											<a style="text-decoration:none" :href="'//' +project.Website" class="mr-5">
											<v-icon large>fas fa-globe-americas</v-icon>
											</a>
											<a style="text-decoration:none" :href="'//' +project.Github">
											<v-icon large>fab fa-github </v-icon>
											</a>
										</v-flex>
										<v-flex xs12>
											Contributors: Syahrul Nizam
										</v-flex>
										<v-flex xs12>
											Platform: 
											<v-layout wrap>
													<v-flex xs3 text-center>
														<v-card
														class="mx-auto"
														:hover="false"
														:flat="true"
														:outline="false"
														v-bind:id="`${project.Platform}`">
															{{project.Platform}}
														</v-card>
													</v-flex>
												</v-layout>
										</v-flex>
										<v-flex xs12>
											Frameworks used:
												<v-layout  wrap>
													<v-flex xs3 v-for=" tool in project.Tools"
													v-bind:key="tool.Name"
													text-center>
														<v-card
														class="mx-auto"
														:hover="false"
														:flat="true"
														:outline="false"
														v-bind:id="`${tool.Name}`">
															{{tool.Name}}
														</v-card>
														
													</v-flex>
												</v-layout>
										</v-flex>
										<v-flex xs12>
											Languages:
											<v-layout wrap>
												<v-flex xs3 v-for="languages in project.Languages"
												:key="languages.Name"
												text-center>
													<v-card
														class="mx-auto"
														:hover="false"
														:flat="true"
														:outline="false"
														v-bind:id="`${languages.Name}`">
															{{languages.Name}}
													</v-card>
												</v-flex>
											</v-layout>
										</v-flex>
									</v-layout>
								</v-container>
							</v-card>
						</v-flex>

					</v-layout>
				</v-container>
			</v-card>
			</v-flex>
		</v-layout>
	</v-container>
</v-app>
</template>

<script>
	const axios = require('axios')
	export default {
		components: {
		},
		data: () => ({
			hover:false,
			id:null,
			project:null,
			myFileName:null,
			windowWidth:null,
			
		}),
		async created() {
			await axios.get('/api/user/getAllApps').then((response) => {
				const Projects = response.data.applications.slice(0,4)
				this.id = this.$route.params.id
				Projects.find((item) => {
					if(item.Name == this.id) {
						this.project = item;
						this.myFileName = item.MyFileName
					}
				})
			}).catch((error) => {
				console.log(error)
				console.log("FAILED TO GET APPLICATION LIST")
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
@media(max-width: 959px){
	#mainCard {
		padding-top: 10%
	}
}
@media(min-width: 960px){
	#mainCard {
		padding-left: 12%;
		padding-right: 12%;
		padding-top: 2%
	}
}

@media(min-width: 1200px){
	#mainCard {
		
	}
}
#dappHeader {
	font-family: Amsi Pro Cond SemiBold,Noto Sans SC,sans-serif;
	font-size: 20px;
	padding-left: 15px;
	background-color: #9cf196;

}

#Hyperledger{
	background-color: #b19cd9;
}

#Multichain{
	background-color: #fffe71;
}

#VueJS{
	background-color: #41b883;
}

#embark{
	background-color: #4571da;
}

#web3JS {
	background-color: #e1f7d5;
}
#truffle{
	background-color: #ffbdbd;
}
#ipfs{
	background-color: #c9c9ff;
}
#vueJS{
	background-color: #f1cbff;
}
#solidity{
	background-color: #cfb7ae;
}
#javascript{
	background-color: #aec6cf;
}
#Ethereum{
	background-color: #ff6961;
}
#Quorum{
	background-color: #0e8de1;
}

#Python{
	background-color: #FFD43B;
}
#Solidity{
	background-color: cyan;
}
#Javascript{
	background-color: #6cc24a;
}
#Golang{
        background-color: #29BEB0;
}
#ReactJS{
        background-color:#222222;
        color:white;
}
#Redux{
        background-color:#764abc;
}
#NextJS{
        background-color:white;     
}

</style>
