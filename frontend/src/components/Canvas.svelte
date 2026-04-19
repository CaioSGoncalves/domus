<script lang="ts">
    import { onMount } from "svelte";
    import Panel from "./Panel.svelte";
    import { getPanels, type Panel as PanelType } from "../lib/api";

    let panels: PanelType[] = [];
    let current = 0;

    onMount(async () => {
        panels = await getPanels();
    });
</script>

<div class="p-4 space-y-4">
    <!-- Tabs -->
    <div class="tabs tabs-boxed w-full justify-center">
        {#each panels as p, i}
            <button
                class="tab {i === current ? 'tab-active' : ''}"
                on:click={() => (current = i)}
            >
                {p.name}
            </button>
        {/each}
    </div>

    <!-- Panel -->
    {#if panels.length}
        <Panel panel={panels[current]} />
    {/if}
</div>
