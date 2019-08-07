<template>
  <div>
    <div v-for="task in tasks">
      <TaskItem :task="task"></TaskItem>
    </div>
  </div>
</template>

<script>
import TaskItem from "./TaskItem";
export default {
  name: "Tasks",
  components: { TaskItem },
  data() {
    return {
      loading: false,
      tasks: []
    };
  },
  mounted() {
    this.getTasks();
  },
  methods: {
    getTasks() {
      this.loading = true;
      this.$http
        .get("/tasks")
        .then(res => {
          this.tasks = res.data.data;
        })
        // eslint-disable-next-line
        .catch(err => console.log(err))
        .finally(() => (this.loading = false));
    }
  }
};
</script>

<style>
</style>