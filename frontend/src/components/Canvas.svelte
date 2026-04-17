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
    <div class="tabs">
        {#each panels as p, i}
            <a
                class="tab {i === current ? 'tab-active' : ''}"
                on:click={() => (current = i)}
            >
                {p.name}
            </a>
        {/each}
    </div>

    {#if panels.length}
        <Panel panel={panels[current]} />
    {/if}
</div>
