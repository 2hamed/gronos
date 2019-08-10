<template>
  <div>
    <button
      @click="toggleFilterTasks"
    >{{displayDisableTasks ? 'Filter By Disabled Tasks' : 'Filter By Enabled Tasks'}}</button>
    <h2>{{displayDisableTasks ? 'Disable' : 'Enable'}} List</h2>
    <div v-for="(task, index) in tasks" :key="index">
      <TaskItem
        :task="task"
        :taskItemStatus="displayDisableTasks"
        @listUpdated="listUpdated"
        :taskIndex="index"
      ></TaskItem>
    </div>
    <div v-show="tasks.length == 0">There are no items to display</div>
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
      tasks: [],
      displayDisableTasks: false
    };
  },
  mounted() {
    this.getTasks();
  },
  computed: {
    requestUrl() {
      return this.displayDisableTasks ? "/tasks/disabled" : "/tasks";
    }
  },
  methods: {
    toggleFilterTasks() {
      this.displayDisableTasks = !this.displayDisableTasks;
      this.getTasks();
      return this.displayDisableTasks;
    },
    listUpdated() {
      this.getTasks();
    },
    getTasks() {
      this.loading = true;
      this.$http
        .get(this.requestUrl)
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