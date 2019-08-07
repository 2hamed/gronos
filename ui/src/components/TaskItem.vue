<template>
  <div>
    <h2>{{task['name']}}</h2>
    <ul>
      <li v-for="command in task.command">{{command}}</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "TaskItem",
  props: ["task"],
  mounted() {
    console.log(this.task);
  }
};
</script>

<style>
</style>


