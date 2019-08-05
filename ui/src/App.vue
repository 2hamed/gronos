<template>
  <div id="app">
    
  </div>
</template>

<script>


export default {
  name: 'app',
  data(){
    return {
      tasks: []
    }
  },
  methods: {
    getTasks(){
      axios.get('http//localhost:8080/tasks').then((res)=>{
        
      });
    }
  },
}
</script>

<style>

</style>
