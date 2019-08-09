<template>
  <v-app>
    <link href="https://use.fontawesome.com/releases/v5.0.13/css/all.css" rel="stylesheet">
       <ProfileCard :name="Name" :blurb ="Blurb" :employers="Employers"></ProfileCard>
     <Divider 
     :DividerText="DividerText1" :linkText="'See all DApps'" :link="`/dapps`":fluid="true"></Divider>
    <v-container grid-list-xl text-center :fluid ="true">
      <v-layout wrap>
        <v-flex v-for="project in Projects" :key="project.id" xs3>
          <ProjectCard
          :title="project.name" 
          :blurb="project.blurb" 
          :previewFileName="project.myFileName"
          :dappId="project.id"
          ></ProjectCard>
        </v-flex>
      </v-layout>
    </v-container>
    <Divider 
     :DividerText="DividerText2" :linkText="'See all articles'" :link="`/allblogs`":fluid="true"></Divider>
    <LatestBlog></LatestBlog>
    
  </v-app>
  
</template>

<script>
import ProfileCard from '../components/ProfileCard'
import ProjectCard from '../components/ProjectCard'
import Divider from '../components/Divider' 
import LatestBlog from '../components/LatestBlog'
import Footer from '../components/Footer'
import json from '../assets/data.json'
export default {
  name: 'App',
  components: {
    ProfileCard,
    ProjectCard,
    Divider,
    LatestBlog,
    Footer,
    
  },
  data: () => ({
    Name:'Syahrul Nizam',
    Blurb:"Developer Advocate in the nascent bockchain space. Specializes in developing DApps on top Quroum and Ethereum,wtih a large interest in using Solidity to build Smart Contracts. Currently a Developer Advocate for Chainstack, where I create applications to demonstrate Chainstack's extremely simple Blockchain-on-the-fly deployment process.",
    DividerText1:"Featured DApps",
    DividerText2:"Article Spotlight",
    Employers:[
      {name:"Chainstack",duration:"June",blurb:"Have lots of fun!"},
      {name:"Talenta",duration:"April",blurb:"Have lots of fun!"},
      {name:"Neo Performance Materials",duration:"April",blurb:"Have lots of fun!"},
      {name:"NCS",duration:"April",blurb:"Have lots of fun!"}

    ],
    "Projects":json.Projects.slice(0,4),
    dialog:false
    //
  }),
};
</script>

<style>
  
</style>
