<template>
	<div @click.stop="dialog = false">
	<v-container grid-list-xl text-centre fluid="fluid" style="background-image: linear-gradient(#98FF98, #CCFF98)">
		<v-layout>
			<v-flex xs12 >
				<v-card
				class="mx-auto"
				:outline="outline"
				:hover="hover"
				:flat="flat"
				color="rgba(51,51,51,0)"
				>
					<v-container grid-list-xs text-centre fluid="fluid">
						<v-layout wrap>
							<v-flex xs12 sm12 md10>
								<v-card-text>
									<h1 id="myName"> {{ name }} </h1>
									<h2 id="myBlurb"> {{ blurb }}</h2>
								</v-card-text>
							</v-flex>
							
						</v-layout>
					</v-container>
					<v-container grid-list-xs text-centre fluid="fluid">
						<v-layout wrap>
							<v-card-actions>
								<router-link :to="this.allDappPage" style="text-decoration: none">
								<v-flex xs12>
									<v-btn outlined style="font-size: 15px">Explore All my Dapps</v-btn>
								</v-flex>
								</router-link>
							</v-card-actions>
						</v-layout>
					</v-container>
				</v-card>
				
			</v-flex>
		</v-layout>
	</v-container>
</div>
</template>

<script>
	export default {
		props:['avatar','name' ,'blurb'],
		data: () => ({
			profileLink: "../assets/profile.png",
			fluid:true,
			outline:false,
			hover:false,
			flat:true,
			dialog:false,
			allDappPage:"/dapps"
		}),
		
	}
</script>

<style>

#myName {
	font-family: Amsi Pro Cond SemiBold,Noto Sans SC,sans-serif;
	font-size: 30px;
	padding-bottom: 25px;

}
#myBlurb {
	font-family: Amsi Pro Cond SemiBold,Noto Sans SC,sans-serif;
	font-size: 25px
}
li a {
    text-decoration: none;
}
	
</style>