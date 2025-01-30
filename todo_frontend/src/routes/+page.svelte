<script>
  import { onMount } from "svelte";
  let tasks =
    /** @type {Array<{ ID: number, Title: string, Completed: boolean, CreatedAt: string}>} */ ([]);
  let newTask = "";

  async function fetchTasks() {
    const response = await fetch("http://localhost:8000/tasks");
    tasks = await response.json();
  }

  async function addTask() {
    await fetch("http://localhost:8000/tasks", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: newTask }),
    });
    newTask = "";
    fetchTasks();
  }

  /**
   * @param {number} id
   */
  async function deleteTask(id) {
    await fetch(`http://localhost:8000/tasks/${id}`, {
      method: "DELETE",
    });
    fetchTasks();
  }
  /**
   * @param {number} id
   * @param {any} Completed
   * @param {string} Title
   */
  async function updateTask(id, Title, Completed) {
    await fetch(`http://localhost:8000/tasks/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ Title, Completed }),
    });
    fetchTasks();
  }

  onMount(() => {
    fetchTasks();
  });
</script>

<div class="page-container">
  <div class="container">
    <h1>Todo App</h1>
    <input bind:value={newTask} placeholder="Add a new task" />
    <button on:click={addTask}>Add</button>

    <ul>
      {#each tasks as task}
        <li>
          <div>
            <input type="text" bind:value={task.Title} />
            <input type="checkbox" bind:checked={task.Completed} />
            <span
              ><strong>Created At:</strong>
              {new Date(task.CreatedAt).toLocaleString()}</span
            >
          </div>
          <div>
            <button on:click={() => deleteTask(task.ID)}>Delete</button>
            <button
              on:click={() => updateTask(task.ID, task.Title, task.Completed)}
              >Update</button
            >
          </div>
        </li>
      {/each}
    </ul>
  </div>
</div>

<style>
  .page-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #a7cbc4;
    font-family: Arial, sans-serif;
  }

  .container {
    background: #7dc76d;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    width: 400px;
    max-height: 90vh;
    display: flex;
    flex-direction: column;
  }

  h1 {
    margin-bottom: 20px;
    font-size: 24px;
    text-align: center;
    color: #333;
  }

  input[type="text"] {
    width: calc(100% - 22px);
    padding: 10px;
    margin-bottom: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }

  button {
    padding: 10px 15px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    background-color: #28a745;
    color: #fff;
    margin-right: 10px;
  }

  button:hover {
    background-color: #218838;
  }

  ul {
    list-style: none;
    padding: 0;
    margin: 0;
    overflow-y: auto;
    max-height: 60vh;
  }

  li {
    background: #f9f9f9;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 4px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  li div {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  span {
    font-size: 14px;
    color: #555;
  }

  input[type="checkbox"] {
    transform: scale(1.2);
  }
</style>
