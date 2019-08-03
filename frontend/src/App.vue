<template>
  <v-app>
    <link href="https://use.fontawesome.com/releases/v5.0.13/css/all.css" rel="stylesheet">
     <ProfileCard></ProfileCard>
     <Divider 
     :DividerText="`Decentralized Application (DApp) Projects`"></Divider>
    <v-container grid-list-md>
      <v-layout wrap>
        <v-flex v-for="project in Projects" xs4>
          <ProjectCard
          :title="project.name" 
          :blurb="project.blurb" 
          :url="project.url"
          :mediaSource="project.mediaSource"></ProjectCard>
        </v-flex>
      </v-layout>
    </v-container>

  </v-app>
</template>

<script>
import ProfileCard from './components/ProfileCard'
import ProjectCard from './components/ProjectCard'
import Divider from './components/Divider'

export default {
  name: 'App',
  components: {
    ProfileCard,
    ProjectCard,
    Divider
  },
  data: () => ({
    Projects:[
      {name:'Project 1',blurb:'This is project 1',url:'//www.google.com',mediaSource:'https://cdn.vuetifyjs.com/images/cards/docks.jpg'},
      {name:'Project 2',blurb:'This is project 1',url:'//www.google.com',mediaSource:'https://cdn.vuetifyjs.com/images/cards/docks.jpg'},
      {name:'Project 3',blurb:'This is project 1',url:'//www.google.com',mediaSource:'https://cdn.vuetifyjs.com/images/cards/docks.jpg'},
      {name:'Project 4',blurb:'This is project 1',url:'//www.google.com',mediaSource:'https://cdn.vuetifyjs.com/images/cards/docks.jpg'}
    ]
    //
  }),
};
</script>

<style>
  
</style>
