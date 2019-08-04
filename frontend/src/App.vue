<template>
  <v-app >
    <link href="https://use.fontawesome.com/releases/v5.0.13/css/all.css" rel="stylesheet">
       <ProfileCard :name="Name" :blurb ="Blurb" :employers="Employers"></ProfileCard>
     <Divider 
     :DividerText="DividerText1" :fluid="true"></Divider>
    <v-container grid-list-xl :fluid ="true">
      <v-layout wrap>
        <v-flex v-for="project in Projects" xs3>
          <ProjectCard
          :title="project.name" 
          :blurb="project.blurb" 
          :url="project.url"
          :mediaSource="project.mediaSource"></ProjectCard>
        </v-flex>
      </v-layout>
    </v-container>
    <Divider 
     :DividerText="DividerText2" :fluid="true"></Divider>
    <LatestBlog></LatestBlog>
    
  </v-app>
  
</template>

<script>
import ProfileCard from './components/ProfileCard'
import ProjectCard from './components/ProjectCard'
import Divider from './components/Divider' 
import LatestBlog from './components/LatestBlog'
import Footer from './components/Footer'

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
    DividerText1:"Featured DApps",
    DividerText2:"Published Articles",
    Employers:[
      {name:"Chainstack",duration:"June",blurb:"Have lots of fun!"},
      {name:"Talenta",duration:"April",blurb:"Have lots of fun!"},
      {name:"Neo Performance Materials",duration:"April",blurb:"Have lots of fun!"},
      {name:"NCS",duration:"April",blurb:"Have lots of fun!"}

    ],
    Projects:[
      {name:'Project 1',blurb:'This is project 1',url:'//www.google.com',mediaSource:'https://cdn.vuetifyjs.com/images/cards/docks.jpg'},
      {name:'Project 2',blurb:'This is project 1',url:'//www.google.com',mediaSource:'https://cdn.vuetifyjs.com/images/cards/docks.jpg'},
      {name:'Project 3',blurb:'This is project 1',url:'//www.google.com',mediaSource:'https://cdn.vuetifyjs.com/images/cards/docks.jpg'},
      {name:'Project 4',blurb:'This is project 1',url:'//www.google.com',mediaSource:'https://cdn.vuetifyjs.com/images/cards/docks.jpg'}
    ],
    dialog:false
    //
  }),
};
</script>

<style>
  
</style>
