<template>
	<v-card
	:outline="true"
	:hover="false">
	<v-card-title> Skills </v-card-title>
		<div> 
			<v-card-text> Platforms: 
				<v-container grid-list-xs>
					<v-layout wrap>
						<v-flex xs4 text-center v-for="platform in platforms">
							<v-card
							:text="true"
							:hover="false"
							text-center
							v-bind:id="`${platform}`"
							>
								{{platform}}
							</v-card>
						</v-flex>
					</v-layout>
				</v-container>
			</v-card-text>
		</div>
		<div> 
			<v-card-text> Frameworks: 
				<v-container grid-list-xs>
					<v-layout wrap>
						<v-flex xs4 text-center v-for="framework in frameworks">
							<v-card
							:text="true"
							:hover="false"
							text-center
							v-bind:id="`${framework}`"
							>
								{{framework}}
							</v-card>
						</v-flex>
					</v-layout>
				</v-container>
			</v-card-text>
		</div>
		<div> 
			<v-card-text> Languages: 
				<v-container grid-list-xs>
					<v-layout wrap>
						<v-flex xs4 text-center v-for="language in languages">
							<v-card
							:text="true"
							:hover="false"
							text-center
							v-bind:id="`${language}`"
							>
								{{language}}
							</v-card>
						</v-flex>
					</v-layout>
				</v-container>
			</v-card-text>
		</div>
		
	<div>
		
	</div>

	</v-card>
</template>
<script>
	export default {
		props:[''],
		components: {

		},
		data() {
			return {
				created:null,
				frameworks:["VueJS","truffle","embark","web3JS"],
				languages:["Python","Solidity","Javascript"],
				platforms:["Ethereum","Quorum","Hyperledger","Multichain"]
			}
		}
	}
</script>
<style></style>