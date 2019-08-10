<template>
  <div>
    <h2>{{task['name']}}</h2>
    <ul>
      <li v-for="command in task['command']" :key="command">{{command}}</li>
    </ul>
    <button @click="toggleStatus">{{this.taskStatusContrary}} Task</button>
  </div>
</template>

<script>
export default {
  name: "TaskItem",
  props: ["task", "taskItemStatus", "index"],
  mounted() {
    // eslint-disable-next-line
    console.log(this.task);
  },
  methods: {
    toggleStatus() {
      this.$http
        .get(
          `/tasks/${this.task.name}/${this.taskStatusContrary.toLowerCase()}`
        )
        .then(() => {
          alert(`Task ${this.taskStatusContrary}d`);
          this.$emit('listUpdated', this.task);
        });
    }
  },
  computed: {
    taskStatusContrary() {
      return this.taskItemStatus ? "Enable" : "Disable";
    }
  }
};
</script>

<style>
</style>


