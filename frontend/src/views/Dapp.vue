<template>
<v-app>
	<Header id="header" v-if="windowWidth > 600 "> "></Header>
    <HeaderSmall id="headerSmall" v-if="windowWidth < 600"></HeaderSmall>
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
					<p> {{project.name}} </p>
				</div>
				<v-container grid-list-xs>
					<v-layout wrap>
						<v-flex xs12 lg7 offset-lg1>
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
								{{ project.blurb}}
							</v-card>
						</v-flex>
						</v-flex>
						<v-flex xs12 lg4>
							<v-card
							:hover="false"
							:outline="true"
							:flat="true">
								<!-- Buttons -->
								<v-container grid-list-xs>
									<v-layout justify-center wrap>
										<v-flex xs12 text-center>
											<v-card
											:hover="true"
											:outline="true"
											:flat="false"
											:text="true"
											text-center
											light
											:href="`//`+project.website">
												Visit the website!
											</v-card>
										</v-flex>
										<v-flex xs12 text-center>
											<v-card
											:hover="true"
											:outline="true"
											:flat="false"
											:text="true"
											text-center
											light
											:href="`//`+project.github">
												Github
											</v-card>
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
														v-bind:id="`${project.platform}`">
															{{project.platform}}
														</v-card>
													</v-flex>
												</v-layout>
										</v-flex>
										<v-flex xs12>
											Frameworks used:
												<v-layout  wrap>
													<v-flex xs3 v-for=" tool in project.tools"
													v-bind:key="project.id"
													text-center>
														<v-card
														class="mx-auto"
														:hover="false"
														:flat="true"
														:outline="false"
														v-bind:id="`${tool}`">
															{{tool}}
														</v-card>
														
													</v-flex>
												</v-layout>
										</v-flex>
										<v-flex xs12>
											Languages:
											<v-layout wrap>
												<v-flex xs3 v-for="languages in project.languages"
												:key="project.id"
												text-center>
													<v-card
														class="mx-auto"
														:hover="false"
														:flat="true"
														:outline="false"
														v-bind:id="`${languages}`">
															{{languages}}
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
	import json from '../assets/data.json'
	import Header from '../components/Header'
	import HeaderSmall from '../components/HeaderSmall'
	export default {
		components: {
			Header,
			HeaderSmall
		},
		data: () => ({
			hover:false,
			id:null,
			project:null,
			myFileName:null,
			windowWidth:null,
			
		}),
		created() {
			this.id = this.$route.params.id
			json.Projects.find(item => {
				if(item.id == this.id){
					this.project = item;
					this.myFileName = item.myFileName

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

@media(min-width: 960px){
	#mainCard {
		padding-left: 12%;
		padding-right: 12%;
		padding-top: 2%
	}
}
@media(max-width: 959px){
	#mainCard {
		padding-top: 10%
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

</style>