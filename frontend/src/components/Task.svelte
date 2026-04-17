<script lang="ts">
    import type { Task } from "../lib/api";
    export let task: Task;
</script>

<div class="card bg-base-100 shadow p-2" data-id={task.id}>
    {task.title}
</div>
