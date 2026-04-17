<script lang="ts">
    import { onMount } from "svelte";
    import Sortable from "sortablejs";
    import Task from "./Task.svelte";
    import { moveTask, type Column } from "../lib/api";

    export let column: Column;
    let el: HTMLDivElement;

    onMount(() => {
        Sortable.create(el, {
            group: "tasks",
            animation: 150,
            onEnd(evt) {
                const id = evt.item.dataset.id;
                moveTask(id, column.id);
            },
        });
    });
</script>

<div class="bg-base-200 p-3 rounded w-64">
    <h2 class="font-bold mb-2">{column.title}</h2>

    <div bind:this={el} class="space-y-2 min-h-[50px]">
        {#each column.tasks as task}
            <Task {task} />
        {/each}
    </div>
</div>
