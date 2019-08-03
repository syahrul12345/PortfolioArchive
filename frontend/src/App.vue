<template>
  <div @click.stop="dialog = false">
  <v-app>
    <link href="https://use.fontawesome.com/releases/v5.0.13/css/all.css" rel="stylesheet">
     <div @click.stop="dialog = !dialog">
       <ProfileCard :avatar="DisplayPicture"></ProfileCard>
     </div>
     <v-dialog v-model="dialog" persistent max-width="600px">
          <v-container v-for="employer in Employers">
            <Employment :Employer="employer.name" :Duration ="employer.duration" :Blurb="employer.blurb"></Employment>
          </v-container>
      </v-dialog>
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
  </div>
</template>

<script>
import ProfileCard from './components/ProfileCard'
import ProjectCard from './components/ProjectCard'
import Employment from './components/Employment'
import Divider from './components/Divider' 

export default {
  name: 'App',
  components: {
    ProfileCard,
    ProjectCard,
    Divider,
    Employment
  },
  data: () => ({
    DisplayPicture:"https://cdn.vuetifyjs.com/images/cards/docks.jpg",
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
