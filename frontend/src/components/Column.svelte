<script lang="ts">
    import { tick, onMount } from "svelte";
    import Sortable from "sortablejs";
    import Task from "./Task.svelte";
    import {
        moveTask,
        addTask,
        type Column,
        type Task as TaskModel,
    } from "../lib/api";

    export let column: Column;
    export let panelId: number;

    let isCreating = false;
    let inputEl: HTMLInputElement;
    let draggedEl: HTMLDivElement;

    onMount(() => {
        Sortable.create(draggedEl, {
            group: "tasks",
            animation: 150,
            onEnd(evt: any) {
                const id = evt.item.dataset.id;
                moveTask(panelId, column.id, id);
            },
        });
    });
</script>

<div class="card bg-base-100 shadow w-64 min-w-[16rem]">
    <div class="card-body p-3">
        <!-- Header -->
        <div class="flex justify-between items-center">
            <!-- Header -->
            <h2 class="font-bold">{column.title}</h2>
            <button
                class="btn btn-xs btn-primary btn-circle"
                on:click={async () => {
                    isCreating = true;
                    await tick();
                    inputEl?.focus();
                }}
            >
                +
            </button>
        </div>
        <!-- Tasks -->
        <div bind:this={draggedEl} class="space-y-2 mt-2 min-h-[150px] p-3">
            {#each column.tasks as task}
                <Task {task} />
            {/each}
            {#if isCreating}
                <Task
                    task={null}
                    isAddMode={true}
                    bind:inputEl
                    onSave={async (title) => {
                        if (!title.trim()) {
                            isCreating = false;
                            return;
                        }

                        const t = await addTask(panelId, column.id, title);
                        if (t) {
                            column.tasks = [...column.tasks, t];
                        }

                        isCreating = false;
                    }}
                    onCancel={() => (isCreating = false)}
                />
            {/if}
        </div>
    </div>
</div>
