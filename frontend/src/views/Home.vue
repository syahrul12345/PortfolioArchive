<template>
  <v-app>
    <link href="https://use.fontawesome.com/releases/v5.0.13/css/all.css" rel="stylesheet">
       <ProfileCard :name="Name" :blurb ="Blurb" :employers="Employers"></ProfileCard>
     <Divider 
     :DividerText="DividerText1" :fluid="true"></Divider>
    <v-container grid-list-xl text-center :fluid ="true">
      <v-layout wrap>
        <v-flex v-for="project in Projects" xs3>
          <ProjectCard
          :title="project.name" 
          :blurb="project.blurb" 
          :previewFileName="project.myFileName"
          :dappId="project.id"
          ></ProjectCard>
        </v-flex>
      </v-layout>
    </v-container>
    <!-- <Divider 
     :DividerText="DividerText2" :fluid="true"></Divider>
    <LatestBlog :latestBlogTitle="LatestBlogTitle" :latestBlogBlurb="LatestBlogBlurb"></LatestBlog>
     -->
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
    Blurb:"Developer Advocate in the nascent bockchain space. Specializes in developing DApps on top Quroum and Ethereum. Currently advocating for Chainstack.",
    LatestBlogTitle: "Some Blog",
    LatestBlogBlurb: "Some Blurb2",
    DividerText1:"Featured DApps",
    DividerText2:"Published Articles",
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
