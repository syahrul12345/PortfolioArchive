<template>
  <v-app style="background-color: white">
    <link href="https://use.fontawesome.com/releases/v5.0.13/css/all.css" rel="stylesheet">
       <ProfileCard 
       :name="Name" 
       :blurb ="Blurb" 
       ></ProfileCard>
     <Divider 
     :DividerText="DividerText1" :linkText="'See all DApps'" :link="`/dapps`":fluid="true"
     ></Divider>
    <v-container 
    grid-list-xl 
    text-center 
    :fluid ="true" 
    style="padding-left:4%;padding-right:4%;">
      <v-layout wrap>
        <v-flex v-for="project in Projects" :key="project.id" xs12 sm4 md4 lg3>
          <ProjectCard
          :title="project.Name" 
          :blurb="project.Blurb" 
          :previewFileName="project.MyFileName"
          :dappId="project.Name"
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
const axios = require('axios')
import ProfileCard from '../components/ProfileCard'
import ProjectCard from '../components/ProjectCard'
import Divider from '../components/Divider' 
import LatestBlog from '../components/LatestBlog'
import Footer from '../components/Footer'
export default {
  components: {
    ProfileCard,
    ProjectCard,
    Divider,
    LatestBlog,
    Footer,
  },
  data: () => ({
    Name:'Syahrul Nizam',
    Blurb:"I'm a Solidity Developer and I create Blockchain Applications",
    DividerText1:"Featured DApps",
    DividerText2:"Article Spotlight",
    Employers:[
      {name:"Chainstack",duration:"June",blurb:"Have lots of fun!"},
      {name:"Talenta",duration:"April",blurb:"Have lots of fun!"},
      {name:"Neo Performance Materials",duration:"April",blurb:"Have lots of fun!"},
      {name:"NCS",duration:"April",blurb:"Have lots of fun!"}

    ],
    Projects:null,
    dialog:false,
    windowWidth:null,
    //
  }),
  created() {
    axios.get('/api/user/getAllApps').then((response) => {
      this.Projects = response.data.applications.slice(0,4)
    }).catch((error) => {
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
};
</script>

<style>

</style>
