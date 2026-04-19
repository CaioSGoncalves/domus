<script lang="ts">
    import type { Task } from "../lib/api";
    export let task: Task | null;
    export let isAddMode: boolean = false;
    export let onSave: (title: string) => void;
    export let onCancel: () => void;
    export let inputEl: HTMLInputElement;

    let title = task?.title ?? "";
</script>

{#if isAddMode}
    <div class="card bg-base-100 shadow-sm border border-base-300">
        <div class="card-body p-1 group">
            <input
                class="input"
                bind:this={inputEl}
                bind:value={title}
                placeholder="Task title..."
                on:keydown={(e) => {
                    if (e.key === "Enter") onSave(title);
                    if (e.key === "Escape") onCancel();
                }}
                on:blur={() => onSave(title)}
            />
        </div>
    </div>
{:else}
    <div
        class="card bg-base-200 shadow-sm hover-shadow-md p-2 cursor-grab active:cursor-grabbing hover:ring-2 hover:ring-primary transition"
        data-id={task?.id}
    >
        <div class="card-body p-1 group">
            <div class="flex items-center justify-between">
                <p class="text-sm flex-1">{task?.title}</p>
                <button
                    class="btn btn-ghost btn-xs opacity-0 group-hover:opacity-100 transition"
                >
                    ✏️
                </button>
            </div>
        </div>
    </div>
{/if}
