<script lang="ts">
    import Column from "./Column.svelte";
    import type { Panel } from "../lib/api";

    export let panel: Panel;
</script>

<div class="flex gap-4 overflow-x-auto pb-4">
    {#each panel.columns as col}
        <Column column={col} panelId={panel.id} />
    {/each}
</div>
