<script setup>
import { onMounted, ref } from "vue";
import logoImg from "./assets/images/theoneFROG.svg";

const tasks = ref([]);
const taskTitle = ref("");
const taskDescription = ref("");
const taskTag = ref("");

// Function to add a task
onMounted(async () => {
  try {
    tasks.value = await window.backend.App.LoadTasks();
  } catch (error) {
    console.error("Failed to load tasks:", error);
  }
});

const addTask = async () => {
  if (!taskTitle.value || !taskDescription.value || !taskTag.value) {
    alert("Please fill in all fields.");
    return;
  }
  const newTask = {
    id: tasks.value.length + 1, // Simple ID assignment
    title: taskTitle.value,
    description: taskDescription.value,
    tag: taskTag.value,
  };

  tasks.value.push(newTask);

  try {
    await window.backend.App.SaveTasks();
  } catch (error) {
    console.error("Error saving tasks:", error);
  }

  // Clear input fields after adding
  taskTitle.value = "";
  taskDescription.value = "";
  taskTag.value = "";
};

const deleteTask = (taskId) => {
  tasks.value = tasks.value.filter((task) => task.id !== taskId);
};

const editTask = (taskId) => {
  const taskToEdit = tasks.value.find((task) => task.id === taskId);
  if (taskToEdit) {
    taskTitle.value = taskToEdit.title;
    taskDescription.value = taskToEdit.description;
    taskTag.value = taskToEdit.tag;
    deleteTask(taskId);
  }
};
</script>

<template>
  <div id="app">
    <img :src="logoImg" alt="Logo" class="logo" />
    <h1>Froggy</h1>
    <img href="./assets/images/theoneFROG.svg" />
    <div class="form-container">
      <div class="input-wrapper">
        <textarea
          v-model="taskTitle"
          placeholder="Task Title"
          class="input-field"
        />
      </div>
      <div class="input-wrapper">
        <textarea
          v-model="taskDescription"
          placeholder="Task Description"
          class="input-field"
        ></textarea>
      </div>
      <div class="input-wrapper">
        <textarea
          v-model="taskTag"
          placeholder="Tag"
          class="input-field"
        ></textarea>
      </div>
      <div class="input-wrapper">
        <button @click="addTask">Add Task</button>
      </div>
    </div>
    <div class="tasks">
      <div v-for="task in tasks" :key="task.id" class="task">
        <div class="task-header">
          <h3>{{ task.title }}</h3>
          <p>#{{ task.tag }}</p>
        </div>
        <p class="task-body">{{ task.description }}</p>
        <div class="task-footer">
          <button class="button-mod" @click="editTask(task.id)">Edit</button>
          <button class="button-mod" @click="deleteTask(task.id)">
            Delete
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

/* Style your components here */
<style>
.logo {
  display: block;
  max-width: 200px;
  margin: 0 auto 20px;
}

.form-container {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  padding: 8px;
}

.input-wrapper {
  flex-grow: 1;
}

.input-field {
  padding: 8px;
  margin-bottom: 0;
  background: rgba(255, 255, 255, 0.2); /* Semi-transparent background */
  border-radius: 10px;
  border: 0px solid rgba(255, 130, 255, 0.3); /* Light border for glass effect */
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Soft shadow for depth */
  backdrop-filter: blur(10px); /* Blur effect for the glassmorphism */
  color: #000; /* Adjust text color if needed */
}

.input-field,
button {
  width: 100%;
  box-sizing: border-box;
}

.input-field {
  padding: 8px;
  margin-bottom: 0;
  border-radius: 4px;
  box-shadow: 1px 2px 2px 2px rgba(0, 0, 0, 0.2);
}

textarea.input-field {
  height: 38px;
  resize: vertical;
}

button {
  padding: 10px;
  background-color: rgba(223, 140, 38, 1);
  border: none;
  color: white;
  cursor: pointer;
  box-shadow: 1px 2px 2px 2px rgba(0, 0, 0, 0.2);
  border-radius: 4px;
}

button:hover {
  background-color: #ffa02b;
}

.button-mod {
  padding: 4px 8px !important; /* Increase padding for more internal space */
  margin: 0; /* Ensure no external margins affect layout */
  background-color: rgba(223, 140, 38, 1); /* Use your existing button color */
  border: none;
  color: white;
  cursor: pointer;
  border-radius: 4px; /* Maintain rounded corners */
  box-shadow: 1px 2px 2px rgba(0, 0, 0, 0.2); /* Reapply shadow for consistency */
  transition: background-color 0.3s ease; /* Smooth transition for hover effect */
}

.button-mod:hover {
  background-color: #ffa02b; /* Darker shade on hover for better interaction feedback */
}

.form-container,
.tasks {
  max-width: 600px;
  margin: 0 auto;
}

.input-wrapper,
.task {
  width: 100%;
}

.task-header {
  display: flex;
  justify-content: space-between;
  width: 100%;
  align-items: baseline;
}

.task-body {
  display: flex;
  justify-content: flex-start;
  width: 100%;
}

.task-footer {
  position: absolute;
  bottom: 10px;
  right: 15px;
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  padding: 12px 20px;
}

.task {
  display: flex;
  position: relative;
  flex-direction: column;
  justify-content: space-between;
  gap: 8px;
  margin-top: 10px;
  padding: 15px; /* Adjusted padding */
  background: rgba(255, 255, 255, 0.2); /* Semi-transparent background */
  border-radius: 10px;
  border: 1px solid rgba(255, 255, 255, 0.3); /* Light border for glass effect */
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Soft shadow for depth */
  backdrop-filter: blur(10px); /* Blur effect for the glassmorphism */
  color: #000; /* Adjust text color if needed */
}
.task * {
  margin: 0;
  padding: 0;
}
</style>
